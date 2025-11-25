package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"gintama/internal/app"
	"gintama/internal/dto"
	"gintama/internal/lib"
	"gintama/internal/lib/argon2"
	"gintama/internal/lib/constant"
	"gintama/internal/lib/jwt"
	"gintama/internal/models"
	"gintama/internal/services"
	"gintama/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type authHandler struct {
	app *app.Application
}

func (h *authHandler) SignUp(c *gin.Context) {
	var dto dto.AuthSignUp

	if err := lib.ValidateRequestBody(c, &dto); err != nil {
		switch e := err.(type) {
		case *lib.ErrValidationFailed:
			c.JSON(http.StatusBadRequest, lib.WrapValidationError(e.MessageRecord))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	user := &models.User{
		Base: models.Base{
			ID: uuid.Must(uuid.NewV7()),
		},
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Phone:     dto.Phone,
		Password:  &dto.Password,
		RoleID:    uuid.Must(uuid.Parse(constant.RoleUser)),
	}

	userVerifyAccount := &models.UserVerifyAccount{}

	err := lib.WithTransaction(h.app.Repositories.User.DB, func(tx *sql.Tx) error {
		err := user.BeforeCreate()
		if err != nil {
			return err
		}

		err = h.app.Repositories.User.InsertExec(tx, user)
		if err != nil {
			return err
		}

		jsonWebToken := jwt.New(&h.app.Config.App)
		token, expiresIn, err := jsonWebToken.Generate(&jwt.JWTPayload{
			UID:       user.ID.String(),
			Secret:    h.app.Config.App.JWTSecret,
			ExpiresAt: "1", // 1 day
		})
		if err != nil {
			return err
		}

		userVerifyAccount.ID = user.ID
		userVerifyAccount.Token = token
		userVerifyAccount.ExpiresAt = time.Unix(expiresIn, 0)

		return h.app.Repositories.UserVerifyAccount.InsertExec(tx, userVerifyAccount)
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	link := fmt.Sprintf("%s/verify?token=%s", h.app.Config.App.ClientURL, userVerifyAccount.Token)

	emailForm := struct {
		Fullname string
		Link     string
		AppName  string
	}{
		Fullname: strings.Join([]string{dto.FirstName, *dto.LastName}, " "),
		Link:     link,
		AppName:  h.app.Config.App.Name,
	}

	_, err = h.app.Services.Email.SendEmail(services.SendEmailParams{
		Subject:      "Verify your email address",
		To:           dto.Email,
		Data:         emailForm,
		HtmlTemplate: "templates/emails/registration.html",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sign up successfully",
	})
}

func (h *authHandler) SignIn(c *gin.Context) {
	var dto dto.AuthSignIn

	if err := lib.ValidateRequestBody(c, &dto); err != nil {
		switch e := err.(type) {
		case *lib.ErrValidationFailed:
			c.JSON(http.StatusBadRequest, lib.WrapValidationError(e.MessageRecord))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	user, err := h.app.Repositories.User.GetByEmail(dto.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	hash := argon2.New()
	match, err := hash.Compare(*user.Password, dto.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}

	jsonWebToken := jwt.New(&h.app.Config.App)
	token, expiresIn, err := jsonWebToken.Generate(&jwt.JWTPayload{
		UID:       user.ID.String(),
		Secret:    h.app.Config.App.JWTSecret,
		ExpiresAt: "1", // 1 day
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	session := &models.Session{
		Base: models.Base{
			ID: uuid.Must(uuid.NewV7()),
		},
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Unix(expiresIn, 0),
		IPAddress: c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
	}

	err = h.app.Repositories.Session.Insert(session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[any]{
		Message: "Sign in successfully",
		Data: gin.H{
			"uid":          user.ID.String(),
			"email":        user.Email,
			"display_name": strings.Join([]string{user.FirstName, *user.LastName}, " "),
			"is_admin":     user.RoleID.String() == constant.RoleAdmin,
			"access_token": token,
		},
	})
}

func (h *authHandler) VerifyRegistration(c *gin.Context) {
	var dto dto.AuthVerifyRegistration

	if err := lib.ValidateRequestBody(c, &dto); err != nil {
		switch e := err.(type) {
		case *lib.ErrValidationFailed:
			c.JSON(http.StatusBadRequest, lib.WrapValidationError(e.MessageRecord))
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	jsonWebToken := jwt.New(&h.app.Config.App)
	claims, err := jsonWebToken.Verify(dto.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	userID := uuid.Must(uuid.Parse(claims.UID))

	userVerifyAccount, err := h.app.Repositories.UserVerifyAccount.Get(userID, dto.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if userVerifyAccount.ExpiresAt.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Token expired"})
		return
	}

	user, err := h.app.Repositories.User.Get(userVerifyAccount.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user.ActiveAt = lib.TimePtr(time.Now())

	err = h.app.Repositories.User.Update(user.ID, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Verify registration successfully",
	})
}

func (h *authHandler) VerifySession(c *gin.Context) {
	uid, err := lib.ContextGetUID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	user, err := h.app.Repositories.User.Get(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.ResponseSingleData[*models.User]{
		Message: "Verify session successfully",
		Data:    user,
	})
}

func (h *authHandler) SignOut(c *gin.Context) {
	jwt := jwt.New(&h.app.Config.App)

	extractToken, err := jwt.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("Unauthorized, %s", err.Error())})
		return
	}

	uid, err := lib.ContextGetUID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	err = h.app.Repositories.Session.Delete(uid, extractToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sign out successfully",
	})
}

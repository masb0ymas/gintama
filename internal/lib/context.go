package lib

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ContextGetUID(c *gin.Context) (uuid.UUID, error) {
	if uid, exists := c.Get("uid"); exists {
		uid := uuid.MustParse(uid.(string))
		return uid, nil
	}

	return uuid.Nil, errors.New("can't find get context auth, please check your authorization")
}

func ContextSetUID(c *gin.Context, uid uuid.UUID) {
	c.Set("uid", uid.String())
}

func ContextParamUUID(c *gin.Context, key string) (uuid.UUID, error) {
	str := c.Param(key)
	return uuid.Parse(str)
}

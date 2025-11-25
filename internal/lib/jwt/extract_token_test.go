package jwt

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gintama/internal/config"

	"github.com/gin-gonic/gin"
)

func TestExtractToken(t *testing.T) {
	tests := []struct {
		name      string
		url       string
		headers   map[string]string
		wantToken string
		wantErr   bool
	}{
		{
			name:      "Extract token from query parameter",
			url:       "/test?token=query-token-123",
			headers:   map[string]string{},
			wantToken: "query-token-123",
			wantErr:   false,
		},
		{
			name: "Extract token from cookie",
			url:  "/test",
			headers: map[string]string{
				"Cookie": "token=cookie-token-456",
			},
			wantToken: "cookie-token-456",
			wantErr:   false,
		},
		{
			name: "Extract token from Authorization header with Bearer",
			url:  "/test",
			headers: map[string]string{
				"Authorization": "Bearer header-token-789",
			},
			wantToken: "header-token-789",
			wantErr:   false,
		},
		{
			name: "Query parameter takes precedence over cookie",
			url:  "/test?token=query-token",
			headers: map[string]string{
				"Cookie": "token=cookie-token",
			},
			wantToken: "query-token",
			wantErr:   false,
		},
		{
			name: "Query parameter takes precedence over header",
			url:  "/test?token=query-token",
			headers: map[string]string{
				"Authorization": "Bearer header-token",
			},
			wantToken: "query-token",
			wantErr:   false,
		},
		{
			name: "Cookie takes precedence over header",
			url:  "/test",
			headers: map[string]string{
				"Cookie":        "token=cookie-token",
				"Authorization": "Bearer header-token",
			},
			wantToken: "cookie-token",
			wantErr:   false,
		},
		{
			name: "Invalid Authorization header - missing Bearer prefix",
			url:  "/test",
			headers: map[string]string{
				"Authorization": "token-without-bearer",
			},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "Invalid Authorization header - only Bearer without token",
			url:  "/test",
			headers: map[string]string{
				"Authorization": "Bearer ",
			},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "Invalid Authorization header - Bearer with empty token",
			url:  "/test",
			headers: map[string]string{
				"Authorization": "Bearer  ",
			},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "Invalid Authorization header - wrong prefix",
			url:  "/test",
			headers: map[string]string{
				"Authorization": "Basic token-123",
			},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "Invalid Authorization header - too many parts",
			url:  "/test",
			headers: map[string]string{
				"Authorization": "Bearer token extra-part",
			},
			wantToken: "",
			wantErr:   true,
		},
		{
			name:      "No token provided anywhere",
			url:       "/test",
			headers:   map[string]string{},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "Empty query parameter is ignored",
			url:  "/test?token=",
			headers: map[string]string{
				"Authorization": "Bearer fallback-token",
			},
			wantToken: "fallback-token",
			wantErr:   false,
		},
		{
			name:      "Token with special characters in query",
			url:       "/test?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U",
			headers:   map[string]string{},
			wantToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U",
			wantErr:   false,
		},
		{
			name: "Token with special characters in cookie",
			url:  "/test",
			headers: map[string]string{
				"Cookie": "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U",
			},
			wantToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U",
			wantErr:   false,
		},
		{
			name: "Token with special characters in Authorization header",
			url:  "/test",
			headers: map[string]string{
				"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U",
			},
			wantToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U",
			wantErr:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			app := gin.New()
			j := &JWT{
				config: &config.ConfigApp{
					Name:      "gintama",
					JWTSecret: "test-secret",
				},
			}

			var gotToken string
			var gotErr error

			app.GET("/test", func(c *gin.Context) {
				gotToken, gotErr = j.ExtractToken(c)
				c.String(http.StatusOK, "ok")
			})

			req := httptest.NewRequest("GET", tc.url, nil)
			for key, value := range tc.headers {
				req.Header.Set(key, value)
			}
			w := httptest.NewRecorder()
			app.ServeHTTP(w, req)

			if (gotErr != nil) != tc.wantErr {
				t.Errorf("ExtractToken() error = %v, wantErr %v", gotErr, tc.wantErr)
				return
			}

			if gotToken != tc.wantToken {
				t.Errorf("ExtractToken() token = %v, want %v", gotToken, tc.wantToken)
			}
		})
	}
}

func TestExtractTokenErrorMessages(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		headers        map[string]string
		wantErrMessage string
	}{
		{
			name:           "No token - token not found error",
			url:            "/test",
			headers:        map[string]string{},
			wantErrMessage: "token not found",
		},
		{
			name: "Invalid format - missing Bearer",
			url:  "/test",
			headers: map[string]string{
				"Authorization": "just-a-token",
			},
			wantErrMessage: "invalid token format",
		},
		{
			name: "Invalid format - empty token after Bearer",
			url:  "/test",
			headers: map[string]string{
				"Authorization": "Bearer ",
			},
			wantErrMessage: "invalid token format",
		},
		{
			name: "Invalid format - too many parts",
			url:  "/test",
			headers: map[string]string{
				"Authorization": "Bearer token1 token2",
			},
			wantErrMessage: "invalid token format",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			app := gin.New()
			j := &JWT{
				config: &config.ConfigApp{
					Name:      "gintama",
					JWTSecret: "test-secret",
				},
			}

			var gotErr error

			app.GET("/test", func(c *gin.Context) {
				_, gotErr = j.ExtractToken(c)
				c.String(http.StatusOK, "ok")
			})

			req := httptest.NewRequest("GET", tc.url, nil)
			for key, value := range tc.headers {
				req.Header.Set(key, value)
			}
			w := httptest.NewRecorder()
			app.ServeHTTP(w, req)

			if gotErr == nil {
				t.Errorf("ExtractToken() expected error but got nil")
				return
			}

			if gotErr.Error() != tc.wantErrMessage {
				t.Errorf("ExtractToken() error message = %v, want %v", gotErr.Error(), tc.wantErrMessage)
			}
		})
	}
}

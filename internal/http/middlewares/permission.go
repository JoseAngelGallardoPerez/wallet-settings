package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Confialink/wallet-settings/internal/http/handlers"
	"github.com/Confialink/wallet-settings/internal/http/responses"
	"github.com/Confialink/wallet-settings/internal/service/auth"
)

type PermissionChecker struct {
	authService     *auth.Service
	responseService responses.ResponseHandler
}

func NewPermissionChecker(authService *auth.Service, responseService responses.ResponseHandler) *PermissionChecker {
	return &PermissionChecker{authService, responseService}
}

func (s *PermissionChecker) CanDynamic(action, resource string) func(*gin.Context) {
	return func(c *gin.Context) {
		user := handlers.GetCurrentUser(c)
		if user == nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		if !s.authService.CanDynamic(user, action, resource) {
			s.responseService.ErrorResponse(c, responses.CodeForbidden, "")
			return
		}
	}
}

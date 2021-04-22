package middlewares

import (
	"errors"
	"net/http"

	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
	"github.com/gin-gonic/gin"
)

func AdminOrRoot() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, ok := ctx.Get("_user")
		if !ok {
			// Returns a "403 StatusForbidden" response
			ctx.Status(http.StatusForbidden)
			ctx.Error(errors.New("Access token not found"))
			ctx.Abort()
			return
		}

		rolename := user.(*userpb.User).RoleName
		if rolename != "admin" && rolename != "root" { // TODO: sync with permissions service
			// Returns a "403 StatusForbidden" response
			ctx.Status(http.StatusForbidden)
			ctx.Error(errors.New("User does not have the right roles"))
			ctx.Abort()
			return
		}
	}
}

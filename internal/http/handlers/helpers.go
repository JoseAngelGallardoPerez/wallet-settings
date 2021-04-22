package handlers

import (
	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
	"github.com/gin-gonic/gin"
)

// getCurrentUser returns current user or nil
func GetCurrentUser(c *gin.Context) *userpb.User {
	user, ok := c.Get("_user")
	if !ok {
		return nil
	}
	return user.(*userpb.User)
}

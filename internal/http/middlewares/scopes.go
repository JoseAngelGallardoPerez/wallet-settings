package middlewares

import (
	"github.com/Confialink/wallet-settings/internal/db/models"
	"github.com/gin-gonic/gin"
)

func ScopesMiddleware(c *gin.Context) {
	_, ok := c.Get("_user")
	if !ok {
		c.Set("_scopes", []string{models.ScopePublic})
		return
	}

	c.Set("_scopes", []string{models.ScopePublic, models.ScopePrivate})
}

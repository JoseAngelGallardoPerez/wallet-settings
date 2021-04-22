package validators

import (
	"github.com/Confialink/wallet-settings/internal/db/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdateConfigsValidator is
type UpdateConfigsValidator struct {
	Data []*models.Config `json:"data" binding:"required,dive,required"`
}

// BindJSON binding from JSON
func (s *UpdateConfigsValidator) BindJSON(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(s, b)
	if err != nil {
		return err
	}
	return nil
}

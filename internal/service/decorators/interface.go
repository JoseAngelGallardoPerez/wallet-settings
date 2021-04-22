package decorators

import (
	"github.com/Confialink/wallet-settings/internal/db/models"
)

type Decorator interface {
	// return modified config from decorator
	GetSetting(setting models.Config) (*models.Config, error)
}

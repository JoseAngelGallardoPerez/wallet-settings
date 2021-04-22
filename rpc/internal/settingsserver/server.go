package settingsserver

import (
	"github.com/Confialink/wallet-settings/internal/service"
)

// SettingsHandlerServer implements the notifications service
type SettingsHandlerServer struct {
	SettingsProvider *service.SettingsProvider
}

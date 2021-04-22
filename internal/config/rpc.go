package config

import (
	"github.com/Confialink/wallet-pkg-env_config"
)

// RPCConfiguration is rpc config model
type RPCConfiguration struct {
	SettingsServerPort string
}

// GetSettingsServerPort returns rpc port for settingsserver
func (s *RPCConfiguration) GetSettingsServerPort() string {
	return s.SettingsServerPort
}

// Init initializes enviroment variables
func (s *RPCConfiguration) Init() error {
	s.SettingsServerPort = env_config.Env("VELMIE_WALLET_SETTINGS_RPC_PORT", "")
	return nil
}

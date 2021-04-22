package settingsserver

import (
	"fmt"
	"net/http"

	"github.com/Confialink/wallet-settings/internal/config"
	"github.com/Confialink/wallet-settings/internal/service"
	server "github.com/Confialink/wallet-settings/rpc/internal/settingsserver"
	pb "github.com/Confialink/wallet-settings/rpc/proto/settings"
)

// SettingsServer implements the Users service
type SettingsServer struct {
	SettingsProvider *service.SettingsProvider
}

// Init initializes users rpc server
func (s *SettingsServer) Init() {
	// Retrieve config options.
	conf := config.GetConf()

	server := &server.SettingsHandlerServer{SettingsProvider: s.SettingsProvider}

	twirpHandler := pb.NewSettingsHandlerServer(server, nil)

	mux := http.NewServeMux()
	mux.Handle(pb.SettingsHandlerPathPrefix, twirpHandler)

	go http.ListenAndServe(fmt.Sprintf(":%s", conf.RPC.GetSettingsServerPort()), mux)
}

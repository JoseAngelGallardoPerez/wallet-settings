package main

import (
	"log"

	"github.com/Confialink/wallet-pkg-env_mods"
	"github.com/gin-gonic/gin"

	"github.com/Confialink/wallet-settings/internal/app/di"
	"github.com/Confialink/wallet-settings/internal/http/routes"
	"github.com/Confialink/wallet-settings/rpc/cmd/server/settingsserver"
)

func main() {
	c := di.Container
	conf := c.Config().GetServer()

	ginMode := env_mods.GetMode(conf.GetEnv())
	gin.SetMode(ginMode)

	// Run the rpc server.
	rpc := &settingsserver.SettingsServer{SettingsProvider: c.SettingsProvider()}
	rpc.Init()

	// Register routes to be used.
	r := routes.RegisterRoutes()

	log.Printf("Listening on : %s", conf.GetPort())

	// Run the server
	r.Run(":" + conf.GetPort())
}

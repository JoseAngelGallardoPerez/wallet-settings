package routes

import (
	"net/http"

	"github.com/Confialink/wallet-pkg-env_mods"
	errorsPkg "github.com/Confialink/wallet-pkg-errors"
	"github.com/gin-gonic/gin"

	"github.com/Confialink/wallet-settings/internal/app/di"
	"github.com/Confialink/wallet-settings/internal/authentication"
	"github.com/Confialink/wallet-settings/internal/http/handlers"
	"github.com/Confialink/wallet-settings/internal/http/middlewares"
	"github.com/Confialink/wallet-settings/internal/http/responses"
	"github.com/Confialink/wallet-settings/internal/service/auth"
	"github.com/Confialink/wallet-settings/internal/version"
)

var container = di.Container

// RegisterRoutes is where you can register all of the routes for an service.
func RegisterRoutes() *gin.Engine {
	// Retrieve config options.
	conf := container.Config()
	ginMode := env_mods.GetMode(conf.GetServer().GetEnv())
	gin.SetMode(ginMode)

	// Creates a gin router with default middleware:
	r := gin.Default()

	r.GET("/settings/health-check", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/settings/build", func(c *gin.Context) {
		c.JSON(http.StatusOK, version.BuildInfo)
	})

	errorsMiddleware := errorsPkg.ErrorHandler(container.Logger().New("Middleware", "Errors"))

	r.Use(
		middlewares.CorsMiddleware(),
		gin.Recovery(),
		gin.Logger(),
		errorsMiddleware,
	)

	apiGroup := r.Group("settings")

	configHandler := &handlers.ConfigHandler{
		container.Repository(),
		container.Response(),
		container.SettingsProvider(),
		container.ConfigValidator(),
	}
	mwPermissions := middlewares.NewPermissionChecker(container.AuthService(), container.Response())

	/*
	 |---------------------------------------------------
	 | Private router group
	 |---------------------------------------------------
	*/
	privateGroup := apiGroup.Group("/private")
	{
		v1Group := privateGroup.Group("/v1", authentication.Middleware(container.Logger().New("Middleware", "Authentication")))
		{
			configGroup := v1Group.Group("/config")
			{
				// GET /settings/private/v1/config/:section
				configGroup.GET("/:section", middlewares.ScopesMiddleware, configHandler.ListConfigHandler)
				// GET /settings/private/v1/config/:section/:group
				configGroup.GET("/:section/:group", middlewares.ScopesMiddleware, configHandler.ListConfigHandler)
				// GET /settings/private/v1/config/:section/:group/:field
				configGroup.GET("/:section/:group/:field", middlewares.ScopesMiddleware, configHandler.GetConfigHandler)

				// PUT /settings/private/v1/config
				configGroup.PUT("", middlewares.AdminOrRoot(), mwPermissions.CanDynamic(auth.ActionUpdate, auth.ResourceSettings), configHandler.UpdateConfigsHandler)
			}
		}
	}

	/*
	 |---------------------------------------------------
	 | Public router group
	 |---------------------------------------------------
	*/
	publicGroup := apiGroup.Group("/public")
	{
		v1Group := publicGroup.Group("/v1")
		{
			configGroup := v1Group.Group("/config", middlewares.ScopesMiddleware)
			{
				// GET /settings/private/v1/config/:section
				configGroup.GET("/:section", configHandler.ListConfigHandler)
				// GET /settings/private/v1/config/:section/:group
				configGroup.GET("/:section/:group", configHandler.ListConfigHandler)
				// GET /settings/private/v1/config/:section/:group/:field
				configGroup.GET("/:section/:group/:field", configHandler.GetConfigHandler)
			}
		}
	}

	// If route not found returns StatusNotFound
	r.NoRoute(NotFound)

	// Handle OPTIONS request
	r.OPTIONS("/*cors", func(c *gin.Context) {
		c.Status(http.StatusOK)
		c.Abort()
		return
	})
	return r
}

// NotFound returns 404 NotFound
func NotFound(c *gin.Context) {
	publicError := errorsPkg.PublicError{Code: responses.NotFound, HttpStatus: http.StatusNotFound}
	errorsPkg.AddErrors(c, &publicError)
	c.Abort()
	return
}

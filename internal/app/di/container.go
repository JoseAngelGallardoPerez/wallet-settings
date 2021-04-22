package di

import (
	"log"

	"github.com/Confialink/wallet-pkg-service_names"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/inconshreveable/log15"
	"github.com/jinzhu/gorm"

	"github.com/Confialink/wallet-settings/internal/config"
	"github.com/Confialink/wallet-settings/internal/db"
	"github.com/Confialink/wallet-settings/internal/db/repositories"
	"github.com/Confialink/wallet-settings/internal/http/responses"
	"github.com/Confialink/wallet-settings/internal/service"
	"github.com/Confialink/wallet-settings/internal/service/auth"
	"github.com/Confialink/wallet-settings/internal/srvdiscovery"
	"github.com/Confialink/wallet-settings/internal/validators"
)

// container is the implementation of the Container interface.
type container struct {
	config           *config.Configuration
	repository       repositories.RepositoryInterface
	dbConnection     *gorm.DB
	response         responses.ResponseHandler
	logger           log15.Logger
	settingsProvider *service.SettingsProvider
	authService      *auth.Service
	configValidator  *validators.SettingsValidator
}

// Container represents a dependency injection container.
var Container *container

func init() {
	Container = new(container)
	// Retrieve config options.
	config.InitConfig(Container.Logger().New("service", "configReader"))
	Container.config = config.GetConf()
}

// Config returns config
func (c *container) Config() *config.Configuration {
	return c.config
}

// Repository creates new repository if not exists and return
func (c *container) Repository() repositories.RepositoryInterface {
	if nil == c.repository {
		c.repository = repositories.NewRepository(c.DbConnection())
	}
	return c.repository
}

// Response creates new response if not exists and return
func (c *container) Response() responses.ResponseHandler {
	if nil == c.response {
		c.response = responses.NewResponseService()
	}
	return c.response
}

// DbConnection creates new DB connection if not exists and return
func (c *container) DbConnection() *gorm.DB {
	var err error
	if nil == c.dbConnection {
		c.dbConnection, err = db.CreateConnection(c.Config().GetDatabase())
		// defer database.Close()
		if nil != err {
			log.Fatalf("Could not connect to DB: %v", err)
		}

		if c.Config().GetDatabase().IsDebugMode {
			c.dbConnection.LogMode(true)
		}
	}
	return c.dbConnection
}

func (c *container) Logger() log15.Logger {
	if c.logger == nil {
		c.logger = log15.New("service", service_names.Settings.Internal)
	}
	return c.logger
}

func (c *container) SettingsProvider() *service.SettingsProvider {
	if c.settingsProvider == nil {
		decoratorsProvider, err := service.NewDecoratorsProvider(c.Logger().New("service", "DecoratorsProvider"))
		if err != nil {
			panic(err.Error())
		}
		c.settingsProvider = service.NewSettingsProvider(
			c.Repository().GetConfigRepository(),
			decoratorsProvider,
			c.Logger().New("service", "SettingsProvider"),
		)
	}
	return c.settingsProvider
}

func (c *container) AuthService() *auth.Service {
	if c.authService == nil {
		c.authService = auth.NewService(auth.NewPermissionService(srvdiscovery.Resolver()))
	}
	return c.authService
}

func (c *container) ConfigValidator() *validators.SettingsValidator {
	if c.configValidator == nil {
		c.configValidator = validators.NewSettingsValidator()
		structValidator, ok := binding.Validator.Engine().(*validator.Validate)
		if !ok {
			panic("cannot init engine validator")
		}
		c.configValidator.Add(validators.NewAutoLogoutValidator(structValidator))
	}
	return c.configValidator
}

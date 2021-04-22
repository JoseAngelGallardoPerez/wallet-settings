package handlers

import (
	"errors"
	"fmt"

	errPkg "github.com/Confialink/wallet-pkg-errors"
	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
	"github.com/gin-gonic/gin"

	"github.com/Confialink/wallet-settings/internal/db/repositories"
	"github.com/Confialink/wallet-settings/internal/http/responses"
	"github.com/Confialink/wallet-settings/internal/service"
	"github.com/Confialink/wallet-settings/internal/validators"
)

type ConfigHandler struct {
	Repository        repositories.RepositoryInterface
	ResponseService   responses.ResponseHandler
	SettingsProvider  *service.SettingsProvider
	SettingsValidator *validators.SettingsValidator
}

// GetConfigHandler returns config by path
func (h *ConfigHandler) GetConfigHandler(c *gin.Context) {
	section := c.Params.ByName("section")
	group := c.Params.ByName("group")
	field := c.Params.ByName("field")

	path := fmt.Sprintf("%s/%s/%s", section, group, field)

	scopes, err := h.getScopesFromContext(c)
	if nil != err {
		// Returns a "400 BadRequest" response
		h.ResponseService.ErrorResponse(c, responses.CannotRetrieveSetting, err.Error())
		return
	}

	config, err := h.SettingsProvider.FindOneByPathAndScopes(path, scopes)
	if nil != err {
		// Returns a "404 StatusNotFound" response
		h.ResponseService.ErrorResponse(c, responses.SettingsNotFound, "Settings not found")
		return
	}

	// Returns a "200 OK" response
	h.ResponseService.OkResponse(c, config)
	return
}

// ListConfigHandler returns list of configs by path
func (h *ConfigHandler) ListConfigHandler(c *gin.Context) {
	section := c.Params.ByName("section")
	group := c.Params.ByName("group")

	var path string

	if group == "" {
		path = fmt.Sprintf("%s/%%", section)
	} else {
		path = fmt.Sprintf("%s/%s/%%", section, group)
	}

	scopes, err := h.getScopesFromContext(c)
	if nil != err {
		// Returns a "400 BadRequest" response
		h.ResponseService.ErrorResponse(c, responses.CanNotRetrieveCollection, err.Error())
		return
	}

	configs, err := h.SettingsProvider.FindByPathAndScopes(path, scopes)
	if nil != err {
		// Returns a "404 StatusNotFound" response
		h.ResponseService.ErrorResponse(c, responses.SettingsNotFound, "Settings not found")
		return
	}

	// Returns a "200 OK" response
	h.ResponseService.OkResponse(c, configs)
	return
}

// UpdateConfigsHandler updates configs
func (h *ConfigHandler) UpdateConfigsHandler(c *gin.Context) {
	user := h.getUserFromContext(c)

	// Checks if the query entry is valid
	validator := validators.UpdateConfigsValidator{}
	if err := validator.BindJSON(c); err != nil {
		errPkg.AddShouldBindError(c, err)
		return
	}

	if err := h.SettingsValidator.Validate(validator.Data); err != nil {
		errPkg.AddShouldBindError(c, err)
		return
	}

	// Update answers
	for _, data := range validator.Data {
		item, err := h.Repository.GetConfigRepository().FindOneByPath(string(data.Path))
		if err != nil {
			// Returns a "500 StatusInternalServerError" response
			h.ResponseService.ErrorResponse(c, responses.CanNotUpdateCollection, "Could not update settings")
			return
		}

		if item.RootOnly && user.RoleName != "root" {
			// Returns a "500 StatusInternalServerError" response
			h.ResponseService.ErrorResponse(c, responses.CanNotUpdateCollection, "Could not update settings")
			return
		}

		_, err = h.Repository.GetConfigRepository().Update(item, data)
		if err != nil {
			// Returns a "500 StatusInternalServerError" response
			h.ResponseService.ErrorResponse(c, responses.CanNotUpdateCollection, "Could not update settings")
			return
		}
	}
	// Returns a "200 OK" response
	h.ResponseService.OkResponse(c, validator.Data)
}

func (h *ConfigHandler) getScopesFromContext(c *gin.Context) ([]string, error) {
	scopes, ok := c.Get("_scopes")
	if !ok {
		return []string{}, errors.New("cannot retrieve scopes from the context")
	}
	return scopes.([]string), nil
}

func (h *ConfigHandler) getUserFromContext(c *gin.Context) *userpb.User {
	user := GetCurrentUser(c)
	if user == nil {
		panic("Can't get user from context")
	}

	return user
}

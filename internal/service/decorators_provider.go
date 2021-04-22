package service

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-settings/internal/db/models"
	"github.com/Confialink/wallet-settings/internal/service/decorators"
)

type DecoratorsProvider struct {
	decorators map[models.Path]decorators.Decorator
}

func NewDecoratorsProvider(logger log15.Logger) (*DecoratorsProvider, error) {
	provider := &DecoratorsProvider{}
	provider.decorators = make(map[models.Path]decorators.Decorator)
	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 3 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   3 * time.Second,
			ResponseHeaderTimeout: 3 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: 3 * time.Second,
	}
	if err := provider.RegisterDecorator(models.PathSiteUrl, decorators.NewAddSchemeToUrl(httpClient, 1, logger, nil)); err != nil {
		logger.Error("cannot register decorator", "err", err)
		return nil, err
	}
	return provider, nil
}

// register a new decorator for setting path
func (d *DecoratorsProvider) RegisterDecorator(path models.Path, decorator decorators.Decorator) error {
	if _, ok := d.decorators[path]; ok {
		return errors.New(fmt.Sprintf("cannot register a decorator for path %s. You can register only one Decorator for each path. You can put the decorator into the previous registered decorator", path))
	}
	d.decorators[path] = decorator

	return nil
}

// get setting after decorator processing
func (d *DecoratorsProvider) GetSetting(setting models.Config) (*models.Config, error) {
	if _, ok := d.decorators[setting.Path]; ok {
		return d.decorators[setting.Path].GetSetting(setting)
	}

	return &setting, nil
}

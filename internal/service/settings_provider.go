package service

import (
	"github.com/Confialink/wallet-settings/internal/db/models"
	"github.com/Confialink/wallet-settings/internal/db/repositories"
	"github.com/inconshreveable/log15"
)

type SettingsProvider struct {
	repo               *repositories.ConfigRepository
	decoratorsProvider *DecoratorsProvider
	logger             log15.Logger
}

func NewSettingsProvider(repo *repositories.ConfigRepository, decoratorsProvider *DecoratorsProvider, logger log15.Logger) *SettingsProvider {
	return &SettingsProvider{repo, decoratorsProvider, logger}
}

// find setting by path
func (s *SettingsProvider) FindOneByPath(path string) (*models.Config, error) {
	setting, err := s.repo.FindOneByPath(path)
	if nil != err {
		logger := s.logger.New("method", "FindOneByPath")
		logger.Error("cannot retrieve setting", "err", err, "path", path)
		return nil, err
	}

	return s.decoratedSetting(setting)
}

// find settings by path
func (s *SettingsProvider) FindByPath(path string) ([]*models.Config, error) {
	settings, err := s.repo.FindByPath(path)
	if nil != err {
		logger := s.logger.New("method", "FindByPath")
		logger.Error("cannot retrieve settings", "err", err, "path", path)
		return nil, err
	}

	return s.decoratedSettings(settings)
}

// find one setting by path and scopes
func (s *SettingsProvider) FindOneByPathAndScopes(path string, scopes []string) (*models.Config, error) {
	setting, err := s.repo.FindOneByPathAndScopes(path, scopes)
	if nil != err {
		logger := s.logger.New("method", "FindByPath")
		logger.Error("cannot retrieve setting", "err", err, "path", path, "scopes", scopes)
		return nil, err
	}

	return s.decoratedSetting(setting)
}

// find settings by path and scopes
func (s *SettingsProvider) FindByPathAndScopes(path string, scopes []string) ([]*models.Config, error) {
	settings, err := s.repo.FindByPathAndScopes(path, scopes)
	if nil != err {
		logger := s.logger.New("method", "FindByPath")
		logger.Error("cannot retrieve settings", "err", err, "path", path, "scopes", scopes)
		return nil, err
	}

	return s.decoratedSettings(settings)
}

// return decorated setting
func (s *SettingsProvider) decoratedSetting(setting *models.Config) (*models.Config, error) {
	decoratedSetting, err := s.decoratorsProvider.GetSetting(*setting)
	if err != nil {
		return nil, err
	}

	return decoratedSetting, nil
}

// return decorated settings
func (s *SettingsProvider) decoratedSettings(settings []*models.Config) ([]*models.Config, error) {
	decoratedSettings := make([]*models.Config, 0, len(settings))
	for _, setting := range settings {
		decoratedSetting, err := s.decoratedSetting(setting)
		if err != nil {
			return nil, err
		}
		decoratedSettings = append(decoratedSettings, decoratedSetting)
	}

	return decoratedSettings, nil
}

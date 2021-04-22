package repositories

import (
	"fmt"

	"github.com/Confialink/wallet-settings/internal/db/models"
	"github.com/jinzhu/gorm"
)

// ConfigRepositoryInterface is interface for repository functionality
// that ought to be implemented manually.
type ConfigRepositoryInterface interface {
}

// ConfigRepository is config repository for CRUD operations.
type ConfigRepository struct {
	db *gorm.DB
}

// NewConfigRepository creates new repository
func NewConfigRepository(db *gorm.DB) ConfigRepositoryInterface {
	return &ConfigRepository{db}
}

// FindOneByPath find config by path
func (repo *ConfigRepository) FindOneByPath(path string) (*models.Config, error) {
	config := &models.Config{}
	if err := repo.db.Where("path = ?", path).
		First(&config).Error; err != nil {
		return nil, fmt.Errorf("Could not find config with path `%s` in database", path)
	}
	return config, nil
}

// FindByPath find configs by path
func (repo *ConfigRepository) FindByPath(path string) ([]*models.Config, error) {
	var configs []*models.Config
	if err := repo.db.Where("path LIKE ?", path).Find(&configs).Error; err != nil {
		return nil, fmt.Errorf("Could not find config with path `%s` in database", path)
	}
	return configs, nil
}

// FindOneByPathAndScopes find config by path and scopes
func (repo *ConfigRepository) FindOneByPathAndScopes(path string, scopes []string) (*models.Config, error) {
	config := &models.Config{}
	if err := repo.db.Where("path = ? AND scope IN (?)", path, scopes).
		First(&config).Error; err != nil {
		return nil, fmt.Errorf("Could not find config with path `%s` in database", path)
	}
	return config, nil
}

// FindByPathAndScopes find configs by path and scopes
func (repo *ConfigRepository) FindByPathAndScopes(path string, scopes []string) ([]*models.Config, error) {
	var configs []*models.Config
	if err := repo.db.Where("path LIKE ? AND scope IN (?)", path, scopes).Find(&configs).Error; err != nil {
		return nil, fmt.Errorf("Could not find config with path `%s` in database", path)
	}
	return configs, nil
}

// FirstOrCreate updates or create if config don't exists
func (repo *ConfigRepository) FirstOrCreate(data *models.Config) (*models.Config, error) {
	err := repo.db.Where(models.Config{Path: data.Path}).
		Assign(map[string]interface{}{
			"path":  data.Path,
			"value": data.Value,
		}).
		FirstOrCreate(&data).Error
	if nil != err {
		return nil, err
	}
	return data, nil
}

// Update updates config
func (repo *ConfigRepository) Update(model *models.Config, data *models.Config) (*models.Config, error) {
	err := repo.db.Model(model).Updates(data).Error
	if nil != err {
		return nil, err
	}
	return data, nil
}

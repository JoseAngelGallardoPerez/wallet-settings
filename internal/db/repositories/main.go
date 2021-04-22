package repositories

import (
	"github.com/jinzhu/gorm"
)

// RepositoryInterface is interface for repository functionality
type RepositoryInterface interface {
	GetConfigRepository() *ConfigRepository
}

// Repository is config repository for CRUD operations.
type Repository struct {
	DB *gorm.DB
}

// NewRepository creates new config repository
func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{db}
}

// GetConfigRepository gets the repository for a config
func (repo *Repository) GetConfigRepository() *ConfigRepository {
	return &ConfigRepository{repo.DB}
}

package repository

import (
	"go-url-shortener/internal/model"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type URLRepository interface {
	Create(url *model.URL) error
	FindByCode(code string) (*model.URL, error)
	IncrementClick(id uuid.UUID) error
}

type urlRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &urlRepository{db}
}

func (r *urlRepository) Create(url *model.URL) error {
	return r.db.Create(url).Error
}

func (r *urlRepository) FindByCode(code string) (*model.URL, error) {
	var url model.URL
	err := r.db.Where("short_code = ?", code).First(&url).Error
	return &url, err
}

func (r *urlRepository) IncrementClick(id uuid.UUID) error {
	return r.db.Model(&model.URL{}).
		Where("id = ?", id).
		UpdateColumn("click_count", gorm.Expr("click_count + 1")).Error
}

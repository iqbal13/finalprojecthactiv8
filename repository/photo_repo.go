package repository

import (
	"context"

	"github.com/iqbal13/finalprojecthactiv8/models"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	Fetch(context.Context, *[]models.Photo) error
	Store(context.Context, *models.Photo) error
	GetByID(context.Context, *models.Photo, string) error
	Update(context.Context, models.Photo, string) (models.Photo, error)
	Delete(context.Context, string) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (photoRepository *photoRepository) Fetch(ctx context.Context, photos *[]models.Photo) (err error) {

	if err = photoRepository.db.WithContext(ctx).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username", "email")
	}).Find(&photos).Error; err != nil {
		return err
	}

	return
}

func (photoRepository *photoRepository) Store(ctx context.Context, photo *models.Photo) (err error) {

	if err := photoRepository.db.WithContext(ctx).Create(&photo).Error; err != nil {
		return err
	}

	return
}

func (photoRepository *photoRepository) GetByID(ctx context.Context, photo *models.Photo, id string) (err error) {

	if err = photoRepository.db.WithContext(ctx).First(&photo, &id).Error; err != nil {
		return err
	}

	return
}

func (photoRepository *photoRepository) Update(ctx context.Context, photo models.Photo, id string) (p models.Photo, err error) {

	p = models.Photo{}

	if err = photoRepository.db.WithContext(ctx).First(&p, &id).Error; err != nil {
		return p, err
	}

	if err = photoRepository.db.WithContext(ctx).Model(&p).Updates(photo).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (photoRepository *photoRepository) Delete(ctx context.Context, id string) (err error) {

	if err = photoRepository.db.WithContext(ctx).First(&models.Photo{}, &id).Error; err != nil {
		return err
	}

	if err = photoRepository.db.WithContext(ctx).Delete(&models.Photo{}, &id).Error; err != nil {
		return err
	}

	return
}

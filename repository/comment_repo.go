package repository

import (
	"context"
	"time"

	"github.com/iqbal13/finalprojecthactiv8/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Fetch(context.Context, *[]models.Comment, string) error
	Store(context.Context, *models.Comment) error
	GetByID(context.Context, *models.Comment, string) error
	Update(context.Context, models.Comment, string) (models.Photo, error)
	Delete(context.Context, string) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (commentRepository *commentRepository) Fetch(ctx context.Context, comments *[]models.Comment, userID string) (err error) {

	if err = commentRepository.db.WithContext(ctx).Where("user_id = ?", userID).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "email", "username", "profile_image_url")
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "user_id", "title", "photo_url", "caption")
	}).Find(&comments).Error; err != nil {
		return err
	}

	return
}

func (commentRepository *commentRepository) Store(ctx context.Context, comment *models.Comment) (err error) {

	if err = commentRepository.db.WithContext(ctx).Create(&comment).Error; err != nil {
		return err
	}

	return
}

func (commentRepository *commentRepository) GetByID(ctx context.Context, comment *models.Comment, id string) (err error) {

	if err = commentRepository.db.WithContext(ctx).First(&comment, &id).Error; err != nil {
		return err
	}

	return
}

func (commentRepository *commentRepository) Update(ctx context.Context, comment models.Comment, id string) (photo models.Photo, err error) {

	c := models.Comment{}

	photo = models.Photo{}

	if err = commentRepository.db.WithContext(ctx).First(&c, &id).Error; err != nil {
		return photo, err
	}

	if err = commentRepository.db.WithContext(ctx).Model(&c).Updates(comment).Error; err != nil {
		return photo, err
	}

	if err = commentRepository.db.WithContext(ctx).First(&photo, comment.PhotoID).Error; err != nil {
		return photo, err
	}

	return photo, nil
}

func (commentRepository *commentRepository) Delete(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = commentRepository.db.WithContext(ctx).First(&models.Comment{}, &id).Error; err != nil {
		return err
	}

	if err = commentRepository.db.WithContext(ctx).Delete(&models.Comment{}, &id).Error; err != nil {
		return err
	}

	return
}

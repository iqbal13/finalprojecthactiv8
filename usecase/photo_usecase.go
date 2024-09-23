package usecase

import (
	"context"

	"github.com/iqbal13/finalprojecthactiv8/models"
	"github.com/iqbal13/finalprojecthactiv8/repository"
)

type PhotoUseCase interface {
	Fetch(context.Context, *[]models.Photo) error
	Store(context.Context, *models.Photo) error
	GetByID(context.Context, *models.Photo, string) error
	Update(context.Context, models.Photo, string) (models.Photo, error)
	Delete(context.Context, string) error
}

type photoUseCase struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoUseCase(photoRepository repository.PhotoRepository) *photoUseCase {
	return &photoUseCase{photoRepository}
}

func (photoUseCase *photoUseCase) Fetch(ctx context.Context, photos *[]models.Photo) (err error) {
	if err = photoUseCase.photoRepository.Fetch(ctx, photos); err != nil {
		return err
	}

	return
}

func (photoUseCase *photoUseCase) Store(ctx context.Context, photo *models.Photo) (err error) {
	if err = photoUseCase.photoRepository.Store(ctx, photo); err != nil {
		return err
	}

	return
}

func (photoUseCase *photoUseCase) GetByID(ctx context.Context, photo *models.Photo, id string) (err error) {
	if err = photoUseCase.photoRepository.GetByID(ctx, photo, id); err != nil {
		return err
	}

	return
}

func (photoUseCase *photoUseCase) Update(ctx context.Context, photo models.Photo, id string) (p models.Photo, err error) {
	if p, err = photoUseCase.photoRepository.Update(ctx, photo, id); err != nil {
		return p, err
	}

	return p, nil
}

func (photoUseCase *photoUseCase) Delete(ctx context.Context, id string) (err error) {
	if err = photoUseCase.photoRepository.Delete(ctx, id); err != nil {
		return err
	}

	return
}

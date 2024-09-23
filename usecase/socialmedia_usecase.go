package usecase

import (
	"context"

	"github.com/iqbal13/finalprojecthactiv8/models"
	"github.com/iqbal13/finalprojecthactiv8/repository"
)

type SocialMediaUseCase interface {
	Fetch(context.Context, *[]models.SocialMedia, string) error
	Store(context.Context, *models.SocialMedia) error
	GetByID(context.Context, *models.SocialMedia, string) error
	Update(context.Context, models.SocialMedia, string) (models.SocialMedia, error)
	Delete(context.Context, string) error
}

type socialMediaUseCase struct {
	socialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaUseCase(socialMediaRepository repository.SocialMediaRepository) *socialMediaUseCase {
	return &socialMediaUseCase{socialMediaRepository}
}

func (socialMediaUseCase *socialMediaUseCase) Fetch(ctx context.Context, socialMedias *[]models.SocialMedia, userID string) (err error) {
	if err = socialMediaUseCase.socialMediaRepository.Fetch(ctx, socialMedias, userID); err != nil {
		return err
	}

	return
}

func (socialMediaUseCase *socialMediaUseCase) Store(ctx context.Context, socialMedia *models.SocialMedia) (err error) {
	if err = socialMediaUseCase.socialMediaRepository.Store(ctx, socialMedia); err != nil {
		return err
	}

	return
}

func (socialMediaUseCase *socialMediaUseCase) GetByID(ctx context.Context, socialMedia *models.SocialMedia, id string) (err error) {
	if err = socialMediaUseCase.socialMediaRepository.GetByID(ctx, socialMedia, id); err != nil {
		return err
	}

	return
}

func (socialMediaUseCase *socialMediaUseCase) Update(ctx context.Context, socialMedia models.SocialMedia, id string) (socmed models.SocialMedia, err error) {
	if socmed, err = socialMediaUseCase.socialMediaRepository.Update(ctx, socialMedia, id); err != nil {
		return socmed, err
	}

	return socmed, nil
}

func (socialMediaUseCase *socialMediaUseCase) Delete(ctx context.Context, id string) (err error) {
	if err = socialMediaUseCase.socialMediaRepository.Delete(ctx, id); err != nil {
		return err
	}

	return
}

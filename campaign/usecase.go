package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Usecase interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
}

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := u.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := u.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (u *usecase) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := u.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (u *usecase) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{
		Name:        input.Name,
		Summary:     input.Summary,
		Description: input.Description,
		GoalAmount:  input.GoalAmount,
		Perks:       input.Perks,
		UserID:      input.User.ID,
	}

	sluggedString := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	slug := slug.Make(sluggedString)

	campaign.Slug = slug

	newCampaign, err := u.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

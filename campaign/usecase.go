package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Usecase interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CampaignInput) (Campaign, error)
	Update(inputID GetCampaignDetailInput, inputData CampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
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

func (u *usecase) CreateCampaign(input CampaignInput) (Campaign, error) {
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

func (u *usecase) Update(inputID GetCampaignDetailInput, inputData CampaignInput) (Campaign, error) {
	campaign, err := u.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	campaign.Name = inputData.Name
	campaign.Summary = inputData.Summary
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	if inputID.ID != inputData.User.ID {
		err := errors.New("user not authorized")
		return campaign, err
	}

	updatedCampaign, err := u.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (u *usecase) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	var isPrimary = 0

	if input.IsPrimary {
		_, err := u.repository.MarkAllAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
		isPrimary = 1
	}

	campaignImage := CampaignImage{
		ID:        input.CampaignID,
		IsPrimary: isPrimary,
		FileName:  fileLocation,
	}

	newCampaignImage, err := u.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil

}

package campaign

import "strings"

type CampaignFormatter struct {
	ID            int    `json:"id"`
	UserID        int    `json:"userId"`
	Name          string `json:"name"`
	Summary       string `json:"summary"`
	ImageUrl      string `json:"imageUrl"`
	GoalAmount    int    `json:"goalAmount"`
	CurrentAmount int    `json:"currentAmount"`
	Slug          string `json:"slug"`
}

type CampaignDetailFormatter struct {
	ID            int                      `json:"id"`
	Name          string                   `json:"name"`
	Summary       string                   `json:"summary"`
	Description   string                   `json:"description"`
	ImageUrl      string                   `json:"imageUrl"`
	GoalAmount    int                      `json:"goalAmount"`
	CurrentAmount int                      `json:"currentAmount"`
	DonatorCount  int                      `json:"donatorCount"`
	UserID        int                      `json:"userId"`
	Slug          string                   `json:"slug"`
	Perks         []string                 `json:"perks"`
	User          CampaignUserFormatter    `json:"user"`
	Images        []CampaignImageFormatter `json:"image"`
}

type CampaignUserFormatter struct {
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

type CampaignImageFormatter struct {
	ImageUrl  string `json:"imageUrl"`
	IsPrimary bool   `json:"isPrimary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formattedCampaign := CampaignFormatter{
		ID:            campaign.ID,
		UserID:        campaign.UserID,
		Name:          campaign.Name,
		Summary:       campaign.Summary,
		ImageUrl:      "",
		GoalAmount:    campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		Slug:          campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		formattedCampaign.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return formattedCampaign
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{
		ID:            campaign.ID,
		UserID:        campaign.UserID,
		Name:          campaign.Name,
		Summary:       campaign.Summary,
		ImageUrl:      "",
		DonatorCount:  campaign.DonatorCount,
		GoalAmount:    campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		Slug:          campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	perks := strings.Split(campaign.Perks, ",")
	campaignDetailFormatter.Perks = perks

	UserFormatter := CampaignUserFormatter{
		Name:      campaign.User.Name,
		AvatarUrl: campaign.User.AvatarFileName,
	}

	var images []CampaignImageFormatter

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{
			ImageUrl: image.FileName,
		}

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}

		campaignImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignImageFormatter)
	}

	campaignDetailFormatter.User = UserFormatter
	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}

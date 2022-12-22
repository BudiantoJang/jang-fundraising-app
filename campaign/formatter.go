package campaign

type CampaignFormatter struct {
	ID             int    `json:"id"`
	UserID         int    `json:"userId"`
	Name           string `json:"name"`
	Summary        string `json:"summary"`
	ImageURL       string `json:"imageURL"`
	GoalAmmount    int    `json:"goalAmmount"`
	CurrentAmmount int    `json:"currentAmmount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formattedCampaign := CampaignFormatter{
		ID:             campaign.ID,
		UserID:         campaign.UserID,
		Name:           campaign.Name,
		Summary:        campaign.Summary,
		ImageURL:       "",
		GoalAmmount:    campaign.GoalAmmount,
		CurrentAmmount: campaign.CurrentAmmount,
	}

	if len(campaign.CampaignImages) > 0 {
		formattedCampaign.ImageURL = campaign.CampaignImages[0].FileName
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

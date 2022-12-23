package campaign

type Usecase interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
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

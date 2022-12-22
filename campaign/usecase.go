package campaign

type Usecase interface {
	GetCampaigns(userID int) ([]Campaign, error)
}

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{repository}
}

func (s *usecase) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

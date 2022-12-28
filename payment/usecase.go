package payment

import (
	"jangFundraising/campaign"
	"jangFundraising/user"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type Usecase interface {
	GetPaymentUrl(transaction Transaction, user user.User) (string, error)
}

type usecase struct {
	campaignRepository campaign.Repository
}

func NewUsecase(campaignRepository campaign.Repository) *usecase {
	return &usecase{campaignRepository}
}

func (u *usecase) GetPaymentUrl(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "" // INPUT MIDTRANS SERVER KEY
	midclient.ClientKey = "" // INPUT MIDTRANS CLIENT KEY
	midclient.APIEnvType = midtrans.Sandbox
	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, err
}

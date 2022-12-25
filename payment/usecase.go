package payment

import (
	"jangFundraising/user"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type Usecase interface {
	GetPaymentUrl(transaction Transaction, user user.User) (string, error)
}

type usecase struct {
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (u *usecase) GetPaymentUrl(transaction Transaction, user user.User) (string, error) {
	// secret, err := u.readMidtransSecret()
	// if err != nil {
	// 	return "", err
	// }
	midclient := midtrans.NewClient()
	midclient.ServerKey = ""
	midclient.ClientKey = ""
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

// func (u *usecase) readMidtransSecret() (MidtransSecret, error) {
// 	var secret MidtransSecret
// 	yamlFile, err := ioutil.ReadFile("secret.yaml")
// 	if err != nil {
// 		return MidtransSecret{}, err
// 	}

// 	data := make(map[interface{}]interface{})

// 	err = yaml.Unmarshal(yamlFile, &data)
// 	if err != nil {
// 		return MidtransSecret{}, err
// 	}

// 	secret.MerchantID = fmt.Sprint(data["merchantId"])
// 	secret.ClientKey = fmt.Sprint(data["clientKey"])
// 	secret.ServerKey = fmt.Sprint(data["serverKey"])

// 	return secret, nil

// }

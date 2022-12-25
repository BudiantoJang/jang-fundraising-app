package payment

type MidtransSecret struct {
	MerchantID string `yaml:"merchantId"`
	ClientKey  string `yaml:"clientKey"`
	ServerKey  string `yaml:"serverKey"`
}

type Transaction struct {
	ID     int
	Amount int
}

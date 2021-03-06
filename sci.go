package paykassasci

import (
	"strconv"
)

// SCI contains basic settings
type SCI struct {
	ID   int
	Key  string
	Test bool
}

// CheckPaymentResponse is wrapper response CheckPayment
type CheckPaymentResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Transaction string `json:"transaction"`
		ShopID      string `json:"shop_id"`
		OrderID     string `json:"order_id"`
		Amount      string `json:"amount"`
		Currency    string `json:"currency"`
		System      string `json:"system"`
		Address     string `json:"address"`
		Hash        string `json:"hash"`
		Partial     string `json:"partial"`
	} `json:"data"`
}

// CheckPayment is wrapper sci_confirm_order
// https://paykassa.pro/docs/#api-SCI-sci_confirm_order
func (s *SCI) CheckPayment(privateHash string) (CheckPaymentResponse, error) {

	var responseCheckPayment = &CheckPaymentResponse{}

	var param = s.getParamMap()
	param["func"] = ConfirmOrder
	param["private_hash"] = privateHash

	data, statusCode, err := sendRequest(param)
	if err != nil {
		return *responseCheckPayment, err
	}

	err = handlerResp(data, statusCode, responseCheckPayment)

	return *responseCheckPayment, err
}

// GetCryptocurrencyAddressForDepositResponse is wrapper response GetCryptocurrencyAddressForDeposit
type GetCryptocurrencyAddressForDepositResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Invoice  int    `json:"invoice"`
		OrderID  string `json:"order_id"`
		Wallet   string `json:"wallet"`
		Amount   string `json:"amount"`
		System   string `json:"system"`
		Currency string `json:"currency"`
		URL      string `json:"url"`
		Tag      bool   `json:"tag"`
	} `json:"data"`
}

// GetCryptocurrencyAddressForDeposit is wrapper sci_create_order_get_data
// https://paykassa.pro/docs/#api-SCI-sci_create_order_get_data
func (s *SCI) GetCryptocurrencyAddressForDeposit(orderID int, amount float64, currencyID int, comment string, phone bool, paidCommission string) (GetCryptocurrencyAddressForDepositResponse, error) {

	var responseGetCryptocurrency = &GetCryptocurrencyAddressForDepositResponse{}

	var param = s.getParamPayMap(orderID, amount, currencyID, comment, phone, paidCommission)
	param["func"] = CreateOrderGetData

	data, statusCode, err := sendRequest(param)
	if err != nil {
		return *responseGetCryptocurrency, err
	}

	err = handlerResp(data, statusCode, responseGetCryptocurrency)

	return *responseGetCryptocurrency, err
}

// GetLinkForDepositResponse is wrapper response GetLinkForDeposit
type GetLinkForDepositResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Invoice  int    `json:"invoice"`
		OrderID  string `json:"order_id"`
		Amount   string `json:"amount"`
		System   string `json:"system"`
		Currency string `json:"currency"`
		URL      string `json:"url"`
	} `json:"data"`
}

// GetLinkForDeposit is wrapper sci_create_order
// https://paykassa.pro/docs/#api-SCI-sci_create_order
func (s *SCI) GetLinkForDeposit(orderID int, amount float64, currencyID int, comment string, phone bool, paidCommission string) (GetLinkForDepositResponse, error) {

	var responseGetLinkForDeposit = &GetLinkForDepositResponse{}

	var param = s.getParamPayMap(orderID, amount, currencyID, comment, phone, paidCommission)
	param["func"] = CreateOrder

	data, statusCode, err := sendRequest(param)
	if err != nil {
		return *responseGetLinkForDeposit, err
	}

	err = handlerResp(data, statusCode, responseGetLinkForDeposit)

	return *responseGetLinkForDeposit, err
}

func (s *SCI) getParamPayMap(orderID int, amount float64, currencyID int, comment string, phone bool, paidCommission string) map[string]string {
	var param = s.getParamMap()
	param["order_id"] = strconv.Itoa(orderID)
	param["amount"] = strconv.FormatFloat(amount, 'E', -1, 64)
	param["currency"] = CurrencyCode(currencyID)
	param["system"] = strconv.Itoa(currencyID)
	param["comment"] = comment
	param["phone"] = strconv.FormatBool(phone)
	param["paid_commission"] = paidCommission

	return param
}

func (s *SCI) getParamMap() map[string]string {
	var param = make(map[string]string)
	param["sci_id"] = strconv.Itoa(s.ID)
	param["sci_key"] = s.Key
	param["test"] = strconv.FormatBool(s.Test)

	return param
}

// InitSCI initializes SCI
func InitSCI(merchantID int, merchantKey string, test bool) SCI {
	return SCI{
		ID:   merchantID,
		Key:  merchantKey,
		Test: test,
	}
}

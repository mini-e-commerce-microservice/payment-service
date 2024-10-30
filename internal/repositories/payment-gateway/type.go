package payment_gateway

import "github.com/guregu/null/v5"

type ChargeInput struct {
	PaymentType       string                              `json:"payment_type"`
	TransactionDetail ChargeInputTransactionDetail        `json:"transaction_details"`
	CustomExpiry      null.Value[ChargeInputCustomExpiry] `json:"custom_expiry,omitempty"`
	ItemDetails       null.Value[[]ChargeInputItemDetail] `json:"item_details,omitempty"`
	Gopay             null.Value[Gopay]                   `json:"gopay,omitempty"`
}

type ChargeCustomerAddress struct {
	FName       string `json:"first_name,omitempty"`
	LName       string `json:"last_name,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	Postcode    string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type ChargeInputItemDetail struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Qty          int32  `json:"quantity"`
	Brand        string `json:"brand,omitempty"`
	Category     string `json:"category,omitempty"`
	MerchantName string `json:"merchant_name,omitempty"`
}

type ChargeInputTransactionDetail struct {
	OrderID     string `json:"order_id"`
	GrossAmount int64  `json:"gross_amount"`
}

type ChargeInputCustomExpiry struct {
	ExpiryDuration int32  `json:"expiry_duration"`
	Unit           string `json:"unit"` // Unit by default minute
}

type ChargeOutput struct{}

type Gopay struct {
	EnableCallback     bool   `json:"enable_callback,omitempty"`
	CallbackUrl        string `json:"callback_url,omitempty"`
	AccountID          string `json:"account_id,omitempty"`
	PaymentOptionToken string `json:"payment_option_token,omitempty"`
	PreAuth            bool   `json:"pre_auth,omitempty"`
	Recurring          bool   `json:"recurring,omitempty"`
	PromotionIDs       string `json:"promotion_ids,omitempty"`
}

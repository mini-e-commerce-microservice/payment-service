package midtrans

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/go-midtrans-sdk"
	coreapi_midtrans "github.com/SyaibanAhmadRamadhan/go-midtrans-sdk/coreapi"
	"testing"
)

func TestChargeQRIS(t *testing.T) {
	serverKey := "SB-Mid-server-LR5SazLSkB3wdU2MRwIpkXFp"
	coreapi := coreapi_midtrans.NewAPI(coreapi_midtrans.ServerKey(serverKey))
	input := coreapi_midtrans.ChargeQRISInput{
		TransactionDetail: midtrans.TransactionDetail{
			OrderID:     "ORDER-123422s5s2", // Unique order ID
			GrossAmount: 20000,              // Total amount in IDR (e.g., 25,000 IDR)
		},
		ItemDetails: []midtrans.ItemDetail{
			{
				ID:           "ITEM1",      // Item ID or SKU
				Price:        10000,        // Price per item in IDR
				Qty:          2,            // Quantity of the item
				Name:         "Product A",  // Name of the item
				Brand:        "Brand A",    // Optional brand
				Category:     "Category X", // Optional category
				MerchantName: "Merchant Y", // Optional merchant name
			},
		},
		Acquirer: "gopay", // Optional acquirer, specific to your integration
		CustomerDetail: &midtrans.CustomerDetail{
			FirstName: "John", // Customer's first name
			LastName:  "Doe",  // Customer's last name
			Email:     "john@example.com",
			Phone:     "+62123456789",
			BillingAddress: &midtrans.CustomerBillingAddress{
				FirstName:   "John",
				LastName:    "Doe",
				Phone:       "+62123456789",
				Address:     "Jl. Sudirman No. 1",
				City:        "Jakarta",
				PostalCode:  "12345",
				CountryCode: "IDN", // Indonesia country code
			},
			ShippingAddress: &midtrans.CustomerShippingAddress{
				FirstName:   "John",
				LastName:    "Doe",
				Phone:       "+62123456789",
				Address:     "Jl. Thamrin No. 10",
				City:        "Jakarta",
				PostalCode:  "12345",
				CountryCode: "IDN",
			},
		},
		CustomExpiry: &midtrans.CustomExpiry{
			ExpiryDuration: 10,
			Unit:           "minute", // Unit can be "minute" or "hour"
		},
		MetaData: map[string]string{
			"note": "Payment for invoice #12345", // Example of metadata
		},
	}

	output, err := coreapi.ChargeQRIS(context.Background(), input)
	t.Log(err)
	t.Log(output.ErrorBadReqResponse)
	t.Log(output.ResponseSuccess)
}

func TestChargeGOPAY(t *testing.T) {
	serverKey := "SB-Mid-server-LR5SazLSkB3wdU2MRwIpkXFp"
	coreapi := coreapi_midtrans.NewAPI(coreapi_midtrans.ServerKey(serverKey))
	input := coreapi_midtrans.ChargeGoPayInput{
		TransactionDetail: midtrans.TransactionDetail{
			OrderID:     "ORDER-1s2sss3422252", // Unique order ID
			GrossAmount: 20000,                 // Total amount in IDR (e.g., 25,000 IDR)
		},
	}

	output, err := coreapi.ChargeGoPay(context.Background(), input)
	t.Log(err)
	t.Log(output.ErrorBadReqResponse)
	t.Log(output.ResponseSuccess)
	t.Log(output.ResponseSuccess.ActionDeepLinkRedirect)
	t.Log(output.ResponseSuccess.ActionGetStatus)
	t.Log(output.ResponseSuccess.ActionGenerateQRCode)
	t.Log(output.ResponseSuccess.ActionCancel)
}

func TestChargeShopeePay(t *testing.T) {
	serverKey := "SB-Mid-server-LR5SazLSkB3wdU2MRwIpkXFp"
	coreapi := coreapi_midtrans.NewAPI(coreapi_midtrans.ServerKey(serverKey))
	input := coreapi_midtrans.ChargeShopeePayInput{
		TransactionDetail: midtrans.TransactionDetail{
			OrderID:     "ORDER-1ss2ss342s2252", // Unique order ID
			GrossAmount: 20000,                  // Total amount in IDR (e.g., 25,000 IDR)
		},
		ShopeePay: &midtrans.ShopeePay{
			CallbackURL: "http://localhost:3000",
		},
	}

	output, err := coreapi.ChargeShopeePay(context.Background(), input)
	t.Log(err)
	t.Log(output.ErrorBadReqResponse)
	t.Log(output.ResponseSuccess)
}

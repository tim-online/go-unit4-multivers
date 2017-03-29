package multivers

import (
	"context"
	"fmt"
)

const (
	CustomerPath = "/api/%s/Customer/%s.json"
)

func NewCustomerService(client *Client) *CustomerService {
	return &CustomerService{Client: client}
}

type CustomerService struct {
	Client *Client
}

func (s *CustomerService) Get(database string, customerID string, ctx context.Context) (*CustomerGetResponse, error) {
	method := "GET"
	responseBody := NewCustomerGetResponse()
	path := fmt.Sprintf(CustomerPath, database, customerID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewCustomerGetResponse() *CustomerGetResponse {
	return &CustomerGetResponse{}
}

type CustomerGetResponse Customer

type CustomProperties struct {
	//TODO is the name correctly formatted?
	REL_DEB_BRANCHE          string               `json:"REL_DEB.BRANCHE"`
	REL_DEB_LICENTIENUMMER   string               `json:"REL_DEB.LICENTIENUMMER"`
	REL_DEB_OPTICROP         string               `json:"REL_DEB.OPTICROP"`
	REL_DEB_VERTEGENWOORDIGE string               `json:"REL_DEB.VERTEGENWOORDIGE"`
	REL_DEB_EMAILFACTUUR     string               `json:"REL_DEB.EMAILFACTUUR"`
	REL_DEB_EMAILDSB         string               `json:"REL_DEB.EMAILDSB"`
	REL_DEB_VERVALDATUM      *TimeWithoutTimeZone `json:"REL_DEB.VERVALDATUM"`
}

type Customer struct {
	AccountManagerID                    string               `json:"accountManagerId"`
	Addresses                           []Address            `json:"addresses"`
	Messages                            []Message            `json:"messages"`
	ApplyOrderSurcharge                 bool                 `json:"applyOrderSurcharge"`
	BusinessNumber                      string               `json:"businessNumber"`
	CanChange                           bool                 `json:"canChange"`
	ChargeVatTypeID                     int                  `json:"chargeVatTypeId"`
	City                                string               `json:"city"`
	CocCity                             string               `json:"cocCity"`
	CocDate                             *TimeWithoutTimeZone `json:"cocDate"`
	CocRegistration                     string               `json:"cocRegistration"`
	CollectiveInvoiceSystemID           string               `json:"collectiveInvoiceSystemId"`
	CombineInvoicesForElectronicBanking bool                 `json:"combineInvoicesForElectronicBanking"`
	CountryID                           string               `json:"countryId"`
	CreditLimit                         float64              `json:"creditLimit"`
	CreditSqueezeID                     string               `json:"creditSqueezeId"`
	CurrencyID                          string               `json:"currencyId"`
	CustomerID                          string               `json:"customerId"`
	CustomerStateID                     string               `json:"customerStateId"`
	Database                            string               `json:"database"`
	DateChanged                         DateNLNL             `json:"dateChanged"`
	DateCreated                         DateNLNL             `json:"dateCreated"`
	DeliveryConditionID                 string               `json:"deliveryConditionId"`
	DiscountPercentage                  float64              `json:"discountPercentage"`
	Email                               string               `json:"email"`
	Fax                                 string               `json:"fax"`
	FullAddress                         string               `json:"fullAddress"`
	FullDeliveryAddress                 string               `json:"fullDeliveryAddress"`
	GoogleMapsDirectionsUrl             string               `json:"fullDeliveryAddress"`
	GoogleMapsUrl                       string               `json:"googleMapsUrl"`
	HasOutstandingBalance               bool                 `json:"hasOutstandingBalance"`
	Homepage                            string               `json:"homepage"`
	IntratatGoodsCodeID                 int                  `json:"intrastatGoodsCodeId"`
	IntrastatGoodsDistributionID        int                  `json:"intrastatGoodsDistributionId"`
	IntrastatStatSystemID               int                  `json:"intrastatStatSystemId"`
	IntrastatTrafficRegionID            int                  `json:"intrastatTrafficRegionId"`
	IntrastatTransactionTypeID          string               `json:"intrastatTransactionTypeId"`
	IntrastatTransportTypeID            int                  `json:"intrastatTransportTypeId"`
	InvoiceOnBehalfOfMembers            bool                 `json:"invoiceOnBehalfOfMembers"`
	IsDunForPayment                     bool                 `json:"isDunForPayment"`
	IsInFactoring                       bool                 `json:"isInFactoring"`
	IsPaymentRefRequired                bool                 `json:"isPaymentRefRequired"`
	IsPurchaseOrganization              bool                 `json:"isPurchaseOrganization"`
	LanguageID                          string               `json:"languageId"`
	MobilePhone                         string               `json:"mobilePhone"`
	Name                                string               `json:"name"`
	OrganizationID                      int                  `json:"organizationId"`
	PaymentConditionID                  string               `json:"paymentConditionId"`
	Person                              string               `json:"person"`
	PricelistID                         string               `json:"pricelistId"`
	PrintPurchaseDetails                bool                 `json:"printPurchaseDetails"`
	PurchaseOrganizationID              string               `json:"purchaseOrganizationId"`
	PurchaseOrganizationMemberID        string               `json:"purchaseOrganizationMemberId"`
	RevenueAccountID                    string               `json:"revenueAccountId"`
	ShortName                           string               `json:"shortName"`
	Street1                             string               `json:"street1"`
	Street2                             string               `json:"street2"`
	SupplierID                          string               `json:"supplierId"`
	Telephone                           string               `json:"telephone"`
	UsesUBLInvoice                      bool                 `json:"usesUBLInvoice"`
	VatNumber                           string               `json:"vatNumber"`
	VatScenarioID                       int                  `json:"vatScenarioId"`
	VatVerificationDate                 *TimeWithoutTimeZone `json:"vatVerificationDate"`
	ZipCode                             string               `json:"zipCode"`
	CustomProperties                    CustomProperties     `json:"customProperties"`
}

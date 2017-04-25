package multivers

import (
	"context"
	"fmt"
)

const (
	CompanyDetailsInfoPath = "/api/%s/CompanyDetailsInfo.json"
)

func NewCompanyDetailsInfoService(client *Client) *CompanyDetailsInfoService {
	return &CompanyDetailsInfoService{Client: client}
}

type CompanyDetailsInfoService struct {
	Client *Client
}

func (s *CompanyDetailsInfoService) Get(database string, ctx context.Context) (*CompanyDetailsInfoGetResponse, error) {
	method := "GET"
	responseBody := NewCompanyDetailsInfoGetResponse()
	path := fmt.Sprintf(CompanyDetailsInfoPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewCompanyDetailsInfoGetResponse() *CompanyDetailsInfoGetResponse {
	return &CompanyDetailsInfoGetResponse{}
}

type CompanyDetailsInfoGetResponse CompanyDetailsInfo

type CompanyDetailsInfo struct {
	AdministrationDescription        string  `json:"administrationDescription"`
	ApprovePurchaseOrders            bool    `json:"approvePurchaseOrders"`
	ApproveQuotations                bool    `json:"approveQuotations"`
	ApproveSalesOrders               bool    `json:"approveSalesOrders"`
	ApproveSupplierInvoices          bool    `json:"approveSupplierInvoices"`
	AuditionCategory                 int     `json:"auditionCategory"`
	AutoNumberCustomerID             bool    `json:"autoNumberCustomerId"`
	AutoNumberSupplierID             bool    `json:"autoNumberSupplierId"`
	B2bMandateTypeApplicable         bool    `json:"b2bMandateTypeApplicable"`
	BlockUnmatchedSupplierInvoice    bool    `json:"blockUnmatchedSupplierInvoice"`
	CheckMinimumTotalHours           bool    `json:"checkMinimumTotalHours"`
	City                             string  `json:"city"`
	CocRegistration                  string  `json:"cocRegistration"`
	CoreMandateTypeApplicable        bool    `json:"coreMandateTypeApplicable"`
	CostSpecPerOrderLine             bool    `json:"costSpecPerOrderLine"`
	CreditSqueezeOnCreditNotes       bool    `json:"creditSqueezeOnCreditNotes"`
	CurrencyDecimals                 float64 `json:"currencyDecimals"`
	CurrencyID                       string  `json:"currencyId"`
	CustomerDiscountIsOrderDiscount  bool    `json:"customerDiscountIsOrderDiscount"`
	CustomerIDLength                 int     `json:"customerIdLength"`
	CustomerIDPrecededByZeroes       bool    `json:"customerIdPrecededByZeroes"`
	Email                            string  `json:"email"`
	EmailAccountType                 int     `json:"emailAccountType"`
	IctListing                       bool    `json:"ictListing"`
	IsAuditionApplicable             bool    `json:"isAuditionApplicable"`
	JournalizeCostPriceOnTimesheet   bool    `json:"journalizeCostPriceOnTimesheet"`
	LanguageID                       string  `json:"languageId"`
	Legislation                      string  `json:"legislation"`
	LowestPrice                      bool    `json:"lowestPrice"`
	MatchSupplierInvoice             bool    `json:"matchSupplierInvoice"`
	MaximumPaymentDifference         float64 `json:"maximumPaymentDifference"`
	NextCustomerID                   int     `json:"nextCustomerId"`
	NextSupplierID                   int     `json:"nextSupplierId"`
	PersonIDsOfCurrentUser           []int   `json:"personIdsOfCurrentUser"`
	Phone                            string  `json:"phone"`
	PrintMandateFormApplicable       bool    `json:"printMandateFormApplicable"`
	RestrictedSubAdmins              bool    `json:"restrictedSubAdmins"`
	SpecifyQuantity                  bool    `json:"specifyQuantity"`
	StockLevelBasedOnQuantityOrdered bool    `json:"stockLevelBasedOnQuantityOrdered"`
	StockValuation                   string  `json:"stockValuation"`
	Street                           string  `json:"street"`
	Street2                          string  `json:"street2"`
	SupplierIDLength                 int     `json:"supplierIdLength"`
	SupplierIDPrecededByZeroes       bool    `json:"supplierIdPrecededByZeroes"`
	ValidateProductEanNumber         bool    `json:"validateProductEanNumber"`
	VatDeclarationIsPeriod           bool    `json:"vatDeclarationIsPeriod"`
	VatEnabled                       bool    `json:"vatEnabled"`
	VatNumber                        string  `json:"vatNumber"`
	VatOnCreditSqueeze               bool    `json:"vatOnCreditSqueeze"`
	VatPerJournalLine                string  `json:"vatPerJournalLine"`
	VersionNumberPlatinum            int     `json:"versionNumberPlatinum"`
	ZipCode                          string  `json:"zipCode"`
}

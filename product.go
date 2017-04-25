package multivers

import (
	"context"
	"fmt"
)

const (
	ProductPath = "/api/%s/Product/%s.json"
)

func NewProductService(client *Client) *ProductService {
	return &ProductService{Client: client}
}

type ProductService struct {
	Client *Client
}

func (s *ProductService) Get(database string, productID string, ctx context.Context) (*ProductGetResponse, error) {
	method := "GET"
	responseBody := NewProductGetResponse()
	path := fmt.Sprintf(ProductPath, database, productID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewProductGetResponse() *ProductGetResponse {
	return &ProductGetResponse{}
}

type ProductGetResponse Product

type Product struct {
	AccountID                string            `json:"accountId"`
	AdditionalUnits          float64           `json:"additionalUnits"`
	Messages                 []Message         `json:"messages"`
	CanChange                bool              `json:"canChange"`
	Components               []string          `json:"components"`
	ConstellationNotition    string            `json:"constellationNotition"`
	CoverageAccountID        string            `json:"coverageAccountId"`
	CurrentStock             float64           `json:"currentStock"`
	DateCreated              DateNLNL          `json:"dateCreated"`
	Description              string            `json:"description"`
	DiscountAccountID        string            `json:"discountAccountId"`
	EanCode                  string            `json:"eanCode"`
	IntrastatCode            int               `json:"intrastatCode"`
	LastUpdate               DateNLNL          `json:"lastUpdate"`
	Margin                   float64           `json:"margin"`
	MaxStock                 float64           `json:"maxStock"`
	MinStock                 float64           `json:"minStock"`
	PreferredProductSupplier ProductSupplier   `json:"preferredProductSupplier"`
	PricePer                 float64           `json:"pricePer"`
	PriceVatExcl             float64           `json:"priceVatExcl"`
	PriceVatIncl             float64           `json:"priceVatIncl"`
	PrintOnInvoice           bool              `json:"printOnInvoice"`
	PrintOnQuote             bool              `json:"printOnQuote"`
	PrintOnShippingOrder     bool              `json:"printOnShippingOrder"`
	PrintOnWarehouseOrder    bool              `json:"printOnWarehouseOrder"`
	ProductDescriptions      []string          `json:"productDescriptions"`
	ProductDiscountGroupId   string            `json:"productDiscountGroupId"`
	ProductGroupId           string            `json:"productGroupId"`
	ProductId                string            `json:"productId"`
	ProductState             int               `json:"productState"`
	ProductSuppliers         []ProductSupplier `json:"productSuppliers"`
	ProductType              int               `json:"productType"`
	ProjectEntryType         int               `json:"projectEntryType"`
	ProjectSurchargePerc     float64           `json:"projectSurchargePerc"`
	PublishOnWeb             bool              `json:"publishOnWeb"`
	PurchaseOrderCount       int               `json:"purchaseOrderCount"`
	QuantityScale            int               `json:"quantityScale"`
	ShortName                string            `json:"shortName"`
	Stocks                   []Stock           `json:"stocks"`
	StockTransactionCount    int               `json:"stockTransactionCount"`
	StockTransferPrice       float64           `json:"stockTransferPrice"`
	Unit                     string            `json:"unit"`
	VatCodeId                int               `json:"vatCodeId"`
	Weight                   float64           `json:"weight"`
	CustomProperties         CustomProperties  `json:"customProperties"`
}

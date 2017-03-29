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

//TODO is the name correctly formatted?
type CustomPropertiesProduct struct {
	ARTIKEL_VERPAKKING      string  `json:"ARTIKEL.VERPAKKING"`
	ARTIKEL_OMVERPAKKING    string  `json:"ARTIKEL.OMVERPAKKING"`
	ARTIKEL_OMVPAANTAL      float64 `json:"ARTIKEL.OMVPAANTAL"`
	ARTIKEL_NUMINVERPAKKING float64 `json:"ARTIKEL.NUMINVERPAKKING"`
	ARTIKEL_GEVARENKLASSE   string  `json:"ARTIKEL.GEVARENKLASSE"`
	ARTIKEL_TOELATING       string  `json:"ARTIKEL.TOELATING"`
	ARTIKEL_UNNR            string  `json:"ARTIKEL.UNNR"`
	ARTIKEL_STOFNAAM        string  `json:"ARTIKEL.STOFNAAM"`
	ARTIKEL_KAARTNR         string  `json:"ARTIKEL.KAARTNR"`
	ARTIKEL_MENGBAAR        string  `json:"ARTIKEL.MENGBAAR"`
	ARTIKEL_KLEUR           string  `json:"ARTIKEL.KLEUR"`
	ARTIKEL_LOODS           string  `json:"ARTIKEL.LOODS"`
	ARTIKEL_KLASSE          string  `json:"ARTIKEL.KLASSE"`
	ARTIKEL_FLEVO           string  `json:"ARTIKEL.FLEVO"`
	ARTIKEL_STIKSTOF        string  `json:"ARTIKEL.STIKSTOF"`
	ARTIKEL_FOSFAAT         string  `json:"ARTIKEL.FOSFAAT"`
	ARTIKEL_DREMPEL         float64 `json:"ARTIKEL.DREMPEL"`
	ARTIKEL_LINK            string  `json:"ARTIKEL.LINK"`
	ARTIKEL_CATEGORIE       string  `json:"ARTIKEL.CATEGORIE"`
	ARTIKEL_TOELAT          string  `json:"ARTIKEL.TOELAT"`
}

type Stock struct {
	Messages         []Message `json:"messages"`
	CanChange        bool      `json:"canChange"`
	Description      string    `json:"description"`
	Location         string    `json:"location"`
	MaximumStock     float64   `json:"maximumStock"`
	MinimumStock     float64   `json:"minimumStock"`
	QuantityOrdered  float64   `json:"quantityOrdered"`
	QuantityReserved float64   `json:"quantityReserved"`
	TechnicalStock   float64   `json:"technicalStock"`
	WarehouseID      string    `json:"warehouseId"`
}

type ProductSuppliers struct {
	Messages                    []Message `json:"messages"`
	CanChange                   bool      `json:"canChange"`
	CurrencyID                  string    `json:"currencyId"`
	DiscountPercentage          float64   `json:"discountPercentage"`
	MinOrderQuantity            float64   `json:"minOrderQuantity"`
	PartNumber                  string    `json:"partNumber"`
	Preferred                   bool      `json:"preferred"`
	Price                       float64   `json:"price"`
	productSupplierID           string    `json:"productSupplierId"`
	PurchaseSurchargePercentage float64   `json:"purchaseSurchargePercentage"`
}

type Product struct {
	AccountID                string                  `json:"accountId"`
	AdditionalUnits          float64                 `json:"additionalUnits"`
	Messages                 []Message               `json:"messages"`
	CanChange                bool                    `json:"canChange"`
	Components               []string                `json:"components"`
	ConstellationNotition    string                  `json:"constellationNotition"`
	CoverageAccountID        string                  `json:"coverageAccountId"`
	CurrentStock             float64                 `json:"currentStock"`
	DateCreated              DateNLNL                `json:"dateCreated"`
	Description              string                  `json:"description"`
	DiscountAccountID        string                  `json:"discountAccountId"`
	EanCode                  string                  `json:"eanCode"`
	IntrastatCode            int                     `json:"intrastatCode"`
	LastUpdate               DateNLNL                `json:"lastUpdate"`
	Margin                   float64                 `json:"margin"`
	MaxStock                 float64                 `json:"maxStock"`
	MinStock                 float64                 `json:"minStock"`
	PreferredProductSupplier int                     `json:"preferredProductSupplier"`
	PricePer                 float64                 `json:"pricePer"`
	PriceVatExcl             float64                 `json:"priceVatExcl"`
	PriceVatIncl             float64                 `json:"priceVatIncl"`
	PrintOnInvoice           bool                    `json:"printOnInvoice"`
	PrintOnQuote             bool                    `json:"printOnQuote"`
	PrintOnShippingOrder     bool                    `json:"printOnShippingOrder"`
	PrintOnWarehouseOrder    bool                    `json:"printOnWarehouseOrder"`
	ProductDescriptions      []string                `json:"productDescriptions"`
	ProductDiscountGroupId   string                  `json:"productDiscountGroupId"`
	ProductGroupId           string                  `json:"productGroupId"`
	ProductId                string                  `json:"productId"`
	ProductState             int                     `json:"productState"`
	ProductSuppliers         []ProductSuppliers      `json:"productSuppliers"`
	ProductType              int                     `json:"productType"`
	ProjectEntryType         int                     `json:"projectEntryType"`
	ProjectSurchargePerc     float64                 `json:"projectSurchargePerc"`
	PublishOnWeb             bool                    `json:"publishOnWeb"`
	PurchaseOrderCount       int                     `json:"purchaseOrderCount"`
	QuantityScale            int                     `json:"quantityScale"`
	ShortName                string                  `json:"shortName"`
	Stocks                   []Stock                 `json:"stocks"`
	StockTransactionCount    int                     `json:"stockTransactionCount"`
	StockTransferPrice       float64                 `json:"stockTransferPrice"`
	Unit                     string                  `json:"unit"`
	VatCodeId                int                     `json:"vatCodeId"`
	Weight                   float64                 `json:"weight"`
	CustomProperties         CustomPropertiesProduct `json:"customProperties"`
}

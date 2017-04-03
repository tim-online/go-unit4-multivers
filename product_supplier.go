package multivers

type ProductSupplier struct {
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
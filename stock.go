package multivers

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
package multivers

import (
	"context"
	"fmt"
)

const (
	WarehousePath = "/api/%s/Warehouse/%s"
)

func NewWarehouseService(client *Client) *WarehouseService {
	return &WarehouseService{Client: client}
}

type WarehouseService struct {
	Client *Client
}

func (s *WarehouseService) Get(database string, warehouseID string, ctx context.Context) (*WarehouseGetResponse, error) {
	method := "GET"
	responseBody := NewWarehouseGetResponse()
	path := fmt.Sprintf(WarehousePath, database, warehouseID+".json")

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *WarehouseService) GetPreferred(database string, ctx context.Context) (*WarehouseGetResponse, error) {
	method := "GET"
	responseBody := NewWarehouseGetResponse()
	path := fmt.Sprintf(WarehousePath, database, "Preferred.json")

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewWarehouseGetResponse() *WarehouseGetResponse {
	return &WarehouseGetResponse{}
}

type WarehouseGetResponse Warehouse

type Warehouse struct {
	Messages    []Message `json:"messages"`
	CanChange   bool      `json:"canChange"`
	Description string    `json:"description"`
	PreferredID string    `json:"preferredId"`
	WarehouseID string    `json:"warehouseId"`
}

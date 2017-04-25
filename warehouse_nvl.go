package multivers

import (
	"context"
	"fmt"
)

const (
	WarehouseNVLPath = "/api/%s/WarehouseNVL.json"
)

func NewWarehouseNVLService(client *Client) *WarehouseNVLService {
	return &WarehouseNVLService{Client: client}
}

type WarehouseNVLService struct {
	Client *Client
}

func (s *WarehouseNVLService) Get(database string, ctx context.Context) (*WarehouseNVLGetResponse, error) {
	method := "GET"
	responseBody := NewWarehouseNVLGetResponse()
	path := fmt.Sprintf(WarehouseNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewWarehouseNVLGetResponse() *WarehouseNVLGetResponse {
	return &WarehouseNVLGetResponse{}
}

type WarehouseNVLGetResponse []NVL
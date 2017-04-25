package multivers

import (
	"context"
	"fmt"
)

const (
	OrderChargeVatTypeNVLPath = "/api/%s/OrderChargeVatTypeNVL.json"
)

func NewOrderChargeVatTypeNVLService(client *Client) *OrderChargeVatTypeNVLService {
	return &OrderChargeVatTypeNVLService{Client: client}
}

type OrderChargeVatTypeNVLService struct {
	Client *Client
}

func (s *OrderChargeVatTypeNVLService) Get(database string, ctx context.Context) (*OrderChargeVatTypeNVLGetResponse, error) {
	method := "GET"
	responseBody := NewOrderChargeVatTypeNVLGetResponse()
	path := fmt.Sprintf(OrderChargeVatTypeNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewOrderChargeVatTypeNVLGetResponse() *OrderChargeVatTypeNVLGetResponse {
	return &OrderChargeVatTypeNVLGetResponse{}
}

type OrderChargeVatTypeNVLGetResponse []NVL

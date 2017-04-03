package multivers

import (
	"context"
	"fmt"
)

const (
	OrderTypeNVLPath = "/api/%s/OrderTypeNVL.json"
)

func NewOrderTypeNVLService(client *Client) *OrderTypeNVLService {
	return &OrderTypeNVLService{Client: client}
}

type OrderTypeNVLService struct {
	Client *Client
}

func (s *OrderTypeNVLService) Get(database string, ctx context.Context) (*OrderTypeNVLGetResponse, error) {
	method := "GET"
	responseBody := NewOrderTypeNVLGetResponse()
	path := fmt.Sprintf(OrderTypeNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewOrderTypeNVLGetResponse() *OrderTypeNVLGetResponse {
	return &OrderTypeNVLGetResponse{}
}

type OrderTypeNVLGetResponse []NVL
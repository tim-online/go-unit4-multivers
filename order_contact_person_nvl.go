package multivers

import (
	"context"
	"fmt"
)

const (
	OrderContactPersonNVLPath = "/api/%s/OrderContactPersonNVL/%s.json"
)

func NewOrderContactPersonNVLService(client *Client) *OrderContactPersonNVLService {
	return &OrderContactPersonNVLService{Client: client}
}

type OrderContactPersonNVLService struct {
	Client *Client
}

func (s *OrderContactPersonNVLService) Get(database string, customerID string, ctx context.Context) (*OrderContactPersonNVLGetResponse, error) {
	method := "GET"
	responseBody := NewOrderContactPersonNVLGetResponse()
	path := fmt.Sprintf(OrderContactPersonNVLPath, database, customerID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewOrderContactPersonNVLGetResponse() *OrderContactPersonNVLGetResponse {
	return &OrderContactPersonNVLGetResponse{}
}

type OrderContactPersonNVLGetResponse []NVL

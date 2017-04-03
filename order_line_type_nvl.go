package multivers

import (
	"context"
	"fmt"
)

const (
	OrderLineTypeNVLPath = "/api/%s/OrderLineTypeNVL.json"
)

func NewOrderLineTypeNVLService(client *Client) *OrderLineTypeNVLService {
	return &OrderLineTypeNVLService{Client: client}
}

type OrderLineTypeNVLService struct {
	Client *Client
}

func (s *OrderLineTypeNVLService) Get(database string, ctx context.Context) (*OrderLineTypeNVLGetResponse, error) {
	method := "GET"
	responseBody := NewOrderLineTypeNVLGetResponse()
	path := fmt.Sprintf(OrderLineTypeNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewOrderLineTypeNVLGetResponse() *OrderLineTypeNVLGetResponse {
	return &OrderLineTypeNVLGetResponse{}
}

type OrderLineTypeNVLGetResponse []NVL

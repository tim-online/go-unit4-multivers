package multivers

import (
	"context"
	"fmt"
)

const (
	OrderStateNVLPath = "/api/%s/OrderStateNVL.json"
)

func NewOrderStateNVLService(client *Client) *OrderStateNVLService {
	return &OrderStateNVLService{Client: client}
}

type OrderStateNVLService struct {
	Client *Client
}

func (s *OrderStateNVLService) Get(database string, ctx context.Context) (*OrderStateNVLGetResponse, error) {
	method := "GET"
	responseBody := NewOrderStateNVLGetResponse()
	path := fmt.Sprintf(OrderStateNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewOrderStateNVLGetResponse() *OrderStateNVLGetResponse {
	return &OrderStateNVLGetResponse{}
}

type OrderStateNVLGetResponse []OrderStateNVL

type OrderStateNVL struct {
	Name  int    `json:"name"`
	Value string `json:"value"`
}

package multivers

import (
	"context"
	"fmt"
)

const (
	CustomerPersonNVLPath = "/api/%s/CustomerPersonNVL/%s.json"
)

func NewCustomerPersonNVLService(client *Client) *CustomerPersonNVLService {
	return &CustomerPersonNVLService{Client: client}
}

type CustomerPersonNVLService struct {
	Client *Client
}

func (s *CustomerPersonNVLService) Get(database string, customerID string, ctx context.Context) (*CustomerPersonNVLGetResponse, error) {
	method := "GET"
	responseBody := NewCustomerPersonNVLGetResponse()
	path := fmt.Sprintf(CustomerPersonNVLPath, database, customerID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewCustomerPersonNVLGetResponse() *CustomerPersonNVLGetResponse {
	return &CustomerPersonNVLGetResponse{}
}

type CustomerPersonNVLGetResponse []NVL
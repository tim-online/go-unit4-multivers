package multivers

import (
	"context"
	"fmt"
)

const (
	CustomerNVLPath = "/api/%s/CustomerNVL.json"
)

func NewCustomerNVLService(client *Client) *CustomerNVLService {
	return &CustomerNVLService{Client: client}
}

type CustomerNVLService struct {
	Client *Client
}

func (s *CustomerNVLService) Get(database string, ctx context.Context) (*CustomerNVLGetResponse, error) {
	method := "GET"
	responseBody := NewCustomerNVLGetResponse()
	path := fmt.Sprintf(CustomerNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewCustomerNVLGetResponse() *CustomerNVLGetResponse {
	return &CustomerNVLGetResponse{}
}

type CustomerNVLGetResponse []NVL
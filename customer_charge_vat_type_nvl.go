package multivers

import (
	"context"
	"fmt"
)

const (
	CustomerChargeVatTypeNVLPath = "/api/%s/CustomerChargeVatTypeNVL.json"
)

func NewCustomerChargeVatTypeNVLService(client *Client) *CustomerChargeVatTypeNVLService {
	return &CustomerChargeVatTypeNVLService{Client: client}
}

type CustomerChargeVatTypeNVLService struct {
	Client *Client
}

func (s *CustomerChargeVatTypeNVLService) Get(database string, ctx context.Context) (*CustomerChargeVatTypeNVLGetResponse, error) {
	method := "GET"
	responseBody := NewCustomerChargeVatTypeNVLGetResponse()
	path := fmt.Sprintf(CustomerChargeVatTypeNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewCustomerChargeVatTypeNVLGetResponse() *CustomerChargeVatTypeNVLGetResponse {
	return &CustomerChargeVatTypeNVLGetResponse{}
}

type CustomerChargeVatTypeNVLGetResponse []NVL
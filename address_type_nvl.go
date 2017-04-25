package multivers

import (
	"context"
	"fmt"
)

const (
	AddressTypeNVLPath = "/api/%s/AddressTypeNVL.json"
)

func NewAddressTypeNVLService(client *Client) *AddressTypeNVLService {
	return &AddressTypeNVLService{Client: client}
}

type AddressTypeNVLService struct {
	Client *Client
}

func (s *AddressTypeNVLService) Get(database string, ctx context.Context) (*AddressTypeNVLGetResponse, error) {
	method := "GET"
	responseBody := NewAddressTypeNVLGetResponse()
	path := fmt.Sprintf(AddressTypeNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewAddressTypeNVLGetResponse() *AddressTypeNVLGetResponse {
	return &AddressTypeNVLGetResponse{}
}

type AddressTypeNVLGetResponse []NVL
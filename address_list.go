package multivers

import (
	"context"
	"fmt"
)

const (
	AddressListPath = "/api/%s/AddressList/%d.json"
)

func NewAddressListService(client *Client) *AddressListService {
	return &AddressListService{Client: client}
}

type AddressListService struct {
	Client *Client
}

func (s *AddressListService) Get(database string, organizationID int, ctx context.Context) (*AddressListGetResponse, error) {
	method := "GET"
	responseBody := NewAddressListGetResponse()
	path := fmt.Sprintf(AddressListPath, database, organizationID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewAddressListGetResponse() *AddressListGetResponse {
	return &AddressListGetResponse{}
}

type AddressListGetResponse []Address
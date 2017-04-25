package multivers

// https://api.online.unit4.nl/V14/Help/Api/GET-api-database-ProductInfoList_productId_shortName_description_productGroupId

import (
	"context"
	"fmt"
)

const (
	AddressInfoListPath = "/api/%s/AddressInfoList/%d.json"
)

func NewAddressInfoListService(client *Client) *AddressInfoListService {
	return &AddressInfoListService{Client: client}
}

type AddressInfoListService struct {
	Client *Client
}

func (s *AddressInfoListService) Get(database string, organizationID int, ctx context.Context) (*AddressInfoListGetResponse, error) {
	method := "GET"
	responseBody := NewAddressInfoListGetResponse()
	path := fmt.Sprintf(AddressInfoListPath, database, organizationID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}


func NewAddressInfoListGetResponse() *AddressInfoListGetResponse {
	return &AddressInfoListGetResponse{}
}

type AddressInfoListGetResponse []AddressInfo
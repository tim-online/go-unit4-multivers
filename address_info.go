package multivers

import (
	"context"
	"fmt"
)

const (
	AddressInfoPath = "/api/%s/AddressInfo/%s.json"
)

func NewAddressInfoService(client *Client) *AddressInfoService {
	return &AddressInfoService{Client: client}
}

type AddressInfoService struct {
	Client *Client
}

func (s *AddressInfoService) Get(database string, addressGUID string, ctx context.Context) (*AddressInfoGetResponse, error) {
	method := "GET"
	responseBody := NewAddressInfoGetResponse()
	path := fmt.Sprintf(AddressInfoPath, database, addressGUID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewAddressInfoGetResponse() *AddressInfoGetResponse {
	return &AddressInfoGetResponse{}
}

type AddressInfoGetResponse AddressInfo

type AddressInfo struct {
	AddressGUID             string `json:"addressGuid"`
	AddressID               int    `json:"addressId"`
	AddressTypeID           string `json:"addressTypeId"`
	City                    string `json:"city"`
	Country                 string `json:"country"`
	FullAddress             string `json:"fullAddress"`
	GoogleMapsDirectionsUrl string `json:"googleMapsDirectionsUrl"`
	GoogleMapsUrl           string `json:"googleMapsUrl"`
	Name                    string `json:"name"`
	Street1                 string `json:"street1"`
	Street2                 string `json:"street2"`
	ZipCode                 string `json:"zipCode"`
}

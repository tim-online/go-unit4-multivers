package multivers

import (
	"context"
	"fmt"
)

const (
	AddressPath = "/api/%s/Address/%s.json"
)

func NewAddressService(client *Client) *AddressService {
	return &AddressService{Client: client}
}

type AddressService struct {
	Client *Client
}

func (s *AddressService) Get(database string, addressGUID string, ctx context.Context) (*AddressGetResponse, error) {
	method := "GET"
	responseBody := NewAddressGetResponse()
	path := fmt.Sprintf(AddressPath, database, addressGUID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewAddressGetResponse() *AddressGetResponse {
	return &AddressGetResponse{}
}

type AddressGetResponse Address

type Address struct {
	AddressesGuID  string   `json:"addressGuid"`
	AddressID      int      `json:"addressId"`
	AddressType    int      `json:"addressType"`
	Messages       []string `json:"messages"`
	CanChange      bool     `json:"canChange"`
	City           string   `json:"city"`
	ContactPerson  string   `json:"contactPerson"`
	CountryId      string   `json:"countryId"`
	Email          string   `json:"email"`
	Fax            string   `json:"fax"`
	FullAddress    string   `json:"fullAddress"`
	LanguageId     string   `json:"languageId"`
	Name           string   `json:"name"`
	OrganizationId int      `json:"organizationId"`
	Street1        string   `json:"street1"`
	Street2        string   `json:"street2"`
	Telephone      string   `json:"telephone"`
	ZipCode        string   `json:"zipCode"`
}
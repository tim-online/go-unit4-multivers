package multivers

import (
	"context"
	"fmt"
)

const (
	OrganizationPath = "/api/%s/Organization/%d.json"
)

func NewOrganizationService(client *Client) *OrganizationService {
	return &OrganizationService{Client: client}
}

type OrganizationService struct {
	Client *Client
}

func (s *OrganizationService) Get(database string, organizationID int, ctx context.Context) (*OrganizationGetResponse, error) {
	method := "GET"
	responseBody := NewOrganizationGetResponse()
	path := fmt.Sprintf(OrganizationPath, database, organizationID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewOrganizationGetResponse() *OrganizationGetResponse {
	return &OrganizationGetResponse{}
}

type OrganizationGetResponse Organization

type Organization struct {
	AccountManagerID        string               `json:"accountManagerId"`
	Addresses               []Address            `json:"addresses"`
	Messages                []Message            `json:"messages"`
	BusinessNumber          string               `json:"businessNumber"`
	CanChange               bool                 `json:"canChange"`
	City                    string               `json:"city"`
	CocCity                 string               `json:"cocCity"`
	CocDate                 *TimeWithoutTimeZone `json:"cocDate"`
	CocRegistration         string               `json:"cocRegistration"`
	CountryID               string               `json:"countryId"`
	CurrencyID              string               `json:"currencyId"`
	CustomerID              string               `json:"customerId"`
	DateChanged             DateNLNL             `json:"dateChanged"`
	DateCreated             DateNLNL             `json:"dateCreated"`
	Email                   string               `json:"email"`
	Fax                     string               `json:"fax"`
	FullAddress             string               `json:"fullAddress"`
	FullDeliveryAddress     string               `json:"fullDeliveryAddress"`
	GoogleMapsDirectionsUrl string               `json:"googleMapsDirectionsUrl"`
	GoogleMapsUrl           string               `json:"googleMapsUrl"`
	Homepage                string               `json:"homepage"`
	LanguageID              string               `json:"languageId"`
	MobilePhone             string               `json:"mobilePhone"`
	Name                    string               `json:"name"`
	OrganizationID          int                  `json:"organizationId"`
	ShortName               string               `json:"shortName"`
	Street1                 string               `json:"street1"`
	Street2                 string               `json:"street2"`
	SupplierID              string               `json:"supplierId"`
	Telephone               string               `json:"telephone"`
	VatNumber               string               `json:"vatNumber"`
	VatVerificationDate     *TimeWithoutTimeZone `json:"vatVerificationDate"`
	ZipCode                 string               `json:"zipCode"`
}

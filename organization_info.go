package multivers

import (
	"context"
	"fmt"
)

const (
	OrganizationInfoPath = "/api/%s/OrganizationInfo/%d.json"
)

func NewOrganizationInfoService(client *Client) *OrganizationInfoService {
	return &OrganizationInfoService{Client: client}
}

type OrganizationInfoService struct {
	Client *Client
}

func (s *OrganizationInfoService) Get(database string, organizationID int, ctx context.Context) (*OrganizationInfoGetResponse, error) {
	method := "GET"
	responseBody := NewOrganizationInfoGetResponse()
	path := fmt.Sprintf(OrganizationInfoPath, database, organizationID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewOrganizationInfoGetResponse() *OrganizationInfoGetResponse {
	return &OrganizationInfoGetResponse{}
}

type OrganizationInfoGetResponse OrganizationInfo


type OrganizationInfo struct {
	BusinessNumber      string               `json:"businessNumber"`
	City                string               `json:"city"`
	CocCity             string               `json:"cocCity"`
	CocDate             *TimeWithoutTimeZone `json:"cocDate"`
	CocRegistration     string               `json:"cocRegistration"`
	CountryID           string               `json:"countryId"`
	CurrencyID          string               `json:"currencyId"`
	DateChanged         *TimeWithoutTimeZone `json:"dateChanged"`
	DateCreated         *TimeWithoutTimeZone `json:"dateCreated"`
	Email               string               `json:"email"`
	Fax                 string               `json:"fax"`
	Homepage            string               `json:"homepage"`
	LanguageID          string               `json:"languageId"`
	MobilePhone         string               `json:"mobilePhone"`
	Name                string               `json:"name"`
	OrganizationID      int                  `json:"organizationId"`
	ShortName           string               `json:"shortName"`
	Street1             string               `json:"street1"`
	Street2             string               `json:"street2"`
	Telephone           string               `json:"telephone"`
	VatNumber           string               `json:"vatNumber"`
	VatVerificationDate *TimeWithoutTimeZone `json:"vatVerificationDate"`
	ZipCode             string               `json:"zipCode"`
}
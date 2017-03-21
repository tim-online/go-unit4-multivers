package multivers

// https://api.online.unit4.nl/V14/Help/Api/GET-api-database-ProductInfoList_productId_shortName_description_productGroupId

import (
	"context"
	"net/url"

	"fmt"
	"github.com/gorilla/schema"
)

const (
	OrganizationInfoListPath = "/api/%s/OrganizationInfoList.json"
)

func NewOrganizationInfoListService(client *Client) *OrganizationInfoListService {
	return &OrganizationInfoListService{Client: client}
}

type OrganizationInfoListService struct {
	Client *Client
}

func (s *OrganizationInfoListService) Get(database string, requestParams *OrganizationInfoListGetParams, ctx context.Context) (*OrganizationInfoListGetResponse, error) {
	method := "GET"
	responseBody := NewOrganizationInfoListGetResponse()
	path := fmt.Sprintf(OrganizationInfoListPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// Process query parameters
	addQueryParamsToRequest(requestParams, httpReq, false)

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewOrganizationInfoListGetParams() *OrganizationInfoListGetParams {
	return &OrganizationInfoListGetParams{}
}

type OrganizationInfoListGetParams struct {
	Email     string `schema:"email"`
	ShortName string `schema:"shortName"`
	Name      string `schema:"name"`
	City      string `schema:"city"`
}

func (p *OrganizationInfoListGetParams) FromQueryParams(queryParams url.Values) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(p, queryParams)
}

func NewOrganizationInfoListGetResponse() *OrganizationInfoListGetResponse {
	return &OrganizationInfoListGetResponse{}
}

type OrganizationInfoListGetResponse []OrganizationInfo

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
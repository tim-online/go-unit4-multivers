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
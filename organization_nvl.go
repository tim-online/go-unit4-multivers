package multivers

import (
	"context"
	"net/url"

	"fmt"
	"github.com/gorilla/schema"
	"errors"
)

const (
	OrganizationNVLPath = "/api/%s/OrganizationNVL.json"
)

func NewOrganizationNVLService(client *Client) *OrganizationNVLService {
	return &OrganizationNVLService{Client: client}
}

type OrganizationNVLService struct {
	Client *Client
}

func (s *OrganizationNVLService) Get(database string, requestParams *OrganizationNVLGetParams, ctx context.Context) (*OrganizationNVLGetResponse, error) {
	return nil, errors.New("Call doesn't work correctly.")

	method := "GET"
	responseBody := NewOrganizationNVLGetResponse()
	path := fmt.Sprintf(OrganizationNVLPath, database)

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

func NewOrganizationNVLGetParams() *OrganizationNVLGetParams {
	return &OrganizationNVLGetParams{}
}

type OrganizationNVLGetParams struct {
	MustBeCustomer string `schema:"mustBeCustomer"`
	MustBeSupplier string `schema:"mustBeSupplier"`
}

func (p *OrganizationNVLGetParams) FromQueryParams(queryParams url.Values) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(p, queryParams)
}

func NewOrganizationNVLGetResponse() *OrganizationNVLGetResponse {
	return &OrganizationNVLGetResponse{}
}

type OrganizationNVLGetResponse []NVL

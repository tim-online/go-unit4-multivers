package multivers

import (
	"context"
	"fmt"
)

const (
	CustomerGroupNVLPath = "/api/%s/CustomerGroupNVL.json"
)

func NewCustomerGroupNVLService(client *Client) *CustomerGroupNVLService {
	return &CustomerGroupNVLService{Client: client}
}

type CustomerGroupNVLService struct {
	Client *Client
}

func (s *CustomerGroupNVLService) Get(database string,requestParams *CustomerGroupNVLGetParams, ctx context.Context) (*CustomerGroupNVLGetResponse, error) {
	method := "GET"
	responseBody := NewCustomerGroupNVLGetResponse()
	path := fmt.Sprintf(CustomerGroupNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	if requestParams.FiscalYear != 0 {
		// Process query parameters
		addQueryParamsToRequest(requestParams, httpReq, false)
	}


	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewCustomerGroupNVLGetResponse() *CustomerGroupNVLGetResponse {
	return &CustomerGroupNVLGetResponse{}
}

func NewCustomerGroupNVLGetParams() *CustomerGroupNVLGetParams {
	return &CustomerGroupNVLGetParams{}
}

type CustomerGroupNVLGetParams struct {
	FiscalYear int `schema:"fiscalYear"`
}

type CustomerGroupNVLGetResponse []CustomerGroupNVL

type CustomerGroupNVL struct {
	Name  int `json:"name"`
	Value string `json:"value"`
}
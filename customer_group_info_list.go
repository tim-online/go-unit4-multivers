package multivers

import (
	"context"
	"fmt"
)

const (
	CustomerGroupInfoListPath = "/api/%s/CustomerGroupInfoList.json"
)

func NewCustomerGroupInfoListService(client *Client) *CustomerGroupInfoListService {
	return &CustomerGroupInfoListService{Client: client}
}

type CustomerGroupInfoListService struct {
	Client *Client
}
//todo No test data in the database
func (s *CustomerGroupInfoListService) Get(database string, requestParams *CustomerGroupInfoListGetParams, ctx context.Context) (*CustomerGroupInfoListGetResponse, error) {
	method := "GET"
	responseBody := NewCustomerGroupInfoListGetResponse()
	path := fmt.Sprintf(CustomerGroupInfoListPath, database)

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

func NewCustomerGroupInfoListGetResponse() *CustomerGroupInfoListGetResponse {
	return &CustomerGroupInfoListGetResponse{}
}

func NewCustomerGroupInfoListGetParams() *CustomerGroupInfoListGetParams {
	return &CustomerGroupInfoListGetParams{}
}

type CustomerGroupInfoListGetParams struct {
	FiscalYear int `schema:"fiscalYear"`
}

type CustomerGroupInfoListGetResponse []CustomerGroupInfo


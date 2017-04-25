package multivers

import (
	"context"
	"fmt"
)

const (
	CustomerGroupInfoPath = "/api/%s/CustomerGroupInfo/%d.json"
)

func NewCustomerGroupInfoService(client *Client) *CustomerGroupInfoService {
	return &CustomerGroupInfoService{Client: client}
}

type CustomerGroupInfoService struct {
	Client *Client
}

func (s *CustomerGroupInfoService) Get(database string, customerGroupID int,requestParams *CustomerGroupInfoGetParams, ctx context.Context) (*CustomerGroupInfoGetResponse, error) {
	method := "GET"
	responseBody := NewCustomerGroupInfoGetResponse()
	path := fmt.Sprintf(CustomerGroupInfoPath, database, customerGroupID)

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

func NewCustomerGroupInfoGetResponse() *CustomerGroupInfoGetResponse {
	return &CustomerGroupInfoGetResponse{}
}

func NewCustomerGroupInfoGetParams() *CustomerGroupInfoGetParams {
	return &CustomerGroupInfoGetParams{}
}

type CustomerGroupInfoGetParams struct {
	FiscalYear int `schema:"fiscalYear"`
}

type CustomerGroupInfoGetResponse CustomerGroupInfo

type CustomerGroupInfo struct {
	CustomerGroupID int    `json:"customerGroupId"`
	FiscalYear      int    `json:"fiscalYear"`
	LedgerAccountID string `json:"ledgerAccountId"`
}

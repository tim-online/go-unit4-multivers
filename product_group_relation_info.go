package multivers

import (
	"context"
	"fmt"
)

const (
	ProductGroupRelationInfoPath = "/api/%s/ProductGroupRelationInfo/%s/%d.json"
)

func NewProductGroupRelationInfoService(client *Client) *ProductGroupRelationInfoService {
	return &ProductGroupRelationInfoService{Client: client}
}

type ProductGroupRelationInfoService struct {
	Client *Client
}

func (s *ProductGroupRelationInfoService) Get(database string, productGroupID string, fiscalYear int, ctx context.Context) (*ProductGroupRelationInfoGetResponse, error) {
	method := "GET"
	responseBody := NewProductGroupRelationInfoGetResponse()
	path := fmt.Sprintf(ProductGroupRelationInfoPath, database, productGroupID, fiscalYear)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewProductGroupRelationInfoGetResponse() *ProductGroupRelationInfoGetResponse {
	return &ProductGroupRelationInfoGetResponse{}
}

type ProductGroupRelationInfoGetResponse ProductGroupRelationInfo

type ProductGroupRelationInfo struct {
	DiscountAccountID string `json:"discountAccountId"`
	FiscalYear        int    `json:"fiscalYear"`
	ProductGroupID    string `json:"productGroupId"`
	RevenueAccountID  string `json:"revenueAccountId"`
}

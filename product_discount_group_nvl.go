package multivers

import (
	"context"
	"fmt"
)

const (
	ProductDiscountGroupNVLPath = "/api/%s/ProductDiscountGroupNVL.json"
)

func NewProductDiscountGroupNVLService(client *Client) *ProductDiscountGroupNVLService {
	return &ProductDiscountGroupNVLService{Client: client}
}

type ProductDiscountGroupNVLService struct {
	Client *Client
}

func (s *ProductDiscountGroupNVLService) Get(database string, ctx context.Context) (*ProductDiscountGroupNVLGetResponse, error) {
	method := "GET"
	responseBody := NewProductDiscountGroupNVLGetResponse()
	path := fmt.Sprintf(ProductDiscountGroupNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewProductDiscountGroupNVLGetResponse() *ProductDiscountGroupNVLGetResponse {
	return &ProductDiscountGroupNVLGetResponse{}
}

type ProductDiscountGroupNVLGetResponse []NVL

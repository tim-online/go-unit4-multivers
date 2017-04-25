package multivers

import (
	"context"
	"fmt"
)

const (
	ProductNVLPath = "/api/%s/ProductNVL"
)

func NewProductNVLService(client *Client) *ProductNVLService {
	return &ProductNVLService{Client: client}
}

type ProductNVLService struct {
	Client *Client
}

func (s *ProductNVLService) Get(database string, ctx context.Context) (*ProductNVLGetResponse, error) {
	method := "GET"
	responseBody := NewProductNVLGetResponse()
	path := fmt.Sprintf(ProductNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewProductNVLGetResponse() *ProductNVLGetResponse {
	return &ProductNVLGetResponse{}
}

type ProductNVLGetResponse []NVL

package multivers

import (
	"context"
	"fmt"
)

const (
	CustomerInfoListPath = "/api/%s/CustomerInfoList.json"
)

func NewCustomerInfoListService(client *Client) *CustomerInfoListService {
	return &CustomerInfoListService{Client: client}
}

type CustomerInfoListService struct {
	Client *Client
}

func (s *CustomerInfoListService) Get(database string, ctx context.Context) (*CustomerInfoListGetResponse, error) {
	method := "GET"
	responseBody := NewCustomerInfoListGetResponse()
	path := fmt.Sprintf(CustomerInfoListPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewCustomerInfoListGetResponse() *CustomerInfoListGetResponse {
	return &CustomerInfoListGetResponse{}
}

type CustomerInfoListGetResponse []CustomerInfo
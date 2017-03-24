package multivers

import (
	"context"
)

const (
	AdministrationGroupNVLPath = "/api/AdministrationGroupNVL.json"
)

func NewAdministrationGroupNVLService(client *Client) *AdministrationGroupNVLService {
	return &AdministrationGroupNVLService{Client: client}
}

type AdministrationGroupNVLService struct {
	Client *Client
}

//todo No test data in the database
func (s *AdministrationGroupNVLService) Get(requestParams *AdministrationGroupNVLGetParams, ctx context.Context) (*AdministrationGroupNVLGetResponse, error) {
	method := "GET"
	responseBody := NewAdministrationGroupNVLGetResponse()
	path := AdministrationGroupNVLPath

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

func NewAdministrationGroupNVLGetResponse() *AdministrationGroupNVLGetResponse {
	return &AdministrationGroupNVLGetResponse{}
}

func NewAdministrationGroupNVLGetParams() *AdministrationGroupNVLGetParams {
	return &AdministrationGroupNVLGetParams{}
}

type AdministrationGroupNVLGetParams struct {
	UserName string `scheme:"userName"`
}

type AdministrationGroupNVLGetResponse []NVL


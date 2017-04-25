package multivers

import (
	"context"
	"fmt"
)

const (
	AdministrationNVLPath = "/api/AdministrationNVL"
)

func NewAdministrationNVLService(client *Client) *AdministrationNVLService {
	return &AdministrationNVLService{Client: client}
}

type AdministrationNVLService struct {
	Client *Client
}

func (s *AdministrationNVLService) GetByGroupIdAndUsername(groupID int, userName string, ctx context.Context) (*AdministrationNVLGetResponse, error) {
	method := "GET"
	responseBody := NewAdministrationNVLGetResponse()
	prePath := AdministrationNVLPath + "/%d.json"
	path := fmt.Sprintf(prePath, groupID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	type getParams struct{
		UserName string `schema:"userName"`
	}

	requestParams := getParams{userName}

	// Process query parameters
	addQueryParamsToRequest(requestParams, httpReq, false)

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *AdministrationNVLService) GetByGroupId(groupID int, ctx context.Context) (*AdministrationNVLGetResponse, error) {
	method := "GET"
	responseBody := NewAdministrationNVLGetResponse()
	path := AdministrationNVLPath + ".json"

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	type getParams struct{
		GroupID int `schema:"groupId"`
	}

	requestParams := getParams{groupID}

	// Process query parameters
	addQueryParamsToRequest(requestParams, httpReq, false)

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *AdministrationNVLService) GetByUserName(userName string, ctx context.Context) (*AdministrationNVLGetResponse, error) {
	method := "GET"
	responseBody := NewAdministrationNVLGetResponse()
	path := AdministrationNVLPath + ".json"

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	type getParams struct{
		UserName string `schema:"userName"`
	}

	requestParams := getParams{userName}

	// Process query parameters
	addQueryParamsToRequest(requestParams, httpReq, false)

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewAdministrationNVLGetResponse() *AdministrationNVLGetResponse {
	return &AdministrationNVLGetResponse{}
}

type AdministrationNVLGetResponse []NVL


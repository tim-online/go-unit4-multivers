package multivers

import (
	"context"
	"fmt"
)

const (
	OrganizationStateNVLPath = "/api/%s/OrganizationStateNVL.json"
)

func NewOrganizationStateNVLService(client *Client) *OrganizationStateNVLService {
	return &OrganizationStateNVLService{Client: client}
}

type OrganizationStateNVLService struct {
	Client *Client
}

func (s *OrganizationStateNVLService) Get(database string, ctx context.Context) (*OrganizationStateNVLGetResponse, error) {
	method := "GET"
	responseBody := NewOrganizationStateNVLGetResponse()
	path := fmt.Sprintf(OrganizationStateNVLPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewOrganizationStateNVLGetResponse() *OrganizationStateNVLGetResponse {
	return &OrganizationStateNVLGetResponse{}
}

type OrganizationStateNVLGetResponse []NVL

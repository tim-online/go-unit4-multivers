package multivers

import (
	"context"
	"github.com/gorilla/schema"
	"net/url"
)

const (
	AdministrationGroupPath = "/api/AdministrationGroup.json"
)

func NewAdministrationGroupService(client *Client) *AdministrationGroupService {
	return &AdministrationGroupService{Client: client}
}

type AdministrationGroupService struct {
	Client *Client
}

func (s *AdministrationGroupService) Get(requestParams *AdministrationGroupGetParams, ctx context.Context) (*AdministrationGroupGetResponse, error) {
	method := "GET"
	responseBody := NewAdministrationGroupGetResponse()
	path := AdministrationGroupPath

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

func NewAdministrationGroupGetParams() *AdministrationGroupGetParams {
	return &AdministrationGroupGetParams{}
}

type AdministrationGroupGetParams struct {
	GroupID int `schema:"groupId"`
}

func (p *AdministrationGroupGetParams) FromQueryParams(queryParams url.Values) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(p, queryParams)
}

func NewAdministrationGroupGetResponse() *AdministrationGroupGetResponse {
	return &AdministrationGroupGetResponse{}
}

type AdministrationGroupGetResponse AdministrationGroup

type AdministrationGroup struct {
	Messages  []Message `json:"messages"`
	CanChange bool      `json:"canChange"`
	GroupID   int       `json:"groupId"`
	IsDefault bool      `json:"isDefault"`
	Name      string    `json:"name"`
}

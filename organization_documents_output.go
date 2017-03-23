package multivers

import (
	"context"
	"fmt"
)

const (
	OrganizationDocumentsOutputPath = "/api/%s/OrganizationDocumentsOutput/%d.json"
)

func NewOrganizationDocumentsOutputService(client *Client) *OrganizationDocumentsOutputService {
	return &OrganizationDocumentsOutputService{Client: client}
}

type OrganizationDocumentsOutputService struct {
	Client *Client
}

func (s *OrganizationDocumentsOutputService) Get(database string, organizationID int, ctx context.Context) (*OrganizationDocumentsOutputGetResponse, error) {
	method := "GET"
	responseBody := NewOrganizationDocumentsOutputGetResponse()
	path := fmt.Sprintf(OrganizationDocumentsOutputPath, database, organizationID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewOrganizationDocumentsOutputGetResponse() *OrganizationDocumentsOutputGetResponse {
	return &OrganizationDocumentsOutputGetResponse{}
}

type OrganizationDocumentsOutputGetResponse OrganizationDocumentsOutput

type OrganizationDocumentsOutput struct {
	Messages       []string           `json:"messages"`
	CanChange      bool               `json:"canChange"`
	DocumentTypes  []DocumentTypeInfo `json:"documentTypes"`
	OrganizationId int                `json:"organizationId"`
}

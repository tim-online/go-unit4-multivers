package multivers

import (
	"context"
	"fmt"
)

const (
	CompanyDetailsPath = "/api/%s/CompanyDetails.json"
)

func NewCompanyDetailsService(client *Client) *CompanyDetailsService {
	return &CompanyDetailsService{Client: client}
}

type CompanyDetailsService struct {
	Client *Client
}

func (s *CompanyDetailsService) Get(database string, ctx context.Context) (*CompanyDetailsGetResponse, error) {
	method := "GET"
	responseBody := NewCompanyDetailsGetResponse()
	path := fmt.Sprintf(CompanyDetailsPath, database)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewCompanyDetailsGetResponse() *CompanyDetailsGetResponse {
	return &CompanyDetailsGetResponse{}
}

type CompanyDetailsGetResponse CompanyDetails

type CompanyDetails struct {
	AdministrationCode        string    `json:"administrationCode"`
	AdministrationDescription string    `json:"administrationDescription"`
	AdministrationShortName   string    `json:"administrationShortName"`
	Messages                  []Message `json:"messages"`
	CanChange                 bool      `json:"canChange"`
	City                      string    `json:"city"`
	CocRegistration           string    `json:"cocRegistration"`
	Email                     string    `json:"email"`
	EmailAccountType          int       `json:"emailAccountType"`
	PhoneNumber               string    `json:"phoneNumber"`
	Street                    string    `json:"street"`
	VatNumber                 string    `json:"vatNumber"`
	ZipCode                   string    `json:"zipCode"`
}

package multivers

import (
	"context"
	"fmt"
)

const (
	AdministrationInfoPath = "/api/AdministrationInfo/%s"
)

func NewAdministrationInfoService(client *Client) *AdministrationInfoService {
	return &AdministrationInfoService{Client: client}
}

type AdministrationInfoService struct {
	Client *Client
}

func (s *AdministrationInfoService) Get(administrationID string, ctx context.Context) (*AdministrationInfoGetResponse, error) {
	method := "GET"
	responseBody := NewAdministrationInfoGetResponse()
	path := fmt.Sprintf(AdministrationInfoPath, administrationID)

	// create a new HTTP request
	httpReq, err := s.Client.NewRequest(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	// submit the request
	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func NewAdministrationInfoGetResponse() *AdministrationInfoGetResponse {
	return &AdministrationInfoGetResponse{}
}

type AdministrationInfoGetResponse AdministrationInfo

type AdministrationInfo struct {
	AdministrationID    string   `json:"administrationId"`
	BackupDate          DateNLNL `json:"backupDate"`
	Code                string   `json:"code"`
	Description         string   `json:"description"`
	GroupID             int      `json:"groupId"`
	OnlineState         int      `json:"onlineState"`
	PreviousOnlineState int      `json:"previousOnlineState"`
	ReportPath          string   `json:"reportPath"`
	ShortName           string   `json:"shortName"`
	Users               []string `json:"users"`
	Version             int      `json:"version"`
}

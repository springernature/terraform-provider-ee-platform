// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type TeamsResponse struct {
	Teams Teams `json:"teams"`
}

type Teams []Team
type Team struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	SnPaasOrg string `json:"snpaas_org"`
}

type PlatformClient interface {
	GetTeams() (Teams, error)
}

type platformClient struct {
	teamsEndpoint string
}

func (t platformClient) GetTeams() (teams Teams, err error) {
	u, err := url.Parse(t.teamsEndpoint)
	if err != nil {
		return
	}

	response, err := http.Get(u.JoinPath("/api/v1/teams").String())
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = fmt.Errorf("%s", response.Status)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	teamsResponse := TeamsResponse{}
	err = json.Unmarshal(body, &teamsResponse)
	return teamsResponse.Teams, err
}

func NewPlatformClient(teamsEndpoint string) PlatformClient {
	return platformClient{
		teamsEndpoint: teamsEndpoint,
	}
}

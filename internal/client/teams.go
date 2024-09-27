package client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Teams []Team
type Team struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Domain     string `json:"domain"`
	SnPaasOrg  string `json:"snpaas_org"`
}

type TeamsClient interface {
	GetTeams() (Teams, error)
}

type teamsClient struct {
	teamsEndpoint string
}

func (t teamsClient) GetTeams() (teams Teams, err error) {
	u, err := url.Parse(t.teamsEndpoint)
	if err != nil {
		return
	}

	response, err := http.Get(u.JoinPath("/api/teams").String())
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &teams)
	return
}

func NewTeamsClient(teamsEndpoint string) TeamsClient {
	return teamsClient{
		teamsEndpoint: teamsEndpoint,
	}
}

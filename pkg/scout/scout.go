package scout

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Scout struct {
	URL string
}

func New(dbAddress string) Scout {
	return Scout{
		URL: dbAddress,
	}
}

func (s *Scout) post(endpoint string, values url.Values, body string) (*http.Request, error) {
	var query string

	if len(values) > 0 {
		query = "?" + values.Encode()
	} else {
		query = ""
	}

	url, err := url.Parse(s.URL + endpoint + query)

	if err != nil || url.Scheme == "" || url.Host == "" {
		return nil, fmt.Errorf("malformed url")
	}

	request, err := http.NewRequest("POST", url.String(), strings.NewReader(body))

	if err != nil {
		return nil, err
	}

	return request, nil
}

const (
	newTeamURL = "/api/new-team"
)

func (s *Scout) InsertTeam(team Team) error {
	body, err := json.Marshal(team)

	if err != nil {
		return err
	}

	request, err := s.post(newTeamURL, url.Values{}, string(body))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("status code %d", response.StatusCode)
	}

	return nil
}

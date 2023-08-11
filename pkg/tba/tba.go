package tba

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TBA struct {
	APIKey string
}

func New(apiKey string) TBA {
	return TBA{
		APIKey: apiKey,
	}
}

const (
	tbaURL = "https://www.thebluealliance.com/api/v3"
)

func (tba *TBA) get(endpoint string, headers map[string]string) (*http.Request, error) {
	url, err := url.Parse(tbaURL + endpoint)

	if err != nil || url.Scheme == "" || url.Host == "" {
		return nil, err
	}

	request, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return nil, err
	}

	request.Header.Add("X-TBA-Auth-Key", tba.APIKey)

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	return request, nil
}

const (
	eventTeamsURL = "/event/%s/teams"
)

func (tba *TBA) GetTeams(eventKey string) ([]Team, error) {
	endpoint := fmt.Sprintf(eventTeamsURL, eventKey)

	request, err := tba.get(endpoint, nil)

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	var teams []Team
	json.Unmarshal(body, &teams)

	return teams, nil
}

const (
	eventURL = "/event/%s/simple"
)

func (tba *TBA) GetEvent(eventKey string) (Event, error) {
	endpoint := fmt.Sprintf(eventURL, eventKey)

	request, err := tba.get(endpoint, nil)

	if err != nil {
		return Event{}, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return Event{}, err
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return Event{}, err
	}

	var event Event
	json.Unmarshal(body, &event)

	return event, nil
}

const (
	getMatchKeysURL = "/event/%s/matches/keys"
)

func (tba *TBA) GetMatchKeys(eventKey string) ([]string, error) {
	endpoint := fmt.Sprintf(getMatchKeysURL, eventKey)

	request, err := tba.get(endpoint, nil)

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	var matchKeys []string
	json.Unmarshal(body, &matchKeys)

	return matchKeys, nil
}

const (
	getMatchURL = "/match/%s"
)

func (tba *TBA) GetMatch(eventKey string, year int) (any, error) {
	endpoint := fmt.Sprintf(getMatchURL, eventKey)

	request, err := tba.get(endpoint, nil)

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	switch year {
	case 2023:
		var match2023 Match2023
		json.Unmarshal(body, &match2023)

		return match2023, nil
	case 2022:
		var match2022 Match2022
		json.Unmarshal(body, &match2022)

		return match2022, nil
	default:
		return nil, fmt.Errorf("match struct not implemented for %d", year)
	}
}

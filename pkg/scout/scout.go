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

	if response.StatusCode != 200 && response.StatusCode != 500 {
		return fmt.Errorf("status code %d", response.StatusCode)
	}

	return nil
}

const (
	newEventURL = "/api/new-event"
)

func (s *Scout) InsertEvent(event Event) error {
	body, err := json.Marshal(event)

	if err != nil {
		return err
	}

	request, err := s.post(newEventURL, url.Values{}, string(body))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 && response.StatusCode != 500 {
		return fmt.Errorf("status code %d", response.StatusCode)
	}

	return nil
}

const (
	newSeasonURL = "/api/new-season"
)

func (t *Team) ToValues() url.Values {
	values := url.Values{}

	values.Add("team", fmt.Sprint(t.Number))

	return values
}

func (s *Scout) InsertSeason(season Season, team Team) error {
	body, err := json.Marshal(season)

	if err != nil {
		return err
	}

	request, err := s.post(newSeasonURL, team.ToValues(), string(body))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 && response.StatusCode != 500 {
		return fmt.Errorf("status code %d", response.StatusCode)
	}

	return nil
}

const (
	addEventURL = "/api/add-event"
)

func (s *Season) ToValues() url.Values {
	values := url.Values{}

	values.Add("year", fmt.Sprint(s.Year))

	return values
}

func (e *Event) ToValues() url.Values {
	values := url.Values{}

	values.Add("event", e.Name)
	values.Add("region", e.Region)
	values.Add("year", fmt.Sprint(e.Year))
	values.Add("week", fmt.Sprint(e.Week))

	return values
}

func join(vs ...url.Values) url.Values {
	values := url.Values{}

	for _, v := range vs {
		for key, value := range v {
			if values.Has(key) {
				continue
			}

			values.Add(key, value[0])
		}
	}

	return values
}

func (s *Scout) AddEvent(event Event, season Season, team Team) error {
	request, err := s.post(addEventURL, join(event.ToValues(), season.ToValues(), team.ToValues()), "")

	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 && response.StatusCode != 500 {
		return fmt.Errorf("status code %d", response.StatusCode)
	}

	return nil
}

const (
	newRobotURL = "/api/new-robot"
)

func (s *Scout) InsertRobot(robot Robot, season Season, team Team) error {
	body, err := json.Marshal(robot)

	if err != nil {
		return err
	}

	request, err := s.post(newRobotURL, join(season.ToValues(), team.ToValues()), string(body))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 && response.StatusCode != 500 {
		return fmt.Errorf("status code %d", response.StatusCode)
	}

	return nil
}

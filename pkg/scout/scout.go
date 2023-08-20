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

	fmt.Println(request)

	if err != nil {
		return nil, err
	}

	return request, nil
}

const (
	created  = 201
	conflict = 409
)

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

	if response.StatusCode != created && response.StatusCode != conflict {
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

	if response.StatusCode != created && response.StatusCode != conflict {
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

	if response.StatusCode != created && response.StatusCode != conflict {
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

	values.Add("event", e.Code)

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

	if response.StatusCode != created && response.StatusCode != conflict {
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

	if response.StatusCode != created && response.StatusCode != conflict {
		return fmt.Errorf("status code %d", response.StatusCode)
	}

	return nil
}

const (
	newMatchURL = "/api/new-match"
)

func (s *Scout) InsertMatch(match Match, event Event) error {
	body, err := json.Marshal(match)

	if err != nil {
		return err
	}

	request, err := s.post(newMatchURL, event.ToValues(), string(body))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != created && response.StatusCode != conflict {
		return fmt.Errorf("status code %d", response.StatusCode)
	}

	return nil
}

const (
	newParticipantURL = "/api/new-participant"
)

func (r *Robot) ToValues() url.Values {
	values := url.Values{}

	values.Add("robot", r.Name)

	return values
}

func (m *Match) ToValues() url.Values {
	values := url.Values{}

	values.Add("match", m.MatchKey())

	return values
}

func (m *Match) MatchKey() string {
	var s string

	s += m.Type

	if m.Type != "qm" {
		s += fmt.Sprint(m.Set)
		s += "m"
	}

	s += fmt.Sprint(m.Number)

	return s
}

func (s *Scout) InsertParticipant(participant Participant, match Match, event Event) error {
	body, err := json.Marshal(participant)

	if err != nil {
		return err
	}

	request, err := s.post(newParticipantURL, join(match.ToValues(), event.ToValues()), string(body))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != created && response.StatusCode != conflict {
		return fmt.Errorf("status code %d", response.StatusCode)
	}

	return nil
}

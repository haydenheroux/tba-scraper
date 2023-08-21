package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/haydenheroux/adapter"
	"github.com/haydenheroux/data"
	"github.com/haydenheroux/scout"
	"github.com/haydenheroux/tba"
)

const (
	APP_NAME          = "tbascraper"
	DEFAULT_API_KEY   = ""
	DEFAULT_SCOUT_URL = ""
)

var (
	apiKey   string
	scoutURL string
)

func init() {
	flag.StringVar(&apiKey, "apiKey", DEFAULT_API_KEY, "API Key")
	flag.StringVar(&scoutURL, "scoutURL", DEFAULT_SCOUT_URL, "Scout URL")
}

var (
	logger *log.Logger
	api    tba.TBA
	db     scout.Scout
)

func main() {
	flag.Parse()

	logger = log.New(os.Stderr, APP_NAME+": ", 0)

	api = tba.New(apiKey)

	db = scout.New(scoutURL)

	eventKeys := flag.Args()

	for _, eventKey := range eventKeys {
		run(eventKey)
	}
}

func run(eventKey string) {
	logger.Printf("running %s\n", eventKey)

	_event, err := api.GetEvent(eventKey)

	if err != nil {
		logger.Fatalf("Failed to get event: %v\n", err)
	}

	event := adapter.ToEvent(_event)

	println(event.Code)

	if err := db.InsertEvent(event); err != nil {
		logger.Fatalf("Failed to insert event: %v\n%v\n", err, event)
	}

	teams, err := api.GetTeams(eventKey)

	if err != nil {
		logger.Fatalf("Failed to get teams: %v\n", err)
	}

	for _, team := range teams {
		team := adapter.ToTeam(team)

		if err := db.InsertTeam(team); err != nil {
			logger.Fatalf("Failed to insert team: %v\n%v\n", err, team)
		}

		season := scout.Season{
			Year:   event.Year,
			Robots: []scout.Robot{},
			Events: []scout.Event{},
		}

		if err := db.InsertSeason(season, team); err != nil {
			logger.Fatalf("Failed to insert season: %v\n", err)
		}

		robot := scout.Robot{
			Name: fmt.Sprintf("%d %d Robot", team.Number, season.Year),
		}

		if err := db.InsertRobot(robot, season, team); err != nil {
			logger.Fatalf("Failed to insert robot: %v\n%v\n", err, robot)
		}

		if err := db.AddEvent(event, season, team); err != nil {
			logger.Fatalf("Failed to add event: %v\n%v\n", err, event)
		}
	}

	matchKeys, err := api.GetMatchKeys(eventKey)

	if err != nil {
		logger.Fatalf("Failed to get match keys: %v\n", err)
	}

	for _, matchKey := range matchKeys {
		_match, err := api.GetMatch(matchKey, event.Year)

		if err != nil {
			logger.Fatalf("Failed to get match: %v\n", err)
		}

		match := adapter.ToMatch(_match)

		if err := db.InsertMatch(match, event); err != nil {
			logger.Fatalf("Failed to add match: %v\n%v\n", err, match)
		}

		switch _match.(type) {
		case tba.Match2022:
			doMatch2022(_match.(tba.Match2022), _event)
		case tba.Match2023:
			doMatch2023(_match.(tba.Match2023), _event)
		}
	}
}

func doMatch2022(match tba.Match2022, event tba.Event) {
	blueAllianceMetrics := data.AllianceMetrics2022(match.ScoreBreakdown.Blue)
	blueAlliance := scout.Alliance{Color: "blue", Metrics: blueAllianceMetrics, Participants: make([]scout.Participant, 0)}

	if err := db.InsertAlliance(blueAlliance, adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
		logger.Fatalf("Failed to add alliance: %v\n%v\n", err, blueAlliance)
	}

	redAllianceMetrics := data.AllianceMetrics2022(match.ScoreBreakdown.Red)
	redAlliance := scout.Alliance{Color: "red", Metrics: redAllianceMetrics, Participants: make([]scout.Participant, 0)}

	if err := db.InsertAlliance(redAlliance, adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
		logger.Fatalf("Failed to add alliance: %v\n%v\n", err, redAlliance)
	}

	for _, teamKey := range match.Alliances.Blue.TeamKeys {
		teamNumber, err := strconv.Atoi(strings.Split(teamKey, "frc")[1])

		if err != nil {
			logger.Fatalf("Failed to get team number: %v\n", err)
		}

		participant := scout.Participant{
			TeamNumber: teamNumber,
			Metrics:    []scout.Metric{},
		}

		if err != nil {
			logger.Fatalf("Failed to get team: %v\n", err)
		}

		if err := db.InsertParticipant(participant, blueAlliance, adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
			logger.Fatalf("Failed to add participant: %v\n%v\n", err, participant)
		}
	}

	for _, teamKey := range match.Alliances.Red.TeamKeys {
		teamNumber, err := strconv.Atoi(strings.Split(teamKey, "frc")[1])

		if err != nil {
			logger.Fatalf("Failed to get team number: %v\n", err)
		}

		participant := scout.Participant{
			TeamNumber: teamNumber,
			Metrics:    []scout.Metric{},
		}

		if err != nil {
			logger.Fatalf("Failed to get team: %v\n", err)
		}

		if err := db.InsertParticipant(participant, redAlliance, adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
			logger.Fatalf("Failed to add participant: %v\n%v\n", err, participant)
		}
	}

}

func doMatch2023(match tba.Match2023, event tba.Event) {
	blueAllianceMetrics := data.AllianceMetrics2023(match.ScoreBreakdown.Blue)
	blueAlliance := scout.Alliance{Color: "blue", Metrics: blueAllianceMetrics, Participants: make([]scout.Participant, 0)}

	if err := db.InsertAlliance(blueAlliance, adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
		logger.Fatalf("Failed to add alliance: %v\n%v\n", err, blueAlliance)
	}

	redAllianceMetrics := data.AllianceMetrics2023(match.ScoreBreakdown.Red)
	redAlliance := scout.Alliance{Color: "red", Metrics: redAllianceMetrics, Participants: make([]scout.Participant, 0)}

	if err := db.InsertAlliance(redAlliance, adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
		logger.Fatalf("Failed to add alliance: %v\n%v\n", err, redAlliance)
	}

	for _, teamKey := range match.Alliances.Blue.TeamKeys {
		teamNumber, err := strconv.Atoi(strings.Split(teamKey, "frc")[1])

		if err != nil {
			logger.Fatalf("Failed to get team number: %v\n", err)
		}

		participant := scout.Participant{
			TeamNumber: teamNumber,
			Metrics:    []scout.Metric{},
		}

		if err != nil {
			logger.Fatalf("Failed to get team: %v\n", err)
		}

		if err := db.InsertParticipant(participant, blueAlliance, adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
			logger.Fatalf("Failed to add participant: %v\n%v\n", err, participant)
		}
	}

	for _, teamKey := range match.Alliances.Red.TeamKeys {
		teamNumber, err := strconv.Atoi(strings.Split(teamKey, "frc")[1])

		if err != nil {
			logger.Fatalf("Failed to get team number: %v\n", err)
		}

		participant := scout.Participant{
			TeamNumber: teamNumber,
			Metrics:    []scout.Metric{},
		}

		if err != nil {
			logger.Fatalf("Failed to get team: %v\n", err)
		}

		if err := db.InsertParticipant(participant, redAlliance, adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
			logger.Fatalf("Failed to add participant: %v\n%v\n", err, participant)
		}
	}

}

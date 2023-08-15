package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"haydenheroux.github.io/adapter"
	"haydenheroux.github.io/data"
	"haydenheroux.github.io/scout"
	"haydenheroux.github.io/tba"
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

	var wg sync.WaitGroup

	for _, eventKey := range eventKeys {
		wg.Add(1)

		go func(eventKey string) {
			defer wg.Done()

			run(eventKey)
		}(eventKey)
	}

	wg.Wait()
}

func run(eventKey string) {
	logger.Printf("running %s\n", eventKey)

	event, err := api.GetEvent(eventKey)

	if err != nil {
		logger.Fatalf("Failed to get event: %v\n", err)
	}

	if err := db.InsertEvent(adapter.ToEvent(event)); err != nil {
		logger.Fatalf("Failed to insert event: %v\n", err)
	}

	teams, err := api.GetTeams(eventKey)

	if err != nil {
		logger.Fatalf("Failed to get teams: %v\n", err)
	}

	for _, team := range teams {
		team := adapter.ToTeam(team)

		if err := db.InsertTeam(team); err != nil {
			logger.Fatalf("Failed to insert team: %v\n", err)
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
			logger.Fatalf("Failed to insert robot: %v\n", err)
		}

		if err := db.AddEvent(adapter.ToEvent(event), season, team); err != nil {
			logger.Fatalf("Failed to add event: %v\n", err)
		}
	}

	matchKeys, err := api.GetMatchKeys(eventKey)

	if err != nil {
		logger.Fatalf("Failed to get match keys: %v\n", err)
	}

	for _, matchKey := range matchKeys {
		match, err := api.GetMatch(matchKey, event.Year)

		if err != nil {
			logger.Fatalf("Failed to get match: %v\n", err)
		}

		if err := db.InsertMatch(adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
			logger.Fatalf("Failed to add match: %v\n", err)
		}

		switch match.(type) {
		case tba.Match2022:
			doMatch2022(match.(tba.Match2022), event)
		case tba.Match2023:
			doMatch2023(match.(tba.Match2023), event)
		}
	}
}

func doMatch2022(match tba.Match2022, event tba.Event) {
	for n, teamKey := range match.Alliances.Blue.TeamKeys {
		teamNumber, err := strconv.Atoi(strings.Split(teamKey, "frc")[1])

		if err != nil {
			logger.Fatalf("Failed to get team number: %v\n", err)
		}

		participant := scout.Participant{
			Alliance:   "blue",
			TeamNumber: teamNumber,
			Metrics:    data.Metrics2022(match.ScoreBreakdown.Blue, n),
		}

		if err != nil {
			logger.Fatalf("Failed to get team: %v\n", err)
		}

		if err := db.InsertParticipant(participant, adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
			logger.Fatalf("Failed to add participant: %v\n", err)
		}
	}

	for n, teamKey := range match.Alliances.Red.TeamKeys {
		teamNumber, err := strconv.Atoi(strings.Split(teamKey, "frc")[1])

		if err != nil {
			logger.Fatalf("Failed to get team number: %v\n", err)
		}

		participant := scout.Participant{
			Alliance:   "red",
			TeamNumber: teamNumber,
			Metrics:    data.Metrics2022(match.ScoreBreakdown.Red, n),
		}

		if err != nil {
			logger.Fatalf("Failed to get team: %v\n", err)
		}

		if err := db.InsertParticipant(participant, adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
			logger.Fatalf("Failed to add participant: %v\n", err)
		}
	}

}

func doMatch2023(match tba.Match2023, event tba.Event) {
	for n, teamKey := range match.Alliances.Blue.TeamKeys {
		teamNumber, err := strconv.Atoi(strings.Split(teamKey, "frc")[1])

		if err != nil {
			logger.Fatalf("Failed to get team number: %v\n", err)
		}

		participant := scout.Participant{
			Alliance:   "blue",
			TeamNumber: teamNumber,
			Metrics:    data.Metrics2023(match.ScoreBreakdown.Blue, n),
		}

		if err != nil {
			logger.Fatalf("Failed to get team: %v\n", err)
		}

		if err := db.InsertParticipant(participant, adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
			logger.Fatalf("Failed to add participant: %v\n", err)
		}
	}

	for n, teamKey := range match.Alliances.Red.TeamKeys {
		teamNumber, err := strconv.Atoi(strings.Split(teamKey, "frc")[1])

		if err != nil {
			logger.Fatalf("Failed to get team number: %v\n", err)
		}

		participant := scout.Participant{
			Alliance:   "red",
			TeamNumber: teamNumber,
			Metrics:    data.Metrics2023(match.ScoreBreakdown.Red, n),
		}

		if err != nil {
			logger.Fatalf("Failed to get team: %v\n", err)
		}

		if err := db.InsertParticipant(participant, adapter.ToMatch(match), adapter.ToEvent(event)); err != nil {
			logger.Fatalf("Failed to add participant: %v\n", err)
		}
	}

}

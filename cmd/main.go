package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"haydenheroux.github.io/adapter"
	"haydenheroux.github.io/scout"
	"haydenheroux.github.io/tba"
	"haydenheroux.github.io/tbascraper/pkg/memo"
)

const (
	APP_NAME          = "tbascraper"
	DEFAULT_EVENT     = ""
	DEFAULT_API_KEY   = ""
	DEFAULT_SCOUT_URL = ""
)

var (
	eventKey string
	apiKey   string
	scoutURL string
)

func init() {
	flag.StringVar(&eventKey, "eventKey", DEFAULT_EVENT, "Event Key")
	flag.StringVar(&apiKey, "apiKey", DEFAULT_API_KEY, "API Key")
	flag.StringVar(&scoutURL, "scoutURL", DEFAULT_SCOUT_URL, "Scout URL")
}

func main() {
	flag.Parse()

	logger := log.New(os.Stderr, APP_NAME+": ", 0)

	api := tba.New(apiKey)

	db := scout.New(scoutURL)

	event, err := api.GetEvent(eventKey)

	if err != nil {
		logger.Fatalf("Failed to get event: %v\n", err)
	}

	if err := db.InsertEvent(adapter.ToScoutEvent(event)); err != nil {
		logger.Fatalf("Failed to insert event: %v\n", err)
	}

	teams, err := api.GetTeams(eventKey)

	if err != nil {
		logger.Fatalf("Failed to get teams: %v\n", err)
	}

	for _, team := range teams {
		team := adapter.ToScoutTeam(team)

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

		if err := db.AddEvent(adapter.ToScoutEvent(event), season, team); err != nil {
			logger.Fatalf("Failed to add event: %v\n", err)
		}

		memo.Memoize(team, season, robot)
	}

	matchKeys, err := api.GetMatchKeys(eventKey)

	if err != nil {
		logger.Fatalf("Failed to get match keys: %v\n", err)
	}

	for _, matchKey := range matchKeys {
		match, err := api.GetMatch(matchKey, event.Year)

		match2022 := match.(tba.Match2022)

		if err != nil {
			logger.Fatalf("Failed to get match: %v\n", err)
		}

		if err := db.InsertMatch(adapter.ToScoutMatch(match2022), adapter.ToScoutEvent(event)); err != nil {
			logger.Fatalf("Failed to add match: %v\n", err)
		}

		for n, teamKey := range match2022.Alliances.Blue.TeamKeys {
			teamNumber, err := strconv.Atoi(strings.Split(teamKey, "frc")[1])

			if err != nil {
				logger.Fatalf("Failed to get team number: %v\n", err)
			}

			participant := scout.Participant{
				Alliance:   "blue",
				TeamNumber: fmt.Sprint(teamNumber),
				Metrics:    getMetricsFor(match2022.Metrics.Blue, n),
			}

			if err := db.InsertParticipant(participant, adapter.ToScoutMatch(match2022), adapter.ToScoutEvent(event)); err != nil {
				logger.Fatalf("Failed to add participant: %v\n", err)
			}
		}

		for n, teamKey := range match2022.Alliances.Red.TeamKeys {
			teamNumber, err := strconv.Atoi(strings.Split(teamKey, "frc")[1])

			if err != nil {
				logger.Fatalf("Failed to get team number: %v\n", err)
			}

			participant := scout.Participant{
				Alliance:   "red",
				TeamNumber: fmt.Sprint(teamNumber),
				Metrics:    getMetricsFor(match2022.Metrics.Red, n),
			}

			if err := db.InsertParticipant(participant, adapter.ToScoutMatch(match2022), adapter.ToScoutEvent(event)); err != nil {
				logger.Fatalf("Failed to add participant: %v\n", err)
			}
		}

	}
}

func getMetricsFor(m tba.AllianceMetrics2022, robotNumber int) []scout.Metric {
	var autoTaxi string
	var endgameClimb string

	switch robotNumber {
	case 0:
		autoTaxi = m.TaxiRobot1
		endgameClimb = m.EndgameRobot1
	case 1:
		autoTaxi = m.TaxiRobot2
		endgameClimb = m.EndgameRobot2
	case 2:
		autoTaxi = m.TaxiRobot3
		endgameClimb = m.EndgameRobot3
	}

	// TODO Check for double-counting
	autoScoredUpper := m.AutoCargoUpperBlue + m.AutoCargoUpperRed + m.AutoCargoUpperFar + m.AutoCargoUpperNear
	autoScoredLower := m.AutoCargoLowerBlue + m.AutoCargoLowerRed + m.AutoCargoLowerFar + m.AutoCargoLowerNear
	teleopScoredUpper := m.TeleopCargoUpperBlue + m.TeleopCargoUpperRed + m.TeleopCargoUpperFar + m.TeleopCargoUpperNear
	teleopScoredLower := m.TeleopCargoLowerBlue + m.TeleopCargoLowerRed + m.TeleopCargoLowerFar + m.TeleopCargoLowerNear

	var metrics []scout.Metric

	metrics = append(metrics, scout.Metric{Key: "autoTaxi", Value: autoTaxi})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoCargoScored", Value: fmt.Sprint(m.AutoCargoTotal)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoCargoScoredLower", Value: fmt.Sprint(autoScoredLower)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoCargoScoredUpper", Value: fmt.Sprint(autoScoredUpper)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoCargoPoints", Value: fmt.Sprint(m.AutoCargoPoints)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoPoints", Value: fmt.Sprint(m.AutoPoints)})

	metrics = append(metrics, scout.Metric{Key: "allianceTeleopCargoScored", Value: fmt.Sprint(m.TeleopCargoTotal)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopCargoScoredLower", Value: fmt.Sprint(teleopScoredLower)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopCargoScoredUpper", Value: fmt.Sprint(teleopScoredUpper)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopCargoPoints", Value: fmt.Sprint(m.TeleopCargoPoints)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopPoints", Value: fmt.Sprint(m.TeleopPoints)})

	metrics = append(metrics, scout.Metric{Key: "endgameClimb", Value: endgameClimb})

	return metrics
}

package adapter

import (
	"haydenheroux.github.io/scout"
	"haydenheroux.github.io/tba"
)

func ToScoutTeam(team tba.Team) scout.Team {
	return scout.Team{
		Number:  team.Number,
		Name:    team.Nickname,
		Region:  "ne", // TODO
		Seasons: []scout.Season{},
	}
}

func ToScoutEvent(event tba.Event) scout.Event {
	return scout.Event{
		Name:    event.Name,
		Region:  "ne", // TODO event.District.Abbreviation
		Year:    event.Year,
		Week:    0, // TODO
		Matches: []scout.Match{},
	}
}

func ToScoutMatch(match tba.Match2022) scout.Match {
	return scout.Match{
		Number:       match.MatchNumber,
		Type:         getMatchType(match),
		Participants: []scout.Participant{},
	}
}

func getMatchType(match tba.Match2022) string {
	// TODO Take set number into account
	switch match.CompLevel {
	case "q":
		match.MatchNumber += 1000 // TODO
		return "qualification"
	case "qf":
		match.MatchNumber += 10000 // TODO
		return "playoff"
	case "sf":
		match.MatchNumber += 100000 // TODO
		return "playoff"
	case "f":
		match.MatchNumber += 1000000 // TODO
		return "playoff"
	default:
		return "qualification" // TODO
	}
}

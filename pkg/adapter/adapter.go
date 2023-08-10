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
		Set:          match.SetNumber,
		Number:       match.MatchNumber,
		Type:         match.CompLevel,
		Participants: []scout.Participant{},
	}
}

package adapter

import (
	"haydenheroux.github.io/scout"
	"haydenheroux.github.io/tba"
)

func ToTeam(team tba.Team) scout.Team {
	return scout.Team{
		Number:  team.Number,
		Name:    team.Nickname,
		Region:  "ne", // TODO
		Seasons: []scout.Season{},
	}
}

func ToEvent(event tba.Event) scout.Event {
	return scout.Event{
		Name:    event.Name,
		Region:  "ne", // TODO event.District.Abbreviation
		Year:    event.Year,
		Week:    0, // TODO
		Matches: []scout.Match{},
	}
}

func ToMatch(match any) scout.Match {
	switch match.(type) {
	case tba.Match2023:
		match2023 := match.(tba.Match2023)
		return scout.Match{
			Set:          match2023.SetNumber,
			Number:       match2023.MatchNumber,
			Type:         match2023.CompLevel,
			Participants: []scout.Participant{},
		}
	case tba.Match2022:
		match2022 := match.(tba.Match2022)
		return scout.Match{
			Set:          match2022.SetNumber,
			Number:       match2022.MatchNumber,
			Type:         match2022.CompLevel,
			Participants: []scout.Participant{},
		}
	default:
		return scout.Match{}
	}
}

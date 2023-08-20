package adapter

import (
	"github.com/haydenheroux/scout"
	"github.com/haydenheroux/tba"
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
	var region string

	if event.District.Abbreviation != "" {
		region = event.District.Abbreviation
	} else {
		region = "other"
	}

	println(event.Key)

	return scout.Event{
		Code:    event.Key,
		Name:    event.Name,
		Region:  region,
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

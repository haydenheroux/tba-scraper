package adapter

import (
	"haydenheroux.github.io/scout"
	"haydenheroux.github.io/tba"
)

func ToScoutTeam(tbaTeam tba.Team) scout.Team {
	return scout.Team{
		Number:  tbaTeam.Number,
		Name:    tbaTeam.Nickname,
		Region:  scoutRegionOf(tbaTeam),
		Seasons: []scout.Season{},
	}
}

func scoutRegionOf(tbaTeam tba.Team) string {
	// TODO
	return "ne"
}

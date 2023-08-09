package memo

import "haydenheroux.github.io/scout"

var (
	memo map[int]memoized = make(map[int]memoized)
)

type memoized struct {
	team   scout.Team
	season scout.Season
	robot  scout.Robot
}

func Memoize(team scout.Team, season scout.Season, robot scout.Robot) {
	memo[team.Number] = memoized{
		team:   team,
		season: season,
		robot:  robot,
	}
}

func Has(teamNumber int) bool {
	_, exists := memo[teamNumber]
	return exists
}

func Get(teamNumber int) (scout.Team, scout.Season, scout.Robot) {
	if Has(teamNumber) {
		memoized := memo[teamNumber]
		return memoized.team, memoized.season, memoized.robot
	}

	return scout.Team{}, scout.Season{}, scout.Robot{}
}

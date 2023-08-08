package tba

type Team struct {
	Key             string `json:"key"`
	Number          int    `json:"team_number"`
	Nickname        string `json:"nickname"`
	Name            string `json:"name"`
	SchoolName      string `json:"school_name"`
	City            string `json:"city"`
	State           string `json:"state_prov"`
	Country         string `json:"country"`
	Address         string `json:"address"`
	PostalCode      string `json:"postal_code"`
	GoogleMapsPlace string `json:"gmaps_place_id"`
	GoogleMapsURL   string `json:"gmaps_url"`
	Latitude        int    `json:"lat"`
	Longitude       int    `json:"lng"`
	Location        string `json:"location_name"`
	Website         string `json:"website"`
	RookieYear      int    `json:"rookie_year"`
	Motto           string `json:"motto"`
}

type Event struct {
	City      string   `json:"city"`
	Country   string   `json:"country"`
	District  District `json:"district"`
	EndDate   string   `json:"end_date"`
	Code      string   `json:"event_code"`
	Type      int      `json:"event_type"`
	Key       string   `json:"key"`
	Name      string   `json:"name"`
	StartDate string   `json:"start_date"`
	State     string   `json:"state_prov"`
	Year      int      `json:"year"`
}

type District struct {
	Abbreviation string `json:"abbreviation"`
	Name         string `json:"display_name"`
	Key          string `json:"key"`
	Year         int    `json:"year"`
}

type Match2022 struct {
	ActualTime      int           `json:"actual_time"`
	Alliances       Alliances2022 `json:"alliances"`
	CompLevel       string        `json:"comp_level"`
	EventKey        string        `json:"event_key"`
	Key             string        `json:"key"`
	MatchNumber     int           `json:"match_number"`
	PostResultTime  int           `json:"post_result_time"`
	PredictedTime   int           `json:"predicted_time"`
	ScoreBreakdown  Metrics2022   `json:"score_breakdown"`
	SetNumber       int           `json:"set_number"`
	Time            int           `json:"time"`
	Videos          []Videos      `json:"videos"`
	WinningAlliance string        `json:"winning_alliance"`
}
type Alliance2022 struct {
	DqTeamKeys        []any    `json:"dq_team_keys"`
	Score             int      `json:"score"`
	SurrogateTeamKeys []any    `json:"surrogate_team_keys"`
	TeamKeys          []string `json:"team_keys"`
}
type Alliances2022 struct {
	Blue Alliance2022 `json:"blue"`
	Red  Alliance2022 `json:"red"`
}
type AllianceMetrics2022 struct {
	AdjustPoints            int    `json:"adjustPoints"`
	AutoCargoLowerBlue      int    `json:"autoCargoLowerBlue"`
	AutoCargoLowerFar       int    `json:"autoCargoLowerFar"`
	AutoCargoLowerNear      int    `json:"autoCargoLowerNear"`
	AutoCargoLowerRed       int    `json:"autoCargoLowerRed"`
	AutoCargoPoints         int    `json:"autoCargoPoints"`
	AutoCargoTotal          int    `json:"autoCargoTotal"`
	AutoCargoUpperBlue      int    `json:"autoCargoUpperBlue"`
	AutoCargoUpperFar       int    `json:"autoCargoUpperFar"`
	AutoCargoUpperNear      int    `json:"autoCargoUpperNear"`
	AutoCargoUpperRed       int    `json:"autoCargoUpperRed"`
	AutoPoints              int    `json:"autoPoints"`
	AutoTaxiPoints          int    `json:"autoTaxiPoints"`
	CargoBonusRankingPoint  bool   `json:"cargoBonusRankingPoint"`
	EndgamePoints           int    `json:"endgamePoints"`
	EndgameRobot1           string `json:"endgameRobot1"`
	EndgameRobot2           string `json:"endgameRobot2"`
	EndgameRobot3           string `json:"endgameRobot3"`
	FoulCount               int    `json:"foulCount"`
	FoulPoints              int    `json:"foulPoints"`
	HangarBonusRankingPoint bool   `json:"hangarBonusRankingPoint"`
	MatchCargoTotal         int    `json:"matchCargoTotal"`
	QuintetAchieved         bool   `json:"quintetAchieved"`
	Rp                      int    `json:"rp"`
	TaxiRobot1              string `json:"taxiRobot1"`
	TaxiRobot2              string `json:"taxiRobot2"`
	TaxiRobot3              string `json:"taxiRobot3"`
	TechFoulCount           int    `json:"techFoulCount"`
	TeleopCargoLowerBlue    int    `json:"teleopCargoLowerBlue"`
	TeleopCargoLowerFar     int    `json:"teleopCargoLowerFar"`
	TeleopCargoLowerNear    int    `json:"teleopCargoLowerNear"`
	TeleopCargoLowerRed     int    `json:"teleopCargoLowerRed"`
	TeleopCargoPoints       int    `json:"teleopCargoPoints"`
	TeleopCargoTotal        int    `json:"teleopCargoTotal"`
	TeleopCargoUpperBlue    int    `json:"teleopCargoUpperBlue"`
	TeleopCargoUpperFar     int    `json:"teleopCargoUpperFar"`
	TeleopCargoUpperNear    int    `json:"teleopCargoUpperNear"`
	TeleopCargoUpperRed     int    `json:"teleopCargoUpperRed"`
	TeleopPoints            int    `json:"teleopPoints"`
	TotalPoints             int    `json:"totalPoints"`
}
type Metrics2022 struct {
	Blue AllianceMetrics2022 `json:"blue"`
	Red  AllianceMetrics2022 `json:"red"`
}
type Videos struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

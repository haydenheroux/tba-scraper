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
	ActualTime      int                `json:"actual_time"`
	Alliances       Alliances          `json:"alliances"`
	CompLevel       string             `json:"comp_level"`
	EventKey        string             `json:"event_key"`
	Key             string             `json:"key"`
	MatchNumber     int                `json:"match_number"`
	PostResultTime  int                `json:"post_result_time"`
	PredictedTime   int                `json:"predicted_time"`
	ScoreBreakdown  ScoreBreakdown2022 `json:"score_breakdown"`
	SetNumber       int                `json:"set_number"`
	Time            int                `json:"time"`
	Videos          []Videos           `json:"videos"`
	WinningAlliance string             `json:"winning_alliance"`
}

type Alliances struct {
	Blue Alliance `json:"blue"`
	Red  Alliance `json:"red"`
}

type Alliance struct {
	DqTeamKeys        []any    `json:"dq_team_keys"`
	Score             int      `json:"score"`
	SurrogateTeamKeys []any    `json:"surrogate_team_keys"`
	TeamKeys          []string `json:"team_keys"`
}

type ScoreBreakdown2022 struct {
	Blue Scores2022 `json:"blue"`
	Red  Scores2022 `json:"red"`
}

type Scores2022 struct {
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

type Videos struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

type Match2023 struct {
	ActualTime      int                `json:"actual_time"`
	Alliances       Alliances          `json:"alliances"`
	CompLevel       string             `json:"comp_level"`
	EventKey        string             `json:"event_key"`
	Key             string             `json:"key"`
	MatchNumber     int                `json:"match_number"`
	PostResultTime  int                `json:"post_result_time"`
	PredictedTime   int                `json:"predicted_time"`
	ScoreBreakdown  ScoreBreakdown2023 `json:"score_breakdown"`
	SetNumber       int                `json:"set_number"`
	Time            int                `json:"time"`
	Videos          []Videos           `json:"videos"`
	WinningAlliance string             `json:"winning_alliance"`
}

type ScoreBreakdown2023 struct {
	Blue Scores2023 `json:"blue"`
	Red  Scores2023 `json:"red"`
}

type Scores2023 struct {
	ActivationBonusAchieved     bool          `json:"activationBonusAchieved"`
	AdjustPoints                int           `json:"adjustPoints"`
	AutoBridgeState             string        `json:"autoBridgeState"`
	AutoChargeStationPoints     int           `json:"autoChargeStationPoints"`
	AutoChargeStationRobot1     string        `json:"autoChargeStationRobot1"`
	AutoChargeStationRobot2     string        `json:"autoChargeStationRobot2"`
	AutoChargeStationRobot3     string        `json:"autoChargeStationRobot3"`
	AutoCommunity               Community2023 `json:"autoCommunity"`
	AutoDocked                  bool          `json:"autoDocked"`
	AutoGamePieceCount          int           `json:"autoGamePieceCount"`
	AutoGamePiecePoints         int           `json:"autoGamePiecePoints"`
	AutoMobilityPoints          int           `json:"autoMobilityPoints"`
	AutoPoints                  int           `json:"autoPoints"`
	CoopGamePieceCount          int           `json:"coopGamePieceCount"`
	CoopertitionCriteriaMet     bool          `json:"coopertitionCriteriaMet"`
	EndGameBridgeState          string        `json:"endGameBridgeState"`
	EndGameChargeStationPoints  int           `json:"endGameChargeStationPoints"`
	EndGameChargeStationRobot1  string        `json:"endGameChargeStationRobot1"`
	EndGameChargeStationRobot2  string        `json:"endGameChargeStationRobot2"`
	EndGameChargeStationRobot3  string        `json:"endGameChargeStationRobot3"`
	EndGameParkPoints           int           `json:"endGameParkPoints"`
	FoulCount                   int           `json:"foulCount"`
	FoulPoints                  int           `json:"foulPoints"`
	LinkPoints                  int           `json:"linkPoints"`
	Links                       []Link2023    `json:"links"`
	MobilityRobot1              string        `json:"mobilityRobot1"`
	MobilityRobot2              string        `json:"mobilityRobot2"`
	MobilityRobot3              string        `json:"mobilityRobot3"`
	Rp                          int           `json:"rp"`
	SustainabilityBonusAchieved bool          `json:"sustainabilityBonusAchieved"`
	TechFoulCount               int           `json:"techFoulCount"`
	TeleopCommunity             Community2023 `json:"teleopCommunity"`
	TeleopGamePieceCount        int           `json:"teleopGamePieceCount"`
	TeleopGamePiecePoints       int           `json:"teleopGamePiecePoints"`
	TeleopPoints                int           `json:"teleopPoints"`
	TotalChargeStationPoints    int           `json:"totalChargeStationPoints"`
	TotalPoints                 int           `json:"totalPoints"`
}

type Community2023 struct {
	Bottom []string `json:"B"`
	Middle []string `json:"M"`
	Top    []string `json:"T"`
}

type Link2023 struct {
	Nodes []int  `json:"nodes"`
	Row   string `json:"row"`
}

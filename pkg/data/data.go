package data

import (
	"fmt"

	"haydenheroux.github.io/scout"
	"haydenheroux.github.io/tba"
)

func Metrics2022(scores tba.Scores2022, robotIndex int) []scout.Metric {
	var autoTaxi string
	var endgameClimb string

	switch robotIndex {
	case 0:
		autoTaxi = scores.TaxiRobot1
		endgameClimb = scores.EndgameRobot1
	case 1:
		autoTaxi = scores.TaxiRobot2
		endgameClimb = scores.EndgameRobot2
	case 2:
		autoTaxi = scores.TaxiRobot3
		endgameClimb = scores.EndgameRobot3
	}

	autoScoredUpper := scores.AutoCargoUpperBlue + scores.AutoCargoUpperRed + scores.AutoCargoUpperFar + scores.AutoCargoUpperNear
	autoScoredLower := scores.AutoCargoLowerBlue + scores.AutoCargoLowerRed + scores.AutoCargoLowerFar + scores.AutoCargoLowerNear
	teleopScoredUpper := scores.TeleopCargoUpperBlue + scores.TeleopCargoUpperRed + scores.TeleopCargoUpperFar + scores.TeleopCargoUpperNear
	teleopScoredLower := scores.TeleopCargoLowerBlue + scores.TeleopCargoLowerRed + scores.TeleopCargoLowerFar + scores.TeleopCargoLowerNear

	var metrics []scout.Metric

	metrics = append(metrics, scout.Metric{Key: "autoTaxi", Value: autoTaxi})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoCargoScored", Value: fmt.Sprint(scores.AutoCargoTotal)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoCargoScoredLower", Value: fmt.Sprint(autoScoredLower)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoCargoScoredUpper", Value: fmt.Sprint(autoScoredUpper)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoCargoPoints", Value: fmt.Sprint(scores.AutoCargoPoints)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoPoints", Value: fmt.Sprint(scores.AutoPoints)})

	metrics = append(metrics, scout.Metric{Key: "allianceTeleopCargoScored", Value: fmt.Sprint(scores.TeleopCargoTotal)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopCargoScoredLower", Value: fmt.Sprint(teleopScoredLower)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopCargoScoredUpper", Value: fmt.Sprint(teleopScoredUpper)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopCargoPoints", Value: fmt.Sprint(scores.TeleopCargoPoints)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopPoints", Value: fmt.Sprint(scores.TeleopPoints)})

	metrics = append(metrics, scout.Metric{Key: "endgameClimb", Value: endgameClimb})

	return metrics
}

func Metrics2023(scores tba.Scores2023, robotIndex int) []scout.Metric {
	var autoMobility string
	var autoChargeStation string
	var endgameChargeStation string

	switch robotIndex {
	case 0:
		autoMobility = scores.MobilityRobot1
		autoChargeStation = scores.AutoChargeStationRobot1
		endgameChargeStation = scores.EndGameChargeStationRobot1
	case 1:
		autoMobility = scores.MobilityRobot2
		autoChargeStation = scores.AutoChargeStationRobot2
		endgameChargeStation = scores.EndGameChargeStationRobot2
	case 2:
		autoMobility = scores.MobilityRobot3
		autoChargeStation = scores.AutoChargeStationRobot3
		endgameChargeStation = scores.EndGameChargeStationRobot3
	}

	var metrics []scout.Metric

	metrics = append(metrics, scout.Metric{Key: "autoMobility", Value: autoMobility})
	metrics = append(metrics, scout.Metric{Key: "autoChargeStation", Value: autoChargeStation})

	allianceAutoCubesBottom, allianceAutoCubesMiddle, allianceAutoCubesTop := countCommunity(scores.AutoCommunity, IS_CUBE)
	allianceAutoCubesTotal := allianceAutoCubesBottom + allianceAutoCubesMiddle + allianceAutoCubesTop

	metrics = append(metrics, scout.Metric{Key: "allianceAutoCubesBottom", Value: fmt.Sprint(allianceAutoCubesBottom)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoCubesMiddle", Value: fmt.Sprint(allianceAutoCubesMiddle)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoCubesTop", Value: fmt.Sprint(allianceAutoCubesTop)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoCubesTotal", Value: fmt.Sprint(allianceAutoCubesTotal)})

	allianceAutoConesBottom, allianceAutoConesMiddle, allianceAutoConesTop := countCommunity(scores.AutoCommunity, IS_CONE)
	allianceAutoConesTotal := allianceAutoConesBottom + allianceAutoConesMiddle + allianceAutoConesTop

	metrics = append(metrics, scout.Metric{Key: "allianceAutoConesBottom", Value: fmt.Sprint(allianceAutoConesBottom)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoConesMiddle", Value: fmt.Sprint(allianceAutoConesMiddle)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoConesTop", Value: fmt.Sprint(allianceAutoConesTop)})
	metrics = append(metrics, scout.Metric{Key: "allianceAutoConesTotal", Value: fmt.Sprint(allianceAutoConesTotal)})

	allianceTeleopCubesBottom, allianceTeleopCubesMiddle, allianceTeleopCubesTop := countCommunity(scores.TeleopCommunity, IS_CUBE)
	allianceTeleopCubesTotal := allianceTeleopCubesBottom + allianceTeleopCubesMiddle + allianceTeleopCubesTop

	metrics = append(metrics, scout.Metric{Key: "allianceTeleopCubesBottom", Value: fmt.Sprint(allianceTeleopCubesBottom)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopCubesMiddle", Value: fmt.Sprint(allianceTeleopCubesMiddle)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopCubesTop", Value: fmt.Sprint(allianceTeleopCubesTop)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopCubesTotal", Value: fmt.Sprint(allianceTeleopCubesTotal)})

	allianceTeleopConesBottom, allianceTeleopConesMiddle, allianceTeleopConesTop := countCommunity(scores.TeleopCommunity, IS_CONE)
	allianceTeleopConesTotal := allianceTeleopConesBottom + allianceTeleopConesMiddle + allianceTeleopConesTop

	metrics = append(metrics, scout.Metric{Key: "allianceTeleopConesBottom", Value: fmt.Sprint(allianceTeleopConesBottom)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopConesMiddle", Value: fmt.Sprint(allianceTeleopConesMiddle)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopConesTop", Value: fmt.Sprint(allianceTeleopConesTop)})
	metrics = append(metrics, scout.Metric{Key: "allianceTeleopConesTotal", Value: fmt.Sprint(allianceTeleopConesTotal)})

	allianceLinksBottom, allianceLinksMiddle, allianceLinksTop := countLinks(scores.Links)
	allianceLinksTotal := allianceLinksBottom + allianceLinksMiddle + allianceLinksTop

	metrics = append(metrics, scout.Metric{Key: "allianceLinksBottom", Value: fmt.Sprint(allianceLinksBottom)})
	metrics = append(metrics, scout.Metric{Key: "allianceLinksMiddle", Value: fmt.Sprint(allianceLinksMiddle)})
	metrics = append(metrics, scout.Metric{Key: "allianceLinksTop", Value: fmt.Sprint(allianceLinksTop)})
	metrics = append(metrics, scout.Metric{Key: "allianceLinksTotal", Value: fmt.Sprint(allianceLinksTotal)})

	metrics = append(metrics, scout.Metric{Key: "endgameChargeStation", Value: endgameChargeStation})

	return metrics
}

var (
	IS_CUBE = func(s string) bool { return s == "Cube" }
	IS_CONE = func(s string) bool { return s == "Cone" }
)

func countCommunity(community tba.Community2023, pieceTest func(string) bool) (int, int, int) {
	bottom := filterByPiece(community.Bottom, pieceTest)
	middle := filterByPiece(community.Middle, pieceTest)
	top := filterByPiece(community.Top, pieceTest)

	return len(bottom), len(middle), len(top)
}

func filterByPiece(pieces []string, pieceTest func(string) bool) []string {
	filtered := make([]string, 0)

	for _, xpiece := range pieces {
		if pieceTest(xpiece) {
			filtered = append(filtered, xpiece)
		}
	}

	return filtered
}

var (
	BOTTOM_ROW = func(link tba.Link2023) bool { return link.Row == "Bottom" }
	MIDDLE_ROW = func(link tba.Link2023) bool { return link.Row == "Mid" }
	TOP_ROW    = func(link tba.Link2023) bool { return link.Row == "Top" }
)

func countLinks(links []tba.Link2023) (int, int, int) {
	bottom := filterByRow(links, BOTTOM_ROW)
	middle := filterByRow(links, MIDDLE_ROW)
	top := filterByRow(links, TOP_ROW)

	return len(bottom), len(middle), len(top)
}

func filterByRow(links []tba.Link2023, rowTest func(tba.Link2023) bool) []tba.Link2023 {
	filtered := make([]tba.Link2023, 0)

	for _, link := range links {
		if rowTest(link) {
			filtered = append(filtered, link)
		}
	}

	return filtered
}

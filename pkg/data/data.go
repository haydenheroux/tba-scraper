package data

import (
	"fmt"

	"github.com/haydenheroux/scout"
	"github.com/haydenheroux/tba"
)

func AllianceMetrics2022(scores tba.Scores2022) []scout.Metric {
	autoScoredUpper := scores.AutoCargoUpperBlue + scores.AutoCargoUpperRed + scores.AutoCargoUpperFar + scores.AutoCargoUpperNear
	autoScoredLower := scores.AutoCargoLowerBlue + scores.AutoCargoLowerRed + scores.AutoCargoLowerFar + scores.AutoCargoLowerNear
	teleopScoredUpper := scores.TeleopCargoUpperBlue + scores.TeleopCargoUpperRed + scores.TeleopCargoUpperFar + scores.TeleopCargoUpperNear
	teleopScoredLower := scores.TeleopCargoLowerBlue + scores.TeleopCargoLowerRed + scores.TeleopCargoLowerFar + scores.TeleopCargoLowerNear

	var metrics []scout.Metric

	metrics = append(metrics, scout.Metric{Key: "autoCargoScored", Value: fmt.Sprint(scores.AutoCargoTotal)})
	metrics = append(metrics, scout.Metric{Key: "autoCargoScoredLower", Value: fmt.Sprint(autoScoredLower)})
	metrics = append(metrics, scout.Metric{Key: "autoCargoScoredUpper", Value: fmt.Sprint(autoScoredUpper)})
	metrics = append(metrics, scout.Metric{Key: "autoCargoPoints", Value: fmt.Sprint(scores.AutoCargoPoints)})
	metrics = append(metrics, scout.Metric{Key: "autoPoints", Value: fmt.Sprint(scores.AutoPoints)})

	metrics = append(metrics, scout.Metric{Key: "teleopCargoScored", Value: fmt.Sprint(scores.TeleopCargoTotal)})
	metrics = append(metrics, scout.Metric{Key: "teleopCargoScoredLower", Value: fmt.Sprint(teleopScoredLower)})
	metrics = append(metrics, scout.Metric{Key: "teleopCargoScoredUpper", Value: fmt.Sprint(teleopScoredUpper)})
	metrics = append(metrics, scout.Metric{Key: "teleopCargoPoints", Value: fmt.Sprint(scores.TeleopCargoPoints)})
	metrics = append(metrics, scout.Metric{Key: "teleopPoints", Value: fmt.Sprint(scores.TeleopPoints)})

	return metrics
}

func AllianceMetrics2023(scores tba.Scores2023) []scout.Metric {
	var metrics []scout.Metric

	allianceAutoCubesBottom, allianceAutoCubesMiddle, allianceAutoCubesTop := countCommunity(scores.AutoCommunity, IS_CUBE)
	allianceAutoCubesTotal := allianceAutoCubesBottom + allianceAutoCubesMiddle + allianceAutoCubesTop

	metrics = append(metrics, scout.Metric{Key: "autoCubesBottom", Value: fmt.Sprint(allianceAutoCubesBottom)})
	metrics = append(metrics, scout.Metric{Key: "autoCubesMiddle", Value: fmt.Sprint(allianceAutoCubesMiddle)})
	metrics = append(metrics, scout.Metric{Key: "autoCubesTop", Value: fmt.Sprint(allianceAutoCubesTop)})
	metrics = append(metrics, scout.Metric{Key: "autoCubesTotal", Value: fmt.Sprint(allianceAutoCubesTotal)})

	allianceAutoConesBottom, allianceAutoConesMiddle, allianceAutoConesTop := countCommunity(scores.AutoCommunity, IS_CONE)
	allianceAutoConesTotal := allianceAutoConesBottom + allianceAutoConesMiddle + allianceAutoConesTop

	metrics = append(metrics, scout.Metric{Key: "autoConesBottom", Value: fmt.Sprint(allianceAutoConesBottom)})
	metrics = append(metrics, scout.Metric{Key: "autoConesMiddle", Value: fmt.Sprint(allianceAutoConesMiddle)})
	metrics = append(metrics, scout.Metric{Key: "autoConesTop", Value: fmt.Sprint(allianceAutoConesTop)})
	metrics = append(metrics, scout.Metric{Key: "autoConesTotal", Value: fmt.Sprint(allianceAutoConesTotal)})

	allianceTeleopCubesBottom, allianceTeleopCubesMiddle, allianceTeleopCubesTop := countCommunity(scores.TeleopCommunity, IS_CUBE)
	allianceTeleopCubesTotal := allianceTeleopCubesBottom + allianceTeleopCubesMiddle + allianceTeleopCubesTop

	metrics = append(metrics, scout.Metric{Key: "teleopCubesBottom", Value: fmt.Sprint(allianceTeleopCubesBottom)})
	metrics = append(metrics, scout.Metric{Key: "teleopCubesMiddle", Value: fmt.Sprint(allianceTeleopCubesMiddle)})
	metrics = append(metrics, scout.Metric{Key: "teleopCubesTop", Value: fmt.Sprint(allianceTeleopCubesTop)})
	metrics = append(metrics, scout.Metric{Key: "teleopCubesTotal", Value: fmt.Sprint(allianceTeleopCubesTotal)})

	allianceTeleopConesBottom, allianceTeleopConesMiddle, allianceTeleopConesTop := countCommunity(scores.TeleopCommunity, IS_CONE)
	allianceTeleopConesTotal := allianceTeleopConesBottom + allianceTeleopConesMiddle + allianceTeleopConesTop

	metrics = append(metrics, scout.Metric{Key: "teleopConesBottom", Value: fmt.Sprint(allianceTeleopConesBottom)})
	metrics = append(metrics, scout.Metric{Key: "teleopConesMiddle", Value: fmt.Sprint(allianceTeleopConesMiddle)})
	metrics = append(metrics, scout.Metric{Key: "teleopConesTop", Value: fmt.Sprint(allianceTeleopConesTop)})
	metrics = append(metrics, scout.Metric{Key: "teleopConesTotal", Value: fmt.Sprint(allianceTeleopConesTotal)})

	allianceLinksBottom, allianceLinksMiddle, allianceLinksTop := countLinks(scores.Links)
	allianceLinksTotal := allianceLinksBottom + allianceLinksMiddle + allianceLinksTop

	metrics = append(metrics, scout.Metric{Key: "linksBottom", Value: fmt.Sprint(allianceLinksBottom)})
	metrics = append(metrics, scout.Metric{Key: "linksMiddle", Value: fmt.Sprint(allianceLinksMiddle)})
	metrics = append(metrics, scout.Metric{Key: "linksTop", Value: fmt.Sprint(allianceLinksTop)})
	metrics = append(metrics, scout.Metric{Key: "linksTotal", Value: fmt.Sprint(allianceLinksTotal)})

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

package data

import (
	"fmt"

	"github.com/haydenheroux/tba"
)

func AllianceMetrics2022(scores tba.Scores2022) map[string]string {
	autoScoredUpper := scores.AutoCargoUpperBlue + scores.AutoCargoUpperRed + scores.AutoCargoUpperFar + scores.AutoCargoUpperNear
	autoScoredLower := scores.AutoCargoLowerBlue + scores.AutoCargoLowerRed + scores.AutoCargoLowerFar + scores.AutoCargoLowerNear
	teleopScoredUpper := scores.TeleopCargoUpperBlue + scores.TeleopCargoUpperRed + scores.TeleopCargoUpperFar + scores.TeleopCargoUpperNear
	teleopScoredLower := scores.TeleopCargoLowerBlue + scores.TeleopCargoLowerRed + scores.TeleopCargoLowerFar + scores.TeleopCargoLowerNear

	var metrics map[string]string

	metrics["autoCargoScored"] = fmt.Sprint(scores.AutoCargoTotal)
	metrics["autoCargoScoredLower"] = fmt.Sprint(autoScoredLower)
	metrics["autoCargoScoredUpper"] = fmt.Sprint(autoScoredUpper)
	metrics["autoCargoPoints"] = fmt.Sprint(scores.AutoCargoPoints)
	metrics["autoPoints"] = fmt.Sprint(scores.AutoPoints)

	metrics["teleopCargoScored"] = fmt.Sprint(scores.TeleopCargoTotal)
	metrics["teleopCargoScoredLower"] = fmt.Sprint(teleopScoredLower)
	metrics["teleopCargoScoredUpper"] = fmt.Sprint(teleopScoredUpper)
	metrics["teleopCargoPoints"] = fmt.Sprint(scores.TeleopCargoPoints)
	metrics["teleopPoints"] = fmt.Sprint(scores.TeleopPoints)

	return metrics
}

func AllianceMetrics2023(scores tba.Scores2023) map[string]string {
	var metrics map[string]string

	allianceAutoCubesBottom, allianceAutoCubesMiddle, allianceAutoCubesTop := countCommunity(scores.AutoCommunity, IS_CUBE)
	allianceAutoCubesTotal := allianceAutoCubesBottom + allianceAutoCubesMiddle + allianceAutoCubesTop

	metrics["autoCubesBottom"] = fmt.Sprint(allianceAutoCubesBottom)
	metrics["autoCubesMiddle"] = fmt.Sprint(allianceAutoCubesMiddle)
	metrics["autoCubesTop"] = fmt.Sprint(allianceAutoCubesTop)
	metrics["autoCubesTotal"] = fmt.Sprint(allianceAutoCubesTotal)

	allianceAutoConesBottom, allianceAutoConesMiddle, allianceAutoConesTop := countCommunity(scores.AutoCommunity, IS_CONE)
	allianceAutoConesTotal := allianceAutoConesBottom + allianceAutoConesMiddle + allianceAutoConesTop

	metrics["autoConesBottom"] = fmt.Sprint(allianceAutoConesBottom)
	metrics["autoConesMiddle"] = fmt.Sprint(allianceAutoConesMiddle)
	metrics["autoConesTop"] = fmt.Sprint(allianceAutoConesTop)
	metrics["autoConesTotal"] = fmt.Sprint(allianceAutoConesTotal)

	allianceTeleopCubesBottom, allianceTeleopCubesMiddle, allianceTeleopCubesTop := countCommunity(scores.TeleopCommunity, IS_CUBE)
	allianceTeleopCubesTotal := allianceTeleopCubesBottom + allianceTeleopCubesMiddle + allianceTeleopCubesTop

	metrics["teleopCubesBottom"] = fmt.Sprint(allianceTeleopCubesBottom)
	metrics["teleopCubesMiddle"] = fmt.Sprint(allianceTeleopCubesMiddle)
	metrics["teleopCubesTop"] = fmt.Sprint(allianceTeleopCubesTop)
	metrics["teleopCubesTotal"] = fmt.Sprint(allianceTeleopCubesTotal)

	allianceTeleopConesBottom, allianceTeleopConesMiddle, allianceTeleopConesTop := countCommunity(scores.TeleopCommunity, IS_CONE)
	allianceTeleopConesTotal := allianceTeleopConesBottom + allianceTeleopConesMiddle + allianceTeleopConesTop

	metrics["teleopConesBottom"] = fmt.Sprint(allianceTeleopConesBottom)
	metrics["teleopConesMiddle"] = fmt.Sprint(allianceTeleopConesMiddle)
	metrics["teleopConesTop"] = fmt.Sprint(allianceTeleopConesTop)
	metrics["teleopConesTotal"] = fmt.Sprint(allianceTeleopConesTotal)

	allianceLinksBottom, allianceLinksMiddle, allianceLinksTop := countLinks(scores.Links)
	allianceLinksTotal := allianceLinksBottom + allianceLinksMiddle + allianceLinksTop

	metrics["linksBottom"] = fmt.Sprint(allianceLinksBottom)
	metrics["linksMiddle"] = fmt.Sprint(allianceLinksMiddle)
	metrics["linksTop"] = fmt.Sprint(allianceLinksTop)
	metrics["linksTotal"] = fmt.Sprint(allianceLinksTotal)

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

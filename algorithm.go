package main

import (
	"log"
	"math"
)

/*
Calculation of risk score for resources given measured data
*/

// computeScore : compute score given usage statistics
// - risk = [ average + margin * stDev^{1/sensitivity} ] / 2
// - score = ( 1 - risk ) * maxScore
func ComputeScore(rs *ResourceStats, margin float64, sensitivity float64) float64 {
	if rs.Capacity <= 0 {
		log.Printf("ERROR: Invalid resource capacity %v", rs.Capacity)
		return 0
	}

	// make sure values are within bounds
	rs.Req = math.Max(rs.Req, 0)
	rs.UsedAvg = math.Max(math.Min(rs.UsedAvg, rs.Capacity), 0)
	rs.UsedStdev = math.Max(math.Min(rs.UsedStdev, rs.Capacity), 0)

	// calculate average and deviation factors
	mu, sigma := GetMuSigma(rs)

	// apply root power
	if sensitivity >= 0 {
		sigma = math.Pow(sigma, 1/sensitivity)
	}
	// apply multiplier
	sigma *= margin
	sigma = math.Max(math.Min(sigma, 1), 0)

	// evaluate overall risk factor
	risk := (mu + sigma) / 2

	log.Printf("INFO: Evaluating risk factor, mu %v, sigma %v, margin %v, sensitivity %v, risk %v", mu, sigma, margin, sensitivity, risk)

	return (1. - risk) * 100
}

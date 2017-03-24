package main

import (
//"log"
)

/*
	Process has 5 steps:
	1. Preprocessing: From the measurement list & the samples list creates:
		- measurementsByPlaneType map[planeType][]Measurement
		- sampleByID map[id]Sample
	2. Distribution making: From measurementByPlaneType creates:
		- caracteristicsDistributionsByPlaneType map[planeType]CaracteristicDistributions (struct including distributions each caracteristics)
	3. Calculate: calculate probabilities of each type for each sample
		- probabilityOfPlaneTypeByID map[ID]map[planeType]
	4. Final: find the maximum probability
		- return a map[id]type
	5. Format
*/
func process(input *Input) (output *Output) {
	// STEP 1: Preprocessing
	measurementsByPlaneTypeLocal := preprocessMeasurements(input.Measurements)
	sampleBySampleIDLocal := preprocessSamples(input.Samples)

	// STEP 2: Get stats of measurements
	caracteristicsStatsByPlaneTypeLocal := calculateStats(measurementsByPlaneTypeLocal)

	// STEP 3: Calculate weighted distance to mean for each
	distanceToMeanOfPlaneTypeBySampleIDLocal := calculateDistanceToMean(caracteristicsStatsByPlaneTypeLocal, sampleBySampleIDLocal)

	// STEP 4: Find the minimum weighted distance to mean
	planeTypeBySampleIDLocal := match(distanceToMeanOfPlaneTypeBySampleIDLocal)

	// STEP 5: Fomatting
	output = format(planeTypeBySampleIDLocal)

	return
}

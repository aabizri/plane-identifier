package main

import (
	"math"
)

type meanVarianceCouple struct {
	mean     float64
	variance float64
}

type caracteristicsStats struct {
	noiseLevel    meanVarianceCouple
	brakeDistance meanVarianceCouple
	vibrations    meanVarianceCouple
}

type caracteristicsStatsByPlaneType map[PlaneType]caracteristicsStats

func calculateStats(measurementsByPlaneTypeLocal measurementsByPlaneType) caracteristicsStatsByPlaneType {
	var caracteristicsStatsByPlaneTypeLocal caracteristicsStatsByPlaneType = make(caracteristicsStatsByPlaneType, len(measurementsByPlaneTypeLocal))
	for planeType, measurements := range measurementsByPlaneTypeLocal {
		// Calculate the means for each caracteristic
		noiseLevelMean, brakeDistanceMean, vibrationsMean := calculateMeans(measurements)
		// Calculate the variance
		noiseLevelVariance, brakeDistanceVariance, vibrationsVariance := calculateVariance(measurements, noiseLevelMean, brakeDistanceMean, vibrationsMean)

		// Add the caracteristic stats
		var (
			noiseLevelCouple         meanVarianceCouple  = meanVarianceCouple{noiseLevelMean, noiseLevelVariance}
			brakeDistanceCouple      meanVarianceCouple  = meanVarianceCouple{brakeDistanceMean, brakeDistanceVariance}
			vibrationsCouple         meanVarianceCouple  = meanVarianceCouple{vibrationsMean, vibrationsVariance}
			caracteristicsStatsLocal caracteristicsStats = caracteristicsStats{noiseLevelCouple, brakeDistanceCouple, vibrationsCouple}
		)

		logger.Printf("We got for plane %s:\n\tNoise Level mean:\t%f\n\tNoise Level variance:\t%f\n\tBrake Distance mean:\t%f\n\tBrake Distance variance:\t%f\n\tVibrations mean:\t%f\n\tVibrations variance:\t%f\n", planeType, noiseLevelMean, noiseLevelVariance, brakeDistanceMean, brakeDistanceVariance, vibrationsMean, vibrationsVariance)

		// Add it to the map
		caracteristicsStatsByPlaneTypeLocal[planeType] = caracteristicsStatsLocal
	}
	return caracteristicsStatsByPlaneTypeLocal
}

func calculateMeans(measurementList []Measurement) (noiseLevelMean, brakeDistanceMean, vibrationsMean float64) {
	measurementsNumber := float64(len(measurementList))
	var (
		noiseLevelSum    float64
		brakeDistanceSum float64
		vibrationsSum    float64
	)
	// Sum it all
	for _, m := range measurementList {
		noiseLevelSum += m.NoiseLevel
		brakeDistanceSum += m.BrakeDistance
		vibrationsSum += m.Vibrations
	}
	// Divide it
	noiseLevelMean = noiseLevelSum / measurementsNumber
	brakeDistanceMean = brakeDistanceSum / measurementsNumber
	vibrationsMean = vibrationsSum / measurementsNumber

	// Return
	return
}

func calculateVariance(measurementList []Measurement, noiseLevelMean, brakeDistanceMean, vibrationsMean float64) (noiseLevelVariance, brakeDistanceVariance, vibrationsVariance float64) {
	measurementsNumber := float64(len(measurementList))
	var (
		noiseLevelTmp    float64
		brakeDistanceTmp float64
		vibrationsTmp    float64
	)
	// Intermediary calculation
	for _, m := range measurementList {
		noiseLevelTmp += math.Pow(m.NoiseLevel-noiseLevelMean, 2)
		brakeDistanceTmp += math.Pow(m.BrakeDistance-brakeDistanceMean, 2)
		vibrationsTmp += math.Pow(m.Vibrations-vibrationsMean, 2)
	}
	// Divide it
	noiseLevelVariance = noiseLevelTmp / measurementsNumber
	brakeDistanceVariance = brakeDistanceTmp / measurementsNumber
	vibrationsVariance = vibrationsTmp / measurementsNumber

	// Return
	return
}

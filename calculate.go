package main

import "math"

type distanceToMeanOfPlaneTypeBySampleID map[SampleID]map[PlaneType]float64

//const interval float64 = 0.5

func calculateDistanceToMean(caracteristicsStatsByPlaneTypeLocal caracteristicsStatsByPlaneType, sampleBySampleIDLocal sampleBySampleID) distanceToMeanOfPlaneTypeBySampleID {

	var distanceToMeanOfPlaneTypeBySampleIDLocal distanceToMeanOfPlaneTypeBySampleID = make(distanceToMeanOfPlaneTypeBySampleID, len(sampleBySampleIDLocal))

	// For each caracteristic, get probability for this sample
	for sampleID, sample := range sampleBySampleIDLocal {
		// Init the map
		distanceToMeanOfPlaneTypeBySampleIDLocal[sampleID] = make(map[PlaneType]float64, len(caracteristicsStatsByPlaneTypeLocal))

		// Compare it to every type
		for planeType, caracteristicsStats := range caracteristicsStatsByPlaneTypeLocal {
			var (
				noiseLevelDistanceToMean    float64 = math.Abs(caracteristicsStats.noiseLevel.mean-sample.NoiseLevel) / math.Sqrt(caracteristicsStats.noiseLevel.variance)
				brakeDistanceDistanceToMean float64 = math.Abs(caracteristicsStats.brakeDistance.mean-sample.BrakeDistance) / math.Sqrt(caracteristicsStats.brakeDistance.variance)
				vibrationsDistanceToMean    float64 = math.Abs(caracteristicsStats.vibrations.mean-sample.Vibrations) / math.Sqrt(caracteristicsStats.vibrations.variance)
				meanDistanceToMean          float64 = (noiseLevelDistanceToMean + brakeDistanceDistanceToMean + vibrationsDistanceToMean) / 3
			)
			distanceToMeanOfPlaneTypeBySampleIDLocal[sampleID][planeType] = meanDistanceToMean
			logger.Printf("For sample %d and plane %s we got these distance to mean:\n\tNoise Level:\t%f\n\tBrake Distance:\t%f\n\tVibrations:\t%f\n\tMean:\t%f\n", sampleID, planeType, noiseLevelDistanceToMean, brakeDistanceDistanceToMean, vibrationsDistanceToMean, meanDistanceToMean)
		}
	}
	return distanceToMeanOfPlaneTypeBySampleIDLocal
}

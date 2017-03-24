package main

type planeTypeBySampleID map[SampleID]PlaneType

func match(distanceToMeanOfPlaneTypeBySampleIDLocal distanceToMeanOfPlaneTypeBySampleID) planeTypeBySampleID {
	var planeTypeBySampleIDLocal planeTypeBySampleID = make(planeTypeBySampleID, len(distanceToMeanOfPlaneTypeBySampleIDLocal))

	for sampleID, probabilityOfPlaneTypes := range distanceToMeanOfPlaneTypeBySampleIDLocal {
		// Find the minimum match within the list of plane types
		var (
			matchDistanceToMean float64
			matchPlaneType      PlaneType
		)
		for planeType, distanceToMean := range probabilityOfPlaneTypes {
			if distanceToMean < matchDistanceToMean || matchDistanceToMean == 0 {
				matchDistanceToMean = distanceToMean
				matchPlaneType = planeType
			}
		}
		planeTypeBySampleIDLocal[sampleID] = matchPlaneType
	}

	return planeTypeBySampleIDLocal
}

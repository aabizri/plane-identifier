package main

// measurementsByPlaneType associates the measurements associated to a planeType to each planeType
type measurementsByPlaneType map[PlaneType][]Measurement

// preprocessMeasurements associates the measurements associated to a planeType to each planeType
func preprocessMeasurements(measurementList []Measurement) measurementsByPlaneType {
	// As we don't yet know the number of different plane types, we'll have to append
	var measurementsByPlaneTypeLocal measurementsByPlaneType = make(measurementsByPlaneType, 0)

	// For each measurement, check the plane type, if it doesn't yet exist add a new slice and append it to it, if it does then just append
	for _, measurement := range measurementList {
		// If this type hasn't yet been added to the map, then create the hosted slice
		if _, ok := measurementsByPlaneTypeLocal[measurement.Type]; !ok {
			measurementsByPlaneTypeLocal[measurement.Type] = make([]Measurement, 0)
		}

		// Now add the measurement to the slice
		measurementsByPlaneTypeLocal[measurement.Type] = append(measurementsByPlaneTypeLocal[measurement.Type], measurement)
	}

	return measurementsByPlaneTypeLocal
}

type sampleBySampleID map[SampleID]Sample

// preprocessSamples associated the sample to a sample ID
func preprocessSamples(sampleList []Sample) sampleBySampleID {
	var sampleBySampleIDLocal sampleBySampleID = make(sampleBySampleID, len(sampleList))
	for _, s := range sampleList {
		sampleBySampleIDLocal[s.ID] = s
	}
	return sampleBySampleIDLocal
}

package main

// format creates an output
func format(planeTypeBySampleIDLocal planeTypeBySampleID) *Output {
	var output *Output = &Output{}
	output.Results = make([]Result, len(planeTypeBySampleIDLocal))
	i := 0
	for sampleID, planeType := range planeTypeBySampleIDLocal {
		var result Result = Result{
			ID:   sampleID,
			Type: planeType,
		}
		output.Results[i] = result
		i++
	}
	return output
}

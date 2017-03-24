// TODO: Add feasability
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

func init() {
	flag.Parse()
}

type Input struct {
	Measurements []Measurement `json:"measurements"`
	Samples      []Sample      `json:"samples"`
}

type PlaneType string

type Measurement struct {
	Type          PlaneType `json:"type"`
	NoiseLevel    float64   `json:"noise-level"`
	BrakeDistance float64   `json:"brake-distance"`
	Vibrations    float64   `json:"vibrations"`
}

type SampleID uint

type Sample struct {
	ID            SampleID `json:"id"`
	NoiseLevel    float64  `json:"noise-level"`
	BrakeDistance float64  `json:"brake-distance"`
	Vibrations    float64  `json:"vibrations"`
}

type Output struct {
	Results []Result `json:"result"`
}

type Result struct {
	ID   SampleID  `json:"id"`
	Type PlaneType `json:"type"`
}

type ValidationError struct {
	in Input
}

func (err ValidationError) Error() string {
	return "Validation error : error !"
}

func (input *Input) Validate() error {
	var err ValidationError
	_ = err
	return nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Log the request
	err := logRequest(r)
	if err != nil {
		logger.Printf("Handler: request logging failed: %v", err)
	}

	// First let's parse the input
	log.Print("Decoding...")
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var input *Input = &Input{}
	err = dec.Decode(input)
	if err != nil {
		msg := fmt.Sprintf("JSON Decoding failed: %v", err)
		log.Printf("Handler: %v", msg)
		fmt.Fprint(w, msg)
		return
	}
	log.Print("DONE\n")
	log.Printf("%#v", *input)

	// Let's validate it
	err = input.Validate()
	if err != nil {
		msg := fmt.Sprintf("Validation error: %v", err)
		log.Printf("Handler: %s", msg)
		fmt.Fprint(w, msg)
		return
	}

	// PROCESS
	output := process(input)

	// Encode it
	log.Print("Encoding..")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	err = enc.Encode(output)
	if err != nil {
		msg := fmt.Sprintf("JSON Encoding failed: %v", err)
		log.Printf("Handler: %s", msg)
		fmt.Fprint(w, msg)
		return
	}
	log.Print("DONE\n")

	// Return
	return
}

func main() {
	http.HandleFunc("/", Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

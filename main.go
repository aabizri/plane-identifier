// TODO: Add feasability
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

var (
	portFlag = flag.Uint("p", 8080, "Port to listen to")
	pathFlag = flag.String("path", "/", "Path to endpoint folder")
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
	logger.Print("Decoding...")
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var input *Input = &Input{}
	err = dec.Decode(input)
	if err != nil {
		msg := fmt.Sprintf("JSON Decoding failed: %v", err)
		logger.Printf("Handler: %v", msg)
		fmt.Fprint(w, msg)
		return
	}
	logger.Print("DONE\n")
	logger.Printf("%#v", *input)

	// Let's validate it
	err = input.Validate()
	if err != nil {
		msg := fmt.Sprintf("Validation error: %v", err)
		logger.Printf("Handler: %s", msg)
		fmt.Fprint(w, msg)
		return
	}

	// PROCESS
	output := process(input)

	// Encode it
	logger.Print("Encoding..")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	err = enc.Encode(output)
	if err != nil {
		msg := fmt.Sprintf("JSON Encoding failed: %v", err)
		logger.Printf("Handler: %s", msg)
		fmt.Fprint(w, msg)
		return
	}
	logger.Print("DONE\n")

	// Return
	return
}

func main() {
	http.HandleFunc(*pathFlag, Handler)

	portStr := ":" + strconv.FormatUint(uint64(*portFlag), 10)
	logger.Fatal(http.ListenAndServe(portStr, nil))
}

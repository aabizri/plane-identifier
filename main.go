// TODO: Add feasability
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
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
	// Log
	log.Printf("Received request:\n%v\n", r.Body)

	// First let's parse the input
	log.Print("Decoding...")
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var input *Input = &Input{}
	err := dec.Decode(input)
	if err != nil {
		msg := fmt.Sprintf("Error: while decoding json: %v", err)
		log.Print(msg)
		fmt.Fprint(w, msg)
		return
	}
	log.Print("DONE\n")
	log.Printf("%#v", *input)

	// Let's validate it
	err = input.Validate()
	if err != nil {
		msg := fmt.Sprintf("Error: Validation error: %v", err)
		log.Print(msg)
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
		fmt.Fprintf(w, "Error: json encode: %v", err)
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

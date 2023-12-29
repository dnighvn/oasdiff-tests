package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/tufin/oasdiff/diff"
)

func useTufinOasdiff() {
	// Load base specification
	loader := openapi3.NewLoader()
	base, err := loader.LoadFromFile("spec_base.yaml")
	if err != nil {
		log.Fatalf("Error loading base spec: %s", err)
	}

	// Load revision specification
	revision, err := loader.LoadFromFile("spec_revision.yaml")
	if err != nil {
		log.Fatalf("Error loading revised spec: %s", err)
	}

	// Perform diff
	diffReport, err := diff.Get(&diff.Config{}, base, revision)
	if err != nil {
		log.Fatalf("Error performing diff: %s", err)
	}

	jsonReport, err := json.MarshalIndent(diffReport, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling diff to JSON: %s", err)
	}
	fmt.Println(string(jsonReport))
}

func main() {
	useTufinOasdiff()
}

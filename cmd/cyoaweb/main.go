package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/rafalmp/cyoa"
)

func main() {
	jsonFile := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *jsonFile)

	f, err := os.Open(*jsonFile)
	if err != nil {
		fmt.Printf("Error opening CYOA story: %s\n", err)
		os.Exit(1)
	}

	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", story)
}

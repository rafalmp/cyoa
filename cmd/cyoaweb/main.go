package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rafalmp/cyoa"
)

func main() {
	port := flag.Int("port", 8000, "the port to listen on")
	jsonFile := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *jsonFile)

	f, err := os.Open(*jsonFile)
	if err != nil {
		fmt.Printf("Error opening CYOA story: %s\n", err)
		os.Exit(1)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		os.Exit(1)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

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

	tpl := template.Must(template.New("").Parse(customTemplate))

	h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl), cyoa.WithPathFunc(pathFn))
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var customTemplate = `
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
        {{range .Options}}
        <li><a href="/story/{{.Arc}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
</html>`

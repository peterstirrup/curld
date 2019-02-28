package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Check there's an arg
	if len(os.Args) < 2 {
		log.Println("url missing - type 'curld help'")
		return
	}

	// Check arg isn't help
	if os.Args[1] == "help" {
		log.Println(`command are of the type 'curld "[url]"' where the url contains the parameters you want to test`)
		return
	}

	url := os.Args[1]

	// Split by &P
	urlParts := strings.Split(url, "&P=")
	url = urlParts[0]

	for i := 1; i < len(urlParts); i++ {
		// Add the url part to it
		url += "&P=" + urlParts[i]

		req, err := http.Get(url)
		if err != nil || req.StatusCode != 200 {
			log.Printf("\n\nbad parameter!\nparam: %s\nurl: %s\n\n", urlParts[i], url)
			// Remove it from the url
			remURL := "&P=" + urlParts[i]
			url = strings.Replace(url, remURL, "", -1)
		}
	}
}

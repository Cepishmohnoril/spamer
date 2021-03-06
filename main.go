package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"spamer/lib"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var sheetIds = []string{
	"1j-clXuNPrZlcoxNp2wxBbWINKtHBf2-dzMMQMYh5BsA",
	"1UZywZiyQBaekJXVDiXlTw5u2qA9M51Fi0q74yGw5W8I",
}

var wg sync.WaitGroup

func main() {
	for _, sheetId := range sheetIds {
		wg.Add(1)
		go func(sheetId string) {
			requests := getRequests(sheetId)
			sendRequests(requests)
			wg.Done()
		}(sheetId)
	}
	wg.Wait()
}

func getRequests(sheetId string) []string {
	fmt.Println("Parse sheet: ", sheetId)

	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")

	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := lib.GetGoogleClient(config)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	readRange := "Start!J2:DA"
	resp, err := srv.Spreadsheets.Values.Get(sheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	var requests []string

	for _, sl := range resp.Values {
		for _, req := range sl {
			req, _ := req.(string)
			requests = append(requests, req)
		}
	}

	return requests
}

func sendRequests(requests []string) {
	client := lib.GetTorClient()

	for _, url := range requests {
		lib.SendTorPost(url, client)
	}
}

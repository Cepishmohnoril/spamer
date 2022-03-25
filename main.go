package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"spamer/lib"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var sheetIds = []string{
	"1j-clXuNPrZlcoxNp2wxBbWINKtHBf2-dzMMQMYh5BsA",
	//"1UZywZiyQBaekJXVDiXlTw5u2qA9M51Fi0q74yGw5W8I",
}

func main() {

	for _, sheetId := range sheetIds {
		requests := getRequests(sheetId)
		fmt.Println(requests)
		//sendRequests()
	}
}

func getRequests(sheetId string) []string {
	fmt.Println("Parse sheet ")

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
	client := lib.GetClient(config)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	//readRange := "Start!J2:DA"

	readRange := "Start!J2:K3"
	resp, err := srv.Spreadsheets.Values.Get(sheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	var requests []string
	//fmt.Println(resp.Values)

	for _, sl := range resp.Values {
		for _, req := range sl {
			req, _ := req.(string)
			requests = append(requests, req)
		}
	}

	return requests
}

func sendRequests() {

}

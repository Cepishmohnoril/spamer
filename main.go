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
	"1UZywZiyQBaekJXVDiXlTw5u2qA9M51Fi0q74yGw5W8I",
}

func main() {

	for _, sheetId := range sheetIds {
		parseSheet(sheetId)
	}
}

func parseSheet(sheetId string) {
	fmt.Println("Parse sheet ")

	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	fmt.Println(config)

	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := lib.GetClient(config)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	readRange := "Start!J2"
	resp, err := srv.Spreadsheets.Values.Get(sheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	fmt.Println(resp)

	//if len(resp.Values) == 0 {
	//	fmt.Println("No data found.")
	//} else {
	//	fmt.Println("Name, Major:")
	//	for _, row := range resp.Values {
	//		// Print columns A and E, which correspond to indices 0 and 4.
	//		fmt.Printf("%s, %s\n", row[0], row[4])
	//	}
	//}
}

package main

import (
	"context"
	"log"
	"os"

	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func appendToSheet(spreadsheetId string, values []string) {
	ctx := context.Background()
	// obtain this from google cloud console, create a service account and download the credentials.json file, link: https://console.cloud.google.com
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := config.Client(ctx)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	var vr sheets.ValueRange
	for _, value := range values {
		date := time.Now().Format("02/01/2006")
		//these are the values that will be appended to the sheet, edit this to fit your needs, the values are (the task, the start date, the end date, the time spent, the notes, the person who did the task, the status of the task)
		vr.Values = append(vr.Values, []interface{}{value, date, date, "", "", "Rawad", "Done"})
	}
	// the sheet id is the first parameter, the sheet name is the second parameter, and the values are the third parameter
	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, "Sheet1", &vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Unable to write data to sheet: %v", err)
	}
}

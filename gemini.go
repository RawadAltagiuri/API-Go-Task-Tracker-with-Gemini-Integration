package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Content struct {
	Parts []string `json:Parts`
	Role  string   `json:Role`
}
type Candidates struct {
	Content *Content `json:Content`
}
type ContentResponse struct {
	Candidates *[]Candidates `json:Candidates`
}

func GenerateRespone(gist string) []string {

	ctx := context.Background()
	// create a new client with the api key, obtain this from generative ai, link: https://ai.google.dev, set it up as an environment variable
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GAPI_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	//you can make this prompt more detailed and clear :) this just an example
	var prompt string = fmt.Sprintf("I will now give you a gist of my daily tasks at my job and here is what I want you to do:1. I want you to separate and format different tasks. After you seperate them write them into bullet points in a format where they are clear, detailed, and each bullet points maximum word count is 100 words (do not include the word count in your response). Return only the bullet point list\n%q", gist)

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	marshalResponse, _ := json.MarshalIndent(resp, "", "  ")
	// fmt.Println(string(marshalResponse))
	var generateResponse ContentResponse
	if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
		log.Fatal(err)
	}
	var ContentResponse string = ""
	for _, cad := range *generateResponse.Candidates {
		if cad.Content != nil {
			for _, part := range cad.Content.Parts {
				ContentResponse += part
			}
		}
	}

	//split the response at the bullet points
	var bulletPoints []string = []string{}
	var bulletPoint string = ""
	for _, char := range ContentResponse {
		if char == '\n' {
			bulletPoints = append(bulletPoints, bulletPoint)
			bulletPoint = ""
		} else {
			bulletPoint += string(char)
		}
	}

	return bulletPoints
}

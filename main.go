package main

import (
	read_write "email-extractor/read"
	"email-extractor/scraper"
	"fmt"
	"regexp"
)

func processGoRoutines(urls []string) []string {
	// var html string
	chanRes := make(chan string)
	var emails []string
	for _, url := range urls {
		go scraper.GetHTML(url, chanRes)
	}

	for i := 0; i < len(urls); i++ {
		emails = append(emails, extractEmails(<-chanRes)...)
	}

	return emails
}

func extractEmails(html string) []string {
	// Define a regular expression for matching email addresses
	emailRegex := `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`

	// Compile the regular expression
	re := regexp.MustCompile(emailRegex)

	// Find all email matches in the HTML
	matches := re.FindAllString(html, -1)

	return matches
}

func main() {

	// Create an instance of the struct to hold the decoded data
	var urlsToScrap []string

	err := read_write.ReadJSON(&urlsToScrap)
	if err != nil {
		fmt.Println("an Error ocurr while reading json file : ", err)
		return
	}
	emails := processGoRoutines((urlsToScrap))

	if len(emails) == 0 {
		fmt.Println("no email found :(")
		return
	}
	err = read_write.SaveContent(emails)
	if err != nil {
		fmt.Println("an Error ocurr while reading json file : ", err)
	}
}

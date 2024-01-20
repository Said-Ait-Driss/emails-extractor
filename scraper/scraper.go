package scraper

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func GetHTML(url string, channel chan string) {
	// initialize a controllable Chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)

	// to release the browser resources when
	// it is no longer needed
	defer cancel()

	var html string

	err := chromedp.Run(ctx,
		// visit the target page
		chromedp.Navigate(url),
		// wait for the page to load
		chromedp.Sleep(2000*time.Millisecond),
		// extract the raw HTML from the page
		chromedp.OuterHTML("html", &html, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal("Error while performing the automation logic:", err)
		channel <- "<div>Error while performing the automation logic</div>"
	}
	channel <- html
}

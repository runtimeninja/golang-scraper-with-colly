package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	// Initialize collector
	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	// define callback for when an <a> element with an href attribute is found
	collector.OnHTML("a[href]", func(element *colly.HTMLElement) {
		link := element.Attr("href")
		title := element.Text
		fmt.Printf("Link: %s\nTitle: %s\n\n", link, title)
	})

	// define the callback for handling errors
	collector.OnError(func(response *colly.Response, err error) {
		log.Printf("Error occurred: %v", err)
	})

	// URL to scrape
	url := "https://www.daraz.com.bd/routers/"
	fmt.Printf("Starting to scrape: %s\n", url)

	// start scraping
	if err := collector.Visit(url); err != nil {
		log.Fatalf("Failed to scrape the website: %v", err)
	}

	fmt.Println("Scraping completed successfully.")
}

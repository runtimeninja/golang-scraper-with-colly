package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"context"
	"github.com/chromedp/chromedp"
)

func scrapeWithColly(url string) {
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

	// debugging with raw HTML
	//collector.OnResponse(func(response *colly.Response) {
	//	fmt.Printf("Raw HTML:\n%s\n", string(response.Body))
	//})
	
	// started scraping with colly
	fmt.Printf("Starting to scrape with Colly: %s\n", url)
	if err := collector.Visit(url); err != nil {
		log.Fatalf("Failed to scrape the website: %v", err)
	}

	fmt.Println("Scraping completed successfully.")
}

func scrapeWithChromedp(url string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var html string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.InnerHTML("body", &html),
	)
	if err != nil {
		log.Fatalf("Failed to scrape with chromedp: %v", err)
	}

	//fmt.Println("HTML Content from Chromedp:")
	//fmt.Println(html)
}

func main() {
	url := "https://www.daraz.com.bd/routers/"

	// scraping using Colly
	scrapeWithColly(url)

	// scraping using Chromedp
	scrapeWithChromedp(url)
}

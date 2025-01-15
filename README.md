# Web Scraper using Colly

This project demonstrates a simple web scraper built in Go using the [Colly](https://github.com/gocolly/colly) library. The scraper collects and prints all hyperlinks (`<a>` tags with `href` attributes) and their titles from the specified URL.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher is recommended)

## Installation

1. Clone the repository:
   git clone https://github.com/runtimeranger/golang-scraper-with-colly.git
   cd golang-scraper-with-colly

2. Install the Colly library:
   go get -u github.com/gocolly/colly

3. Install the chromedp library:(Its Implemented to hendle JavaScript-rendered content)
   go get -u github.com/chromedp/chromedp

## Usage

1. Modify the `url` variable in the code to the target website you want to scrape.
   url := "https://www.daraz.com.bd/routers/"

2. Run the program:
   go run main.go

3. The program will output:
   - Links (`href` attributes) found on the page
   - Titles (text content) of the links

## Code Explanation

- **Collector Initialization**:
  A new Colly collector is created with a custom `User-Agent` to mimic a browser.

- **Callbacks**:
  - `OnHTML`: Processes `<a>` tags with `href` attributes and extracts the link and title.
  - `OnError`: Handles any errors that occur during the scraping process.

- **Visit**:
  The scraper begins by visiting the specified URL.

## Example Output

Starting to scrape: https://www.daraz.com.bd/routers/
Link: /some-link
Title: Router Name

Link: /another-link
Title: Another Router

Scraping completed successfully.

## Notes

- Ensure the target website allows scraping by reviewing its `robots.txt` file and terms of service.
- Modify the scraper's callbacks as needed to extract additional data from the website.

## Acknowledgments

- [Colly](https://github.com/gocolly/colly): The lightweight and fast web scraping framework for Go.


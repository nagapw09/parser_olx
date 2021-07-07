package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://www.olx.ua/elektronika/noutbuki-i-aksesuary/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("tbody tr h3").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Find("a").Attr("href")
		titleText, _ := s.Find("strong").Html()

		fmt.Println(i, title, titleText)
	})
}

func main() {
	ExampleScrape()
}

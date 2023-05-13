package html_scrapping_service

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"net/http"
)

type HtmlScrapper struct {
	GoQueryHtmlDocument *goquery.Document
	collyHtmlCollector  *colly.Collector
}

func NewHtmlScrapper(url string) *HtmlScrapper {

	htmlScrapper := &HtmlScrapper{}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching URL: %v", err)
	}

	htmlScrapper.GetGoQueryDocument(
		resp)

	htmlScrapper.GetCollyCollector()

	return htmlScrapper

}

func (htmlScrapper *HtmlScrapper) GetGoQueryDocument(resp *http.Response) {

	defer resp.Body.Close()

	htmlDocument, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}

	htmlScrapper.GoQueryHtmlDocument = htmlDocument

}

func (htmlScrapper *HtmlScrapper) GetCollyCollector() {

	collyHtmlCollector := colly.NewCollector()

	htmlScrapper.collyHtmlCollector = collyHtmlCollector

}

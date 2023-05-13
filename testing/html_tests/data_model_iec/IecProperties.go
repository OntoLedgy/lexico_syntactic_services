package data_model_iec

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type IecProperty struct {
	Code             string
	Version          string
	Revision         string
	IRDI             string
	PreferredName    string
	SynonymousName   string
	Symbol           string
	SynonymousSymbol string
	ShortName        string
	Definition       string
	PropertyUrl      string
}

func NewIecProperty(url string) *IecProperty {

	property := &IecProperty{
		Code:             "",
		Version:          "",
		Revision:         "",
		IRDI:             "",
		PreferredName:    "",
		SynonymousName:   "",
		Symbol:           "",
		SynonymousSymbol: "",
		ShortName:        "",
		Definition:       "",
		PropertyUrl:      url,
	}

	property.scrapePropertyPage()

	return property
}

func constructPropertyURL(baseURL string, propertyID string) string {
	// Construct the URL for the property page

	propertyUrlSuffix := strings.ReplaceAll(propertyID, "/", "-")
	propertyUrlSuffix = strings.ReplaceAll(propertyUrlSuffix, "#", "%23")

	return baseURL + propertyUrlSuffix
}

func (iecProperty *IecProperty) scrapePropertyPage() {
	resp, err := http.Get(iecProperty.PropertyUrl)
	if err != nil {
		log.Fatalf("Failed to fetch property page: %v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse property page: %v", err)
	}

	doc.Find("table#contentL1 tr").Each(func(i int, tr *goquery.Selection) {
		label := strings.TrimSpace(tr.Find("td.label").Text())
		value := strings.TrimSpace(tr.Find("td").Eq(1).Text())

		switch label {
		case "Code:":
			iecProperty.Code = value
		case "Version:":
			iecProperty.Version = value
		case "Revision:":
			iecProperty.Revision = value
		case "IRDI:":
			iecProperty.IRDI = value
		case "Preferred name:":
			iecProperty.PreferredName = value
		case "Synonymous name:":
			iecProperty.SynonymousName = value
		case "Symbol:":
			iecProperty.Symbol = value
		case "Synonymous symbol:":
			iecProperty.SynonymousSymbol = value
		case "Short name:":
			iecProperty.ShortName = value
		case "Definition:":
			iecProperty.Definition = value

		}
	})
	fmt.Println(iecProperty)
}

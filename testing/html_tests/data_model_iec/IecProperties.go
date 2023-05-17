package data_model_iec

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"strings"
	"time"
)

type IecProperty struct {
	Code                     string
	Version                  string
	Revision                 string
	IRDI                     string
	PreferredName            string
	SynonymousName           string
	Symbol                   string
	SynonymousSymbol         string
	ShortName                string
	Definition               string
	Note                     string
	Remark                   string
	PrimaryUnit              string
	AlternativeUnits         string
	Level                    string
	DataType                 string //#TODO - chain to pointer to IecClasses
	Format                   string
	PropertyConstraint       string
	DefinitionSource         string
	ValueSource              string
	PropertyDataElementType  string
	Drawing                  string
	Formula                  string
	ValueListCode            string
	ValueList                string
	DETClass                 string
	ApplicableClasses        string
	DefinitionClass          string
	CodeForUnit              string
	CodesForAlternativeUnits string
	CodeForUnitList          string
	StatusLevel              string
	IsDeprecated             string
	PublishedIn              string
	PublishedBy              string
	ProposalDate             time.Time
	VersionInitiationDate    time.Time
	VersionReleaseDate       time.Time
	RevisionReleaseDate      time.Time
	ObsoleteDate             time.Time
	ResponsibleCommittee     string
	Conditions               string
	ChangeRequestID          string
	VersionHistory           string

	PropertyId  string
	PropertyUrl string
}

func constructPropertyURL(
	baseURL string,
	propertyID string) string {
	// Construct the URL for the property page

	propertyUrlSuffix := strings.ReplaceAll(propertyID, "/", "-")
	propertyUrlSuffix = strings.ReplaceAll(propertyUrlSuffix, "#", "%23")

	return baseURL + propertyUrlSuffix
}

func (iecProperty *IecProperty) scrapePropertyPage() {

	iecProperty.scrapePropertyGoQuery()

}

func (iecProperty *IecProperty) scrapePropertyGoQuery() {

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
		case "Note:":
			iecProperty.Note = value
		case "Remark:":
			iecProperty.Remark = value
		case "Primary unit:":
			iecProperty.PrimaryUnit = value
		case "Alternative units:":
			iecProperty.AlternativeUnits = value
		case "Level:":
			iecProperty.Level = value
		case "Data type:":
			iecProperty.DataType = value
		case "Format:":
			iecProperty.Format = value
		case "Property constraint:":
			iecProperty.PropertyConstraint = value
		case "Definition source:":
			iecProperty.DefinitionSource = value
		case "Value source:":
			iecProperty.ValueSource = value
		case "Property data element type:":
			iecProperty.PropertyDataElementType = value
		case "Drawing:":
			iecProperty.Drawing = value
		case "Formula:":
			iecProperty.Formula = value
		case "Value list code:":
			iecProperty.ValueListCode = value
		case "Value list:":
			iecProperty.ValueList = value
		case "DET class:":
			iecProperty.DETClass = value
		case "Applicable classes:":
			iecProperty.ApplicableClasses = value
		case "Definition class:":
			iecProperty.DefinitionClass = value
		case "Code for unit:":
			iecProperty.CodeForUnit = value
		case "Codes for alternative units:":
			iecProperty.CodesForAlternativeUnits = value
		case "Code for unit list:":
			iecProperty.CodeForUnitList = value
		case "Status level:":
			iecProperty.StatusLevel = value
		case "Is deprecated:":
			iecProperty.IsDeprecated = value
		case "Published in:":
			iecProperty.PublishedIn = value
		case "Published by:":
			iecProperty.PublishedBy = value
		case "Proposal date:":
			iecProperty.ProposalDate, _ = time.Parse("2006-01-02", value) // Adjust the date format according to your data
		case "Version initiation date:":
			iecProperty.VersionInitiationDate, _ = time.Parse("2006-01-02", value) // Adjust the date format according to your data
		case "Version release date:":
			iecProperty.VersionReleaseDate, _ = time.Parse("2006-01-02", value) // Adjust the date format according to your data
		case "Revision release date:":
			iecProperty.RevisionReleaseDate, _ = time.Parse("2006-01-02", value) // Adjust the date format according to your data
		case "Obsolete date:":
			iecProperty.ObsoleteDate, _ = time.Parse("2006-01-02", value) // Adjust the date format according to your data
		case "Responsible Committee:":
			iecProperty.ResponsibleCommittee = value
		case "Conditions:":
			iecProperty.Conditions = value
		case "Change request ID:":
			iecProperty.ChangeRequestID = value
		case "Version history:":
			iecProperty.VersionHistory = value
		}
	})
}

func (iecProperty *IecProperty) scrapePropertyColly() (*IecProperty, error) {

	// Create a new Colly collector
	c := colly.NewCollector()

	// Scrape the property page
	c.OnHTML("table#contentL1 tr", func(tr *colly.HTMLElement) {
		label := strings.TrimSpace(tr.ChildText("td.label"))
		value := strings.TrimSpace(tr.ChildText("td:nth-child(2)"))

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
		case "Note:":
			iecProperty.Note = value
		case "Remark:":
			iecProperty.Remark = value
		case "Primary unit:":
			iecProperty.PrimaryUnit = value
		case "Alternative units:":
			iecProperty.AlternativeUnits = value
		case "Level:":
			iecProperty.Level = value
		case "Data type:":
			iecProperty.DataType = value
		case "Format:":
			iecProperty.Format = value
		case "Property constraint:":
			iecProperty.PropertyConstraint = value
		case "Definition source:":
			iecProperty.DefinitionSource = value
		case "Value source:":
			iecProperty.ValueSource = value
		case "Property data element type:":
			iecProperty.PropertyDataElementType = value
		case "Drawing:":
			iecProperty.Drawing = value
		case "Formula:":
			iecProperty.Formula = value
		case "Value list code:":
			iecProperty.ValueListCode = value
		case "Value list:":
			iecProperty.ValueList = value
		case "DET class:":
			iecProperty.DETClass = value
		case "Applicable classes:":
			iecProperty.ApplicableClasses = value
		case "Definition class:":
			iecProperty.DefinitionClass = value
		case "Code for unit:":
			iecProperty.CodeForUnit = value
		case "Codes for alternative units:":
			iecProperty.CodesForAlternativeUnits = value
		case "Code for unit list:":
			iecProperty.CodeForUnitList = value
		case "Status level:":
			iecProperty.StatusLevel = value
		case "Is deprecated:":
			iecProperty.IsDeprecated = value
		case "Published in:":
			iecProperty.PublishedIn = value
		case "Published by:":
			iecProperty.PublishedBy = value
		case "Proposal date:":
			t, _ := time.Parse("2006-01-02", value)
			iecProperty.ProposalDate = t
		case "Version initiation date:":
			t, _ := time.Parse("2006-01-02", value)
			iecProperty.VersionInitiationDate = t
		case "Version release date:":
			t, _ := time.Parse("2006-01-02", value)
			iecProperty.VersionReleaseDate = t
		case "Revision release date:":
			t, _ := time.Parse("2006-01-02", value)
			iecProperty.RevisionReleaseDate = t
		case "Obsolete date:":
			t, _ := time.Parse("2006-01-02", value)
			iecProperty.ObsoleteDate = t
		case "Responsible Committee:":
			iecProperty.ResponsibleCommittee = value
		case "Conditions:":
			iecProperty.Conditions = value
		case "Change request ID:":
			iecProperty.ChangeRequestID = value
		case "Version history:":
			iecProperty.VersionHistory = value
		}
	})

	err := c.Visit(iecProperty.PropertyUrl)
	if err != nil {
		return nil, err
	}

	return iecProperty, nil
}

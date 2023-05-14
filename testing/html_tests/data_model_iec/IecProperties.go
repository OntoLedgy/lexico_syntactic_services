package data_model_iec

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"strings"
)

type IecProperty struct {
	Code                  string
	Version               string
	Revision              string
	IRDI                  string
	PreferredName         string
	SynonymousName        string
	Symbol                string
	SynonymousSymbol      string
	ShortName             string
	Definition            string
	Note                  string
	Remark                string
	DefinitionSource      string
	Drawing               string
	ClassType             string
	ApplicableDocuments   string
	ClassValueAssignment  string
	RequisityOfProperties string
	Superclass            string
	HigherLevelClasses    string
	ClassifyingDET        string
	Properties            string
	PropertiesTree        string
	InheritedProperties   string
	SuperBlocks           string
	IsCaseOf              string
	ImportedProperties    string
	InstanceSharable      string
	StatusLevel           string
	IsDeprecated          string
	PublishedIn           string
	PublishedBy           string
	ProposalDate          string
	VersionInitiationDate string
	VersionReleaseDate    string
	RevisionReleaseDate   string
	ObsoleteDate          string
	ResponsibleCommittee  string
	ChangeRequestID       string
	VersionHistory        string

	PropertyId  string
	PropertyUrl string
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
		case "Synonymous Symbol:":
			iecProperty.SynonymousSymbol = value
		case "Short name:":
			iecProperty.ShortName = value
		case "Definition:":
			iecProperty.Definition = value
		case "Note:":
			iecProperty.Note = value
		case "Remark:":
			iecProperty.Remark = value
		case "Definition source:":
			iecProperty.DefinitionSource = value
		case "Drawing:":
			iecProperty.Drawing = value
		case "Class type:":
			iecProperty.ClassType = value
		case "Applicable documents:":
			iecProperty.ApplicableDocuments = value
		case "Class value assignment:":
			iecProperty.ClassValueAssignment = value
		case "Requisity of properties:":
			iecProperty.RequisityOfProperties = value
		case "Superclass:":
			iecProperty.Superclass = value
		case "Higher level classes:":
			iecProperty.HigherLevelClasses = value
		case "Classifying DET:":
			iecProperty.ClassifyingDET = value
		case "Properties:":
			iecProperty.Properties = value
		case "Properties tree:":
			iecProperty.PropertiesTree = value
		case "Inherited properties:":
			iecProperty.InheritedProperties = value
		case "SuperBlocks:":
			iecProperty.SuperBlocks = value
		case "Is case of:":
			iecProperty.IsCaseOf = value
		case "Imported properties:":
			iecProperty.ImportedProperties = value
		case "Instance sharable:":
			iecProperty.InstanceSharable = value
		case "Status level:":
			iecProperty.StatusLevel = value
		case "Is deprecated:":
			iecProperty.IsDeprecated = value
		case "Published in:":
			iecProperty.PublishedIn = value
		case "Published by:":
			iecProperty.PublishedBy = value
		case "Proposal date:":
			iecProperty.ProposalDate = value
		case "Version initiation date:":
			iecProperty.VersionInitiationDate = value
		case "Version release date:":
			iecProperty.VersionReleaseDate = value
		case "Revision release date:":
			iecProperty.RevisionReleaseDate = value
		case "Obsolete date:":
			iecProperty.ObsoleteDate = value
		case "Responsible Committee:":
			iecProperty.ResponsibleCommittee = value
		case "Change request ID:":
			iecProperty.ChangeRequestID = value
		case "Version history:":
			iecProperty.VersionHistory = value

		}
	})

}

func (iecProperty *IecProperty) scrapeProperty(url string) (*IecProperty, error) {
	property := &IecProperty{PropertyUrl: url}

	// Create a new Colly collector
	c := colly.NewCollector()

	// Scrape the property page
	c.OnHTML("table#contentL1 tr", func(tr *colly.HTMLElement) {
		label := strings.TrimSpace(tr.ChildText("td.label"))
		value := strings.TrimSpace(tr.ChildText("td:nth-child(2)"))

		switch label {
		case "Code:":
			property.Code = value
		case "Version:":
			property.Version = value
		case "Revision:":
			property.Revision = value
		case "IRDI:":
			property.IRDI = value
		case "Preferred name:":
			property.PreferredName = value
		case "Synonymous name:":
			property.SynonymousName = value
		case "Symbol:":
			property.Symbol = value
		case "Synonymous symbol:":
			property.SynonymousSymbol = value
		case "Short name:":
			property.ShortName = value
		case "Definition:":
			property.Definition = value
		case "Note:":
			property.Note = value
		case "Remark:":
			property.Remark = value
		case "Definition source:":
			property.DefinitionSource = value
		case "Drawing:":
			property.Drawing = value
		case "Class type:":
			property.ClassType = value
		case "Applicable documents:":
			property.ApplicableDocuments = value
		case "Class value assignment:":
			property.ClassValueAssignment = value
		case "Requisity of properties:":
			property.RequisityOfProperties = value
		case "Superclass:":
			property.Superclass = value
		case "Higher level classes:":
			property.HigherLevelClasses = value
		case "Classifying DET:":
			property.ClassifyingDET = value
		case "Properties:":
			property.Properties = value
		case "Properties tree:":
			property.PropertiesTree = value
		case "Inherited properties:":
			property.InheritedProperties = value
		case "SuperBlocks:":
			property.SuperBlocks = value
		case "Is case of:":
			property.IsCaseOf = value
		case "Imported properties:":
			property.ImportedProperties = value
		case "Instance sharable:":
			property.InstanceSharable = value
		case "Status level:":
			property.StatusLevel = value
		case "Is deprecated:":
			property.IsDeprecated = value
		case "Published in:":
			property.PublishedIn = value
		case "Published by:":
			property.PublishedBy = value
		case "Proposal date:":
			property.ProposalDate = value
		case "Version initiation date:":
			property.VersionInitiationDate = value
		case "Version release date:":
			property.VersionReleaseDate = value
		case "Revision release date:":
			property.RevisionReleaseDate = value
		case "Obsolete date:":
			property.ObsoleteDate = value
		case "Responsible Committee:":
			property.ResponsibleCommittee = value
		case "Change request ID:":
			property.ChangeRequestID = value
		case "Version history:":
			property.VersionHistory = value
		}
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return property, nil
}

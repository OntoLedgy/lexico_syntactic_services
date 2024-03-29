package data_model_iec

import (
	"fmt"
	"github.com/OntoLedgy/syntactic_checker/testing/html_tests/html_scrapping_service"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

type IecClasses struct {
	Code                  string
	Version               string
	Revision              string
	IRDI                  string
	PreferredName         string
	SynonymousName        string
	CodedName             string
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
	Superclass            *IecClasses
	ClassifyingDET        string
	Properties            []*IecProperty

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

	SuperClassCode        string
	ClassUrl              string
	PropertyLinks         []string
	SuperClassPropertyIds []string
}

func (iecClass *IecClasses) GetClassURL(classIRDI string) string {

	// If not, load the IecClass by scraping the URL
	classIDUrlFriendly, err := iecClass.validateAndConvertClassID(
		classIRDI)

	if err != nil {
		fmt.Printf("invalid class id %s, %s", classIRDI, err)
		return ""
	}

	if iecClass.ClassUrl == "" {
		// Construct the URL using the base URL and class ID
		baseURL := "https://cdd.iec.ch/CDD/IEC61987/iec61987.nsf/Classes/"
		iecClass.ClassUrl = baseURL + classIDUrlFriendly
	}

	return iecClass.ClassUrl
}

func (iecClass *IecClasses) listPropertyLinks() {

	// Prepare a slice to store property URLs.
	var propertyURLs []string
	urlPrefix := "https://cdd.iec.ch/cdd/iec61987/iec61987.nsf/PropertiesAllVersions/"

	// Iterate through each line and extract the property ID.
	for _, iecPropertyLink := range iecClass.PropertyLinks {

		iecPropertyLink = strings.TrimSpace(iecPropertyLink)

		if iecPropertyLink == "" {
			continue
		}

		propertyID := strings.Split(iecPropertyLink, " - ")[0]
		propertyURL := constructPropertyURL(urlPrefix, propertyID)

		propertyURLs = append(propertyURLs, propertyURL)
	}

	iecClass.PropertyLinks = propertyURLs

}

func (iecClass *IecClasses) scrapeClassPage() *IecClasses {

	htmlScrapper := html_scrapping_service.NewHtmlScrapper(iecClass.ClassUrl)

	var properties []string

	htmlScrapper.GoQueryHtmlDocument.Find("table#contentL1 tr").Each(func(i int, tr *goquery.Selection) {
		label := strings.TrimSpace(tr.Find("td.label").Text())
		value := strings.TrimSpace(tr.Find("td").Eq(1).Text())

		switch label {

		case "Code:":
			iecClass.Code = strings.TrimSpace(value)
		case "Version:":
			iecClass.Version = strings.TrimSpace(value)
		case "Revision:":
			iecClass.Revision = strings.TrimSpace(value)
		case "IRDI:":
			iecClass.IRDI = strings.TrimSpace(value)
		case "Preferred name:":
			iecClass.PreferredName = strings.TrimSpace(value)
		case "Synonymous name:":
			iecClass.SynonymousName = strings.TrimSpace(value)
		case "Coded name:":
			iecClass.CodedName = strings.TrimSpace(value)
		case "Short name:":
			iecClass.ShortName = strings.TrimSpace(value)
		case "Definition:":
			iecClass.Definition = strings.TrimSpace(value)
		case "Note:":
			iecClass.Note = strings.TrimSpace(value)
		case "Remark:":
			iecClass.Remark = strings.TrimSpace(value)
		case "Definition source:":
			iecClass.DefinitionSource = strings.TrimSpace(value)
		case "Drawing:":
			iecClass.Drawing = strings.TrimSpace(value)
		case "Class type:":
			iecClass.ClassType = strings.TrimSpace(value)
		case "Applicable documents:":
			iecClass.ApplicableDocuments = strings.TrimSpace(value)
		case "Class value assignment:":
			iecClass.ClassValueAssignment = strings.TrimSpace(value)
		case "Requisity of properties:":
			iecClass.RequisityOfProperties = strings.TrimSpace(value)
		case "Superclass:":
			iecClass.SuperClassCode = strings.TrimSpace(value)

		case "Classifying DET:":
			iecClass.ClassifyingDET = strings.TrimSpace(value)

		case "Inherited properties:":
			iecClass.InheritedProperties = strings.TrimSpace(value)
		case "SuperBlocks:":
			iecClass.SuperBlocks = strings.TrimSpace(value)
		case "Is case of:":
			iecClass.IsCaseOf = strings.TrimSpace(value)
		case "Imported properties:":
			iecClass.ImportedProperties = strings.TrimSpace(value)
		case "Instance sharable:":
			iecClass.InstanceSharable = strings.TrimSpace(value)
		case "Status level:":
			iecClass.StatusLevel = strings.TrimSpace(value)
		case "Is deprecated:":
			iecClass.IsDeprecated = strings.TrimSpace(value)
		case "Published in:":
			iecClass.PublishedIn = strings.TrimSpace(value)
		case "Published by:":
			iecClass.PublishedBy = strings.TrimSpace(value)
		case "Proposal date:":
			iecClass.ProposalDate = strings.TrimSpace(value)
		case "Version initiation date:":
			iecClass.VersionInitiationDate = strings.TrimSpace(value)
		case "Version release date:":
			iecClass.VersionReleaseDate = strings.TrimSpace(value)
		case "Revision release date:":
			iecClass.RevisionReleaseDate = strings.TrimSpace(value)
		case "Obsolete date:":
			iecClass.ObsoleteDate = strings.TrimSpace(value)
		case "Responsible Committee:":
			iecClass.ResponsibleCommittee = strings.TrimSpace(value)
		case "Change request ID:":
			iecClass.ChangeRequestID = strings.TrimSpace(value)
		case "Version history:":
			iecClass.VersionHistory = strings.TrimSpace(value)

		case "Properties:":

			tr.Find("a").Each(func(i int, s *goquery.Selection) {
				properties = append(properties, s.Text())
			})

		}
	})

	htmlScrapper.GoQueryHtmlDocument.Find("div#PropertiesContentExpanded_1 a").Each(func(i int, s *goquery.Selection) {
		properties = append(properties, s.Text())
	})
	iecClass.PropertyLinks = properties

	return iecClass
}

func (iecClass *IecClasses) inheritPropertiesFromSuperClasses() {

	if iecClass.Superclass != nil {
		// Inherit properties from the superclass
		iecClass.Superclass.inheritPropertiesFromSuperClasses()

		// Create a map to store unique properties
		uniqueProperties := make(map[string]*IecProperty)

		// Add properties from the current class
		for _, property := range iecClass.Properties {
			uniqueProperties[property.Code] = property
		}

		// Add inherited properties from the superclass
		for _, property := range iecClass.Superclass.Properties {
			if _, ok := uniqueProperties[property.Code]; !ok {
				uniqueProperties[property.Code] = property
			}
		}

		// Update the Properties slice with the unique properties
		iecClass.SuperClassPropertyIds = make([]string, 0, len(uniqueProperties))

		for _, property := range uniqueProperties {
			//fmt.Print("found inherited properties")
			iecClass.SuperClassPropertyIds = append(
				iecClass.SuperClassPropertyIds,
				property.PropertyId)
		}
	}
}

func (iecClass *IecClasses) validateAndConvertClassID(
	classID string) (string, error) {
	pattern := `^(\d{4})(/|-)(\d{1})(/|-)(/|-)(/|-)(\d{5})(%23|#)([0-9A-Z]{6})(#([0-9]{3}))?$`
	re := regexp.MustCompile(pattern)

	if !re.MatchString(classID) {
		return "", fmt.Errorf("invalid classID format")
	}

	// Trim the last component of the ID that has a hash
	parts := strings.Split(classID, "#")
	if len(parts) > 2 {
		classID = strings.Join(parts[:2], "#")
	}

	// Replace / with -
	classID = strings.ReplaceAll(classID, "/", "-")

	// Replace # with %23
	classID = strings.ReplaceAll(classID, "#", "%23")

	return classID, nil

}

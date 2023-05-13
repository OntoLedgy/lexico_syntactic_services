package html_tests

import (
	"fmt"
	"github.com/OntoLedgy/syntactic_checker/testing/html_tests/data_model_iec"
	"github.com/OntoLedgy/syntactic_checker/testing/html_tests/storage_interop_services"
	"testing"
)

func TestIECScrapper(t *testing.T) {

	// Read the excel file containing the list of class IDs.
	classIDs := storage_interop_services.ReadExcelFile("E:\\ontologies\\iec\\classes_100.xlsx")

	iecCLassFactory := &data_model_iec.IecClassesFactory{}

	for _, classID := range classIDs {

		// 2. Scrape the web page content and extract the table of class information.
		iecCLassFactory.GetIecClass(classID)

		//iecClass := data_model_iec.NewIecClass(classID)

		//if iecClass != nil {
		//	classes = append(classes, iecClass)
		//}

	}

	// 5. Write the IecClasses and IecProperty data structures to an excel file.
	iecCLassFactory.ReportIecModel()
}

func TestScrapeClass(t *testing.T) {

	// Example usage of scrapePropertyPage.
	classID := "0112-2---61987%23ABE309"

	iecClassFactory := data_model_iec.NewIecClassesFactory()

	iecClass := iecClassFactory.NewIecClass(classID)

	fmt.Printf("%+v\n", iecClass)

	iecClassFactory.ReportIecModel()
}

func TestScrapeProperty(t *testing.T) {

	// Example usage of scrapePropertyPage.
	url := "https://cdd.iec.ch/cdd/iec61987/iec61987.nsf/PropertiesAllVersions/0112-2---61987%23ABU125"

	property := data_model_iec.NewIecProperty(url)
	fmt.Printf("%+v\n", property)
}

package html_tests

import (
	"fmt"
	"github.com/OntoLedgy/syntactic_checker/testing/html_tests/data_model_iec"
	"github.com/OntoLedgy/syntactic_checker/testing/html_tests/storage_interop_services"
	"strconv"
	"testing"
)

func TestIECScrapper(t *testing.T) {

	// Read the excel file containing the list of class IDs.
	classIDs := storage_interop_services.ReadExcelFile("E:\\ontologies\\iec\\classes.xlsx")

	iecCLassFactory := data_model_iec.NewIecClassesFactory()

	for rowIndex, classIRDI := range classIDs {

		// 2. Scrape the web page content and extract the table of class information.
		iecCLassFactory.GetIecClass(
			classIRDI)

		if (rowIndex+1)%100 == 0 {
			iecCLassFactory.ReportIecModel("E:\\ontologies\\iec\\output_" + strconv.Itoa(rowIndex) + ".xlsx")
		}

	}

	// 5. Write the IecClasses and IecProperty data structures to an excel file.
	iecCLassFactory.ReportIecModel("D:\\S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\data\\outputs\\html_scrapper\\iec_report_full.xlsx")
}

func TestScrapeClass(t *testing.T) {

	// Example usage of scrapePropertyPage.
	classID := "0112-2---61987%23ABE309"

	iecClassFactory := data_model_iec.NewIecClassesFactory()

	iecClass := iecClassFactory.NewIecClass(classID)

	fmt.Printf("%+v\n", iecClass)

	iecClassFactory.ReportIecModel("E:\\ontologies\\iec\\output.xlsx")
}

func TestScrapeProperty(t *testing.T) {

	// Example usage of scrapePropertyPage.
	url := "https://cdd.iec.ch/cdd/iec61987/iec61987.nsf/PropertiesAllVersions/0112-2---61987%23ABU125"
	propertyFactory := data_model_iec.NewIecPropertiesFactory()

	property, _ := propertyFactory.GetIecProperty(url)
	fmt.Printf("%+v\n", property)
}

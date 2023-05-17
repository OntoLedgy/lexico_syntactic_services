package data_model_iec

import (
	"fmt"
	"github.com/OntoLedgy/syntactic_checker/testing/html_tests/storage_interop_services"
	"reflect"
	"sync"
)

type IecClassesFactory struct {
	ClassRegister      map[string]*IecClasses
	PropertiesRegister map[string]*IecProperty
	mu                 sync.RWMutex
}

func NewIecClassesFactory() *IecClassesFactory {
	return &IecClassesFactory{
		ClassRegister:      make(map[string]*IecClasses),
		PropertiesRegister: make(map[string]*IecProperty),
	}
}

func (iecClassFactory *IecClassesFactory) GetIecClass(
	classID string) (*IecClasses, error) {

	iecClassFactory.mu.RLock()
	defer iecClassFactory.mu.RUnlock()

	return iecClassFactory.getIecClass(classID)
}

func (iecClassFactory *IecClassesFactory) getIecClass(
	classIRDI string) (*IecClasses, error) {
	// Check if the IecClass is already loaded in memory

	if iecClass, ok := iecClassFactory.ClassRegister[classIRDI]; ok {
		return iecClass, nil
	}

	iecClass :=
		iecClassFactory.NewIecClass(
			classIRDI)

	iecClassFactory.ClassRegister[classIRDI] = iecClass

	if iecClass.Properties != nil {
		for _, property := range iecClass.Properties {
			iecClassFactory.PropertiesRegister[property.IRDI] = property
		}

	}

	return iecClass, nil
}

func (iecClassFactory *IecClassesFactory) GetIecClassByCode(IecClassCode string) (*IecClasses, error) {
	iecClassFactory.mu.RLock()
	defer iecClassFactory.mu.RUnlock()

	class, ok := iecClassFactory.ClassRegister[IecClassCode]
	if !ok {
		return nil, fmt.Errorf("class with code %s not found", IecClassCode)
	}

	return class, nil
}

func (iecClassFactory *IecClassesFactory) NewIecClass(
	classID string) *IecClasses {

	iecClass := &IecClasses{}

	// 1. Construct the URL using the base URL and class ID.
	url := iecClass.GetClassURL(classID)

	// Factory method for creating an IecClasses instance
	fmt.Printf("reading class %s\n", url)

	iecClass.scrapeClassPage()

	// Process the superclass
	if iecClass.SuperClassCode != "" {

		superclass, err :=
			iecClassFactory.GetIecClass(iecClass.SuperClassCode)

		if err != nil {
			return nil
		}
		iecClass.Superclass = superclass
	}

	// 3. Identify the list of links to property pages.
	iecClass.listPropertyLinks()

	fmt.Printf("properties found at %s:\n %s", url, iecClass.PropertyLinks)

	// 4. Scrape each property page and extract the property information.

	for _, propertyLink := range iecClass.PropertyLinks {
		// 4.a. Construct the URL for the property page.
		//propertyURL := constructPropertyURL(propertyURLPrefix, propertyLink)

		// 4.b. Scrape the property web page content and extract the property information.
		iecPropertyFactory := NewIecPropertiesFactory()
		iecProperty, _ := iecPropertyFactory.GetIecProperty(propertyLink)

		if iecProperty != nil {
			iecClass.Properties = append(iecClass.Properties, iecProperty)
			iecClassFactory.PropertiesRegister[iecProperty.PropertyId] = iecProperty
		}
	}

	if iecClass != nil {
		iecClassFactory.ClassRegister[iecClass.Code] = iecClass
	}

	iecClass.inheritPropertiesFromSuperClasses()

	return iecClass
}

func (iecClassFactory *IecClassesFactory) ReportIecModel(
	fileNameAndPath string) {

	classKeys := reflect.ValueOf(iecClassFactory.ClassRegister).MapKeys()

	classTableData := convertStructToTable(iecClassFactory, classKeys)

	storage_interop_services.WriteTableDataToSheet("Classes", *classTableData, fileNameAndPath)

	propertyKeys := reflect.ValueOf(iecClassFactory.PropertiesRegister).MapKeys()

	propertyTableData := &storage_interop_services.TableData{}

	propertyTableData.Headers, _ = storage_interop_services.GetStructFieldNames(iecClassFactory.PropertiesRegister[propertyKeys[0].String()])

	for _, property := range iecClassFactory.PropertiesRegister {
		flattenedProperty, _ := storage_interop_services.FlattenAttributes(property)
		propertyTableData.Rows = append(propertyTableData.Rows, flattenedProperty)
	}

	storage_interop_services.WriteTableDataToSheet("Properties", *propertyTableData, fileNameAndPath)

}

func convertStructToTable(
	iecClassFactory *IecClassesFactory,
	keys []reflect.Value) *storage_interop_services.TableData {

	tableData := &storage_interop_services.TableData{}

	tableData.Headers, _ = storage_interop_services.GetStructFieldNames(iecClassFactory.ClassRegister[keys[0].String()])

	for _, class := range iecClassFactory.ClassRegister {
		flattenedClass, err := storage_interop_services.FlattenAttributes(class)
		if err != nil {
			fmt.Println(err)
		}
		tableData.Rows = append(tableData.Rows, flattenedClass)
	}
	return tableData
}

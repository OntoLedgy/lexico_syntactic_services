package data_model_iec

import (
	"fmt"
	"github.com/OntoLedgy/syntactic_checker/testing/html_tests/storage_interop_services"
	"reflect"
	"strconv"
	"sync"
)

type IecClassesFactory struct {
	ClassRegister      map[string]*IecClasses
	PropertiesRegister map[string]*IecProperty
	mu                 sync.Mutex
}

func NewIecClassesFactory() *IecClassesFactory {
	return &IecClassesFactory{
		ClassRegister:      make(map[string]*IecClasses),
		PropertiesRegister: make(map[string]*IecProperty),
	}
}

func (iecClassFactory *IecClassesFactory) GetIecClass(classID string) (*IecClasses, error) {
	iecClassFactory.mu.Lock()
	defer iecClassFactory.mu.Unlock()

	// Check if the IecClass is already loaded in memory
	if iecClass, ok := iecClassFactory.ClassRegister[classID]; ok {
		return iecClass, nil
	}

	// If not, load the IecClass by scraping the URL
	iecClass := iecClassFactory.NewIecClass(classID)

	iecClassFactory.ClassRegister[classID] = iecClass

	if iecClass.Properties != nil {
		for _, property := range iecClass.Properties {
			iecClassFactory.PropertiesRegister[property.IRDI] = property
		}

	}

	return iecClass, nil
}

func (iecClassFactory *IecClassesFactory) NewIecClass(classID string) *IecClasses {

	// 1. Construct the URL using the base URL and class ID.

	baseURL := "https://cdd.iec.ch/CDD/IEC61987/iec61987.nsf/Classes/"

	url := constructClassURL(baseURL, classID)
	fmt.Printf("reading %s\n", url)

	iecClass := &IecClasses{
		ClassUrl: url}

	// Factory method for creating an IecClasses instance
	iecClass.scrapeClassPage()

	// 3. Identify the list of links to property pages.
	iecClass.identifyPropertyLinks()

	fmt.Printf("properties found at %s:\n %s", url, iecClass.PropertyLinks)

	// 4. Scrape each property page and extract the property information.

	for _, propertyLink := range iecClass.PropertyLinks {
		// 4.a. Construct the URL for the property page.
		//propertyURL := constructPropertyURL(propertyURLPrefix, propertyLink)

		// 4.b. Scrape the property web page content and extract the property information.
		iecProperty := NewIecProperty(propertyLink)

		if iecProperty != nil {
			iecClass.Properties = append(iecClass.Properties, iecProperty)
		}
	}

	if iecClass != nil {
		iecClassFactory.ClassRegister[iecClass.IRDI] = iecClass
	}

	return iecClass
}

func (iecClassFactory *IecClassesFactory) ReportIecModel() {

	keys := reflect.ValueOf(iecClassFactory.ClassRegister).MapKeys()

	classTableData := &storage_interop_services.TableData{}

	classTableData.Headers, _ = getStructFieldNames(iecClassFactory.ClassRegister[keys[0].String()])

	for _, class := range iecClassFactory.ClassRegister {
		flattenedClass, err := flattenAttributes(class)
		if err != nil {
			fmt.Println(err)
		}
		classTableData.Rows = append(classTableData.Rows, flattenedClass)
	}

	storage_interop_services.WriteTableDataToSheet("Classes", *classTableData, "E:\\ontologies\\iec\\output.xlsx")

	propertyKeys := reflect.ValueOf(iecClassFactory.ClassRegister).MapKeys()

	propertyTableData := &storage_interop_services.TableData{}

	propertyTableData.Headers, _ = getStructFieldNames(iecClassFactory.ClassRegister[propertyKeys[0].String()])

	for _, property := range iecClassFactory.PropertiesRegister {
		flattenedClass, _ := flattenAttributes(property)
		propertyTableData.Rows = append(propertyTableData.Rows, flattenedClass)
	}

	storage_interop_services.WriteTableDataToSheet("Classes", *propertyTableData, "output.xlsx")

}

func getStructFieldNames(obj interface{}) ([]string, error) {

	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a struct or pointer to struct")
	}

	t := v.Type()

	fieldNames := make([]string, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		fieldNames[i] = t.Field(i).Name
	}

	return fieldNames, nil
}

func flattenAttributes(obj interface{}) ([]string, error) {
	var flattened []string
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		var fieldStr string
		switch field.Kind() {
		case reflect.String:
			fieldStr = field.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldStr = strconv.FormatInt(field.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fieldStr = strconv.FormatUint(field.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			fieldStr = strconv.FormatFloat(field.Float(), 'f', -1, 64)
		case reflect.Bool:
			fieldStr = strconv.FormatBool(field.Bool())
		case reflect.Struct:
			nested, err := flattenAttributes(field.Interface())
			if err != nil {
				return nil, err
			}
			flattened = append(flattened, nested...)
			continue

		case reflect.Ptr:
			if field.IsNil() {
				fieldStr = ""
			} else {
				nested, err := flattenAttributes(field.Interface())
				if err != nil {
					return nil, err
				}
				flattened = append(flattened, nested...)
				continue
			}

		default:
			return nil, fmt.Errorf("unsupported field type: %v", field.Type())
		}

		flattened = append(flattened, fieldStr)
	}

	return flattened, nil
}

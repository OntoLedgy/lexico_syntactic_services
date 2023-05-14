package data_model_iec

import (
	"fmt"
	"github.com/OntoLedgy/syntactic_checker/testing/html_tests/storage_interop_services"
	"reflect"
	"strconv"
	"strings"
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
	classID string) (*IecClasses, error) {
	// Check if the IecClass is already loaded in memory

	if iecClass, ok := iecClassFactory.ClassRegister[classID]; ok {
		return iecClass, nil
	}

	// If not, load the IecClass by scraping the URL
	iecClass :=
		iecClassFactory.NewIecClass(
			classID)

	iecClassFactory.ClassRegister[classID] = iecClass

	if iecClass.Properties != nil {
		for _, property := range iecClass.Properties {
			iecClassFactory.PropertiesRegister[property.IRDI] = property
		}

	}

	return iecClass, nil
}

func (iecClassFactory *IecClassesFactory) NewIecClass(
	classID string) *IecClasses {

	// 1. Construct the URL using the base URL and class ID.

	url := constructClassURL(classID)
	fmt.Printf("reading %s\n", url)

	iecClass := &IecClasses{
		ClassUrl: url}

	// Factory method for creating an IecClasses instance
	iecClass.scrapeClassPage()

	// Process the superclass
	if iecClass.SuperClassUrl != "" {
		superclass, err := iecClassFactory.GetIecClass(iecClass.SuperClassUrl)
		if err != nil {
			return nil
		}
		iecClass.Superclass = superclass
	}

	// 3. Identify the list of links to property pages.
	iecClass.identifyPropertyLinks()

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
		iecClassFactory.ClassRegister[iecClass.IRDI] = iecClass
	}

	return iecClass
}

func (iecClassFactory *IecClassesFactory) ReportIecModel(
	fileNameAndPath string) {

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

	storage_interop_services.WriteTableDataToSheet("Classes", *classTableData, fileNameAndPath)

	propertyKeys := reflect.ValueOf(iecClassFactory.PropertiesRegister).MapKeys()

	propertyTableData := &storage_interop_services.TableData{}

	propertyTableData.Headers, _ = getStructFieldNames(iecClassFactory.PropertiesRegister[propertyKeys[0].String()])

	for _, property := range iecClassFactory.PropertiesRegister {
		flattenedProperty, _ := flattenAttributes(property)
		propertyTableData.Rows = append(propertyTableData.Rows, flattenedProperty)
	}

	storage_interop_services.WriteTableDataToSheet("Properties", *propertyTableData, fileNameAndPath)

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

	var fieldNames []string

	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i).Type

		// Skip struct fields
		if fieldType.Kind() == reflect.Struct || (fieldType.Kind() == reflect.Slice && (fieldType.Elem().Kind() == reflect.Struct || (fieldType.Elem().Kind() == reflect.Ptr && fieldType.Elem().Elem().Kind() == reflect.Struct))) {
			continue
		}

		fieldNames = append(fieldNames, t.Field(i).Name)
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

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i).Type

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
			continue
		case reflect.Ptr:
			if fieldType.Elem().Kind() == reflect.Struct {
				continue
			}
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
		case reflect.Slice:
			elem := field.Type().Elem()
			if elem.Kind() == reflect.String {
				stringComponent := make([]string, field.Len())
				for i := 0; i < field.Len(); i++ {
					stringComponent[i] = field.Index(i).String()
				}
				fieldStr = strings.Join(stringComponent, ",")
			} else {
				continue
			}
		default:
			return nil, fmt.Errorf("unsupported field type: %v", field.Type())
		}

		flattened = append(flattened, fieldStr)
	}

	return flattened, nil
}

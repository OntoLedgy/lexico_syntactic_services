package storage_interop_services

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func GetStructFieldNames(
	obj interface{}) (
	[]string,
	error) {
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

func FlattenAttributes(
	obj interface{}) (
	[]string,
	error) {
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
				nested, err := FlattenAttributes(field.Interface())
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

package core

import (
	"fmt"
	"reflect"
	"strings"
)

//GetFields Devuelvo los atributos del objeto
func GetFields(v interface{}) []string {
	// v = reflect.ValueOf(v).Interface().(interface{})
	val := reflect.ValueOf(v).Elem()
	var fields []string

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		fields = append(fields, typeField.Name)
	}
	return fields
}

//GetValues Devuelvo los values del objeto
func GetValues(v interface{}) []string {
	val := reflect.ValueOf(v).Elem()
	var fields []string
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		field := ""
		if reflect.TypeOf(valueField.Interface()).Kind() == reflect.Slice {
			slice, ok := valueField.Interface().([]string)
			if !ok {
				panic("value not a []string")
			}
			field = strings.Join(slice, ",")
		} else {
			field = fmt.Sprintf("%v", valueField.Interface())
		}
		fields = append(fields, field)
	}
	return fields
}

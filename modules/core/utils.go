package core

import (
	"fmt"
	"reflect"
	"strings"
)

// var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

// fmt.Println(validID.MatchString("adam[23]"))
// fmt.Println(validID.MatchString("eve[7]"))
// fmt.Println(validID.MatchString("Job[48]"))
// fmt.Println(validID.MatchString("snakey"))
//GetPublicFields Devuelvo los atributos del objeto
func GetPublicFields(v interface{}) []string {
	// v = reflect.ValueOf(v).Interface().(interface{})
	val := reflect.ValueOf(v).Elem()
	var fields []string

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		var first = fmt.Sprintf("%v", string(typeField.Name[0]))

		if strings.ToUpper(first) == first {
			fields = append(fields, typeField.Name)
		}
	}
	return fields
}

//PASAR LISTA PARA TOMAR FIELDS
//GetPublicValues Devuelvo los values del objeto
func GetPublicValues(v interface{}) []string {
	val := reflect.ValueOf(v).Elem()
	var fields []string
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		var first = fmt.Sprintf("%v", string(typeField.Name[0]))
		if strings.ToUpper(first) == first {
			valueField := val.Field(i)
			field := ""
			if valueField.CanInterface() {
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
		}

	}
	return fields
}

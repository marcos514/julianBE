package core

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

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

func GetPublicFieldsReflect(v interface{}) []reflect.Value {
	// v = reflect.ValueOf(v).Interface().(interface{})
	fmt.Printf("First val %v\n", v)
	val := reflect.ValueOf(v).Elem()
	var fields []reflect.Value

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		var first = fmt.Sprintf("%v", string(typeField.Name[0]))

		if strings.ToUpper(first) == first {

			fields = append(fields, val.Field(i))
		}
	}
	return fields
}

//PASAR LISTA PARA TOMAR FIELDS
//GetPublicValues Devuelvo los values del objeto
func GetPublicValues(v interface{}) []string {
	val := reflect.ValueOf(v)
	listFields := GetPublicFields(v)
	var fieldsValues []string
	for i := 0; i < len(listFields); i++ {
		field := listFields[i]
		f := reflect.Indirect(val).FieldByName(field)
		fieldValue := f.Interface()
		switch v := fieldValue.(type) {
		case bool:
			b := "false"
			if v {
				b = "true"
			}
			fieldsValues = append(fieldsValues, b)
		case time.Time:
			fmt.Printf("THIS IS THE TIME %v THIS IS THE STRING %v ANSIC %v\n\n\n\n\n\n", v, v.String(), v.Format(time.ANSIC))
			fieldsValues = append(fieldsValues, v.Format(time.ANSIC))
		default:
			fieldsValues = append(fieldsValues, fmt.Sprintf("%v", v))
		}
		// if strings.ToUpper(first) == first {
		// 	valueField := val.Field(i)
		// 	field := ""
		// 	if valueField.CanInterface() {
		// 		// valueType := fmt.Printf("%v", reflect.TypeOf(valueField.Interface()).Kind())
		// 		// switch valueType {
		// 		// case "reflect.Slice":
		// 		// 	slice, ok := valueField.Interface().([]string)
		// 		// 	if !ok {
		// 		// 		panic("value not a []string")
		// 		// 	}
		// 		// 	field = strings.Join(slice, ",")
		// 		// case "time.Time":
		// 		// 	field = time.Parse(time.ANSIC, valueField.Interface())

		// 		// default:
		// 		// 	field = fmt.Sprintf("%v", valueField.Interface())
		// 		// }
		// 		d := valueField.Interface()
		// 		fmt.Printf("Like valueField kind %v and the interface is %v\n\n\n", reflect.TypeOf(d).Kind(), d.(type))
		// 		if reflect.TypeOf(valueField.Interface()).Kind() == reflect.Slice {
		// 			slice, ok := valueField.Interface().([]string)
		// 			if !ok {
		// 				panic("value not a []string")
		// 			}
		// 			field = strings.Join(slice, ",")
		// 		} else {
		// 			field = fmt.Sprintf("%v", valueField.Interface())
		// 		}
		// 		fields = append(fields, field)
		// 	}
	}
	return fieldsValues
}

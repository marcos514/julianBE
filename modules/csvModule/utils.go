package csvmodule

import (
	"encoding/csv"
	"julian_project/modules/core"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Unmarshal(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := core.GetPublicFieldsReflect(v)
	listFields := core.GetPublicFields(v)
	if len(s) != len(record) {
		return &FieldMismatch{len(s), len(record)}
	}
	val := reflect.ValueOf(v)
	for i := 0; i < len(listFields); i++ {
		field := listFields[i]
		f := reflect.Indirect(val).FieldByName(field)
		fieldValue := f.Interface()
		switch fieldValue.(type) {
		case string:
			f.SetString(record[i])
		case int:
			ival, err := strconv.ParseInt(record[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		case float32:
			ival, err := strconv.ParseFloat(record[i], 32)
			if err != nil {
				return err
			}
			f.SetFloat(ival)
		case bool:
			ival, err := strconv.ParseBool(record[i])
			if err != nil {
				return err
			}
			f.SetBool(ival)
		case time.Time:
			date, err := time.Parse(time.ANSIC, record[i])
			if err != nil {
				return err
			}

			f.Set(reflect.ValueOf(date))
		case []string:
			r := reflect.ValueOf(strings.Split(record[i], ","))
			f.Set(r)
		default:
			return &UnsupportedType{f.Type().String()}
		}
	}

	return nil
}

type FieldMismatch struct {
	expected, found int
}

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}

type UnsupportedType struct {
	Type string
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.Type
}

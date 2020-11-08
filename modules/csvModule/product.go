package csvmodule

import (
	"encoding/csv"
	"fmt"
	"io"
	"julian_project/modules/core"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//Producto manejo de productos en los archivos CSVs
type Producto struct {
	core.Producto
}

//GuardarProductos Guardar una lista de Productos en un csv
func GuardarProductos(lp []Producto) {
	csvfile, err := os.Create("./store/productos.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvfile)
	lengthProducts := len(lp)
	for i := 0; i < lengthProducts; i++ {
		p := lp[i]
		if i == 0 {
			csvwriter.Write(p.GetPublicFields())
		}
		csvwriter.Write(p.GetPublicValues())

	}
	csvwriter.Flush()
	csvfile.Close()
	fmt.Printf("This is a Save")
}

func (p *Producto) GetPublicFields() []string {
	return p.Producto.GetPublicFields()
}

func (p *Producto) GetPublicValues() []string {
	return p.Producto.GetPublicValues()
}

func ReadProductos() []Producto {
	csvfile, err := os.Open("./store/productos.csv")
	if err != nil {
		log.Fatalf("failed open file: %s", err)
	}
	var reader = csv.NewReader(csvfile)
	reader.Comma = ','
	var productos []Producto
	reader.Read()
	var prod Producto
	for {
		err := Unmarshal(reader, &prod.Producto)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		productos = append(productos, prod)
	}
	return productos
}

// func (p *Producto) CsvParse(in string) {
// 	tmp := strings.Split(in, ",")
// 	p.ID= tmp[0]
// 	p.CantidadUnidad= 50,
// 	p.Nombre: "Marcosasdsa"
// 	p.Descripcion: "Redsady"
// 	p.Medidas: "15*454"
// 	p.Empresa: "MarcosSA2"
// 	p.Codigo: "2459"
// 	p.Precio:     154,
// 	p.Categorias: []string{"sad", "sad", "154"},
// 	p.Activo:     false,
// }

func Unmarshal(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(v).Elem()
	fmt.Printf("Fields are %v", s.NumField())
	fmt.Printf("Fields are for csv %v", len(record))
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":
			f.SetString(record[i])
		case "int":
			ival, err := strconv.ParseInt(record[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		case "float32":
			ival, err := strconv.ParseFloat(record[i], 32)
			if err != nil {
				return err
			}
			f.SetFloat(ival)
		case "[]string":
			v := reflect.ValueOf(strings.Split(record[i], ","))
			f.Set(v)
		case "bool":
			ival, err := strconv.ParseBool(record[i])
			if err != nil {
				return err
			}
			f.SetBool(ival)
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

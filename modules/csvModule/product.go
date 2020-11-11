package csvmodule

import (
	"encoding/csv"
	"fmt"
	"io"
	"julian_project/modules/core"
	"log"
	"os"
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
	fmt.Println("This is a Save")
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
	csvfile.Close()
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

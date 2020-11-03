package csvmodule

import (
	"encoding/csv"
	"fmt"
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
			csvwriter.Write(p.GetFields())
		}
		csvwriter.Write(p.GetValues())

	}
	csvwriter.Flush()
	csvfile.Close()
	fmt.Printf("This is a Save")
}

func (p *Producto) GetFields() []string {
	return p.Producto.GetFields()
}

func (p *Producto) GetValues() []string {
	return p.Producto.GetValues()
}

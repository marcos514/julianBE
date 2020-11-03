package csvmodule

import (
	"julian_project/modules/core"
)

//Factura manejo de productos en los archivos CSVs
type Factura struct {
	core.Factura
}

//FacturaProducto manejo de productos en los archivos CSVs
type FacturaProducto struct {
	core.FacturaProducto
}

// //GuardarProductos Guardar una lista de Productos en un csv
// func GuardarProductos(lp []Producto) {
// 	csvfile, err := os.Create("./store/productos.csv")
// 	if err != nil {
// 		log.Fatalf("failed creating file: %s", err)
// 	}
// 	csvwriter := csv.NewWriter(csvfile)
// 	lengthProducts := len(lp)
// 	for i := 0; i < lengthProducts; i++ {
// 		p := lp[i]
// 		if i == 0 {
// 			csvwriter.Write(p.GetFields())
// 		}
// 		csvwriter.Write(p.GetValues())

// 	}
// 	csvwriter.Flush()
// 	csvfile.Close()
// 	fmt.Printf("This is a Save")

// }

func (f *Factura) GetFields() []string {
	return f.Factura.GetFields()
}

func (p *Factura) GetValues() []string {
	return p.Factura.GetValues()
}

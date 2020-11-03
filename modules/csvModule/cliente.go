package csvmodule

import (
	"julian_project/modules/core"
)

//Cliente manejo de productos en los archivos CSVs
type Cliente struct {
	core.Cliente
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

func (p *Cliente) GetFields() []string {
	return p.Cliente.GetFields()
}

func (p *Cliente) GetValues() []string {
	return p.Cliente.GetValues()
}

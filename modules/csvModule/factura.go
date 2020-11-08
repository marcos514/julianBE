package csvmodule

import (
	"encoding/csv"
	"fmt"
	"julian_project/modules/core"
	"log"
	"os"
)

//Factura manejo de productos en los archivos CSVs
type Factura struct {
	core.Factura
}

//FacturaProducto manejo de productos en los archivos CSVs
type FacturaProducto struct {
	core.FacturaProducto
}

//GuardarProductos Guardar una lista de Productos en un csv
func GuardarFacturas(lf []Factura) {
	csvFacturas, err := os.Create("./store/facturas.csv")
	csvFacturasProductos, err := os.Create("./store/facturas_productos.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	facturasWriter := csv.NewWriter(csvFacturas)
	facturasProductosWriter := csv.NewWriter(csvFacturasProductos)
	lengthProducts := len(lf)
	for i := 0; i < lengthProducts; i++ {
		f := lf[i]
		if i == 0 {
			facturasWriter.Write(f.GetPublicFields())
		}
		facturasWriter.Write(f.GetValues())
		GuardarFacturaProductos(f.GetFacturaProducto(), facturasProductosWriter, i)
	}
	facturasWriter.Flush()
	facturasProductosWriter.Flush()
	csvFacturas.Close()
	csvFacturasProductos.Close()
	fmt.Printf("This is a Save")
}

//GuardarProductos Guardar una lista de Productos en un csv
func GuardarFacturaProductos(lfp []core.FacturaProducto, w *csv.Writer, index int) {
	lengthProducts := len(lfp)
	for i := 0; i < lengthProducts; i++ {
		fp := lfp[i]
		if i == 0 {
			w.Write(fp.GetPublicFields())
		}
		w.Write(fp.GetPublicValues())
	}
	fmt.Printf("This is a Save")
}

func (f *Factura) GetPublicFields() []string {
	return f.Factura.GetPublicFields()
}

func (f *Factura) GetValues() []string {
	return f.Factura.GetPublicValues()
}

func (fp *FacturaProducto) GetPublicFields() []string {
	return fp.FacturaProducto.GetPublicFields()
}

func (fp *FacturaProducto) GetValues() []string {
	return fp.FacturaProducto.GetPublicValues()
}

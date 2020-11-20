package csvmodule

import (
	"encoding/csv"
	"fmt"
	"io"
	"julian_project/modules/core"
	"log"
	"os"
	"time"
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
		} else {
			f.Fecha.Format(time.ANSIC)
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

func ReadFacturas() []Factura {
	csvfile, err := os.Open("./store/facturas_productos.csv")
	if err != nil {
		log.Fatalf("failed open file: %s", err)
	}
	var reader = csv.NewReader(csvfile)
	reader.Comma = ','
	var facturas []Factura
	var facturasProductos []FacturaProducto
	reader.Read()
	var fac Factura
	var facProd FacturaProducto
	lp := ReadProductos("./store/productos")
	productosMap := MapProducts(lp)
	//Read Factura Productos
	for {
		err := Unmarshal(reader, &facProd.FacturaProducto)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		facProd.AddProducto(productosMap[facProd.ProductoID].Producto)
		facturasProductos = append(facturasProductos, facProd)
	}
	facturaProductosIds := GetFacturaProductosByIds(facturasProductos)
	csvfile.Close()
	lc := ReadClientes()
	clientesMap := MapClientes(lc)

	csvfile, err = os.Open("./store/facturas.csv")
	if err != nil {
		log.Fatalf("failed open file: %s", err)
	}
	reader = csv.NewReader(csvfile)
	reader.Comma = ','
	reader.Read()
	//Read Facturas
	for {
		err := Unmarshal(reader, &fac.Factura)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fac.AppendListFacturasProductos(facturaProductosIds[fac.ID])
		fac.SetCliente(clientesMap[fac.ClienteID].Cliente)
		facturas = append(facturas, fac)
	}
	csvfile.Close()
	return facturas
}

func GetFacturaProductosByIds(lfp []FacturaProducto) map[int][]FacturaProducto {
	facturaProductosDict := make(map[int][]FacturaProducto)
	for i := 0; i < len(lfp); i++ {
		fp := lfp[i]
		facturaProductosDict[fp.ID] = append(facturaProductosDict[fp.ID], []FacturaProducto{fp}...)
	}
	return facturaProductosDict
}

func (f *Factura) AppendListFacturasProductos(lfp []FacturaProducto) {
	f.Factura.AppendListFacturasProductos(ConvertFactProducto(lfp))
}

func ConvertFactProducto(lfp []FacturaProducto) []core.FacturaProducto {
	var facturaProductos []core.FacturaProducto
	for i := 0; i < len(lfp); i++ {
		fp := lfp[i]
		facturaProductos = append(facturaProductos, fp.FacturaProducto)
	}
	return facturaProductos
}

func (f *Factura) PrintFactura() {
	fmt.Printf(
		`
		Factura:
			ID: %v
			cliente: %v
			Fecha: %v
			PrecioTotal: %v
			facturaProductos: %v
		`, f.ID, f.GetCliente(), f.Fecha.Format(time.ANSIC), f.PrecioTotal, f.GetFacturaProducto()[0].GetFactura(),
	)
}

func AgregarFactura(f Factura) []Factura {
	lf := ReadFacturas()
	lastFacturaId := lf[len(lf)-1].ID
	f.ID = lastFacturaId
	lf = append(lf, f)
	GuardarFacturas(lf)
	return lf
}

func ActualizarFactura(f Factura) []Factura {
	lf := ReadFacturas()
	findex := f.IndexFacturaEnLista(lf)
	if findex == -1 {
		lastFacturaId := lf[len(lf)-1].ID
		f.ID = lastFacturaId
		lf = append(lf, f)
	} else {
		lf[findex] = f
	}
	GuardarFacturas(lf)
	return lf
}

func (f *Factura) IndexFacturaEnLista(lf []Factura) int {
	index := -1
	for i := 0; i < len(lf); i++ {
		faux := lf[i]
		if (f.ClienteID == faux.ClienteID && f.Fecha == faux.Fecha) || f.ID == faux.ID {
			index = i
			break
		}
	}
	return index
}

func (f *Factura) ImprimirFacturaCSV() {

}

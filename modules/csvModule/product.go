package csvmodule

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/marcos514/julianBE/modules/core"
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

func ReadProductos(fileName string) []Producto {
	csvfile, err := os.Open(fileName)
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

func MapProducts(lc []Producto) map[int]Producto {
	mapProducto := make(map[int]Producto)
	for i := 0; i < len(lc); i++ {
		c := lc[i]
		mapProducto[c.ID] = c
	}
	return mapProducto
}

func MapToSliceProducts(mp map[int]Producto) []Producto {
	var lp []Producto
	for _, p := range mp {
		lp = append(lp, p)
	}
	return lp
}

func DeshabilitarProductosById(ids []int) []Producto {
	mp := MapProducts(ReadProductos("./store/productos.csv"))
	for i := 0; i < len(ids); i++ {
		p := mp[ids[i]]
		p.Activo = false
		mp[ids[i]] = p
	}
	lp := MapToSliceProducts(mp)
	GuardarProductos(lp)
	return lp
}

func AgregarProductosDeArchivo(filename string) []Producto {
	lp := ReadProductos(filename)
	for i := 0; i < len(lp); i++ {
		p := lp[i]
		p.Activo = true
		lp[i] = p
	}
	return ActualizarProductos(lp)
}

func ActualizarProductos(newlp []Producto) []Producto {
	lp := ReadProductos("./store/productos.csv")
	lastProductId := lp[len(lp)-1].ID
	for i := 0; i < len(newlp); i++ {
		p := newlp[i]
		pindex := p.IndexProductoEnLista(lp)
		if pindex == -1 {
			lastProductId += 1
			p.ID = lastProductId
			lp = append(lp, p)
		} else {
			paux := lp[pindex]
			p.ID = paux.ID
			lp[pindex] = p
		}
	}
	GuardarProductos(lp)
	return lp
}

func (p *Producto) IndexProductoEnLista(lp []Producto) int {
	index := -1
	if p.ID == -6 {
		for i := 0; i < len(lp); i++ {
			paux := lp[i]
			if p.ID == paux.ID || p.Nombre == paux.Nombre && p.Codigo == paux.Codigo && p.Empresa == paux.Empresa {
				index = i
				break
			}
		}
	}
	return index
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

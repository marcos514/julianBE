package main

import (
	"encoding/json"
	"fmt"
	"julian_project/modules/core"
	csvmodule "julian_project/modules/csvModule"
	"time"
	//"julian_project/morestrings"
)

type empData struct {
	Name string
	Age  string
	City string
}

// type Cliente struct {
// id        int
// mail      string
// nombre    string
// direccion string
// numero    string
// }
type Usuario struct {
	nombre   string
	apellido string
}

func (u Usuario) getDatosPersonales() string {
	return fmt.Sprintf("%s, %s", u.nombre, u.apellido)
}

type Administrador struct {
	Usuario
	sector string
}

func (a Administrador) getDatosCompletos() string {
	return fmt.Sprintf("%s - %s", a.getDatosPersonales(), a.sector)
}

type Primate interface {
	Alimentar(string)
}

type Antropoide struct{}

func (t Antropoide) Alimentar(fruta string) {
	fmt.Printf("Comiendo %s \n", fruta)
}

type Gorila struct {
	Antropoide
}

func DarDeComer(primate Primate) {
	primate.Alimentar("banana")
}

func main() {

	//fmt.Println(morestrings.ReverseRunes("!oG ,ollasdasdeH"))

	var clientes = core.Cliente{}

	core.ReadFromFile("cliente.json", &clientes)
	res2B, _ := json.Marshal(clientes)
	fmt.Println(string(res2B))

	var listClientes = []core.Cliente{
		{
			ID:        1,
			Mail:      "marmarreyer@gmail.com",
			Nombre:    "Marcos",
			Direccion: "Amenedo 622",
			Numero:    "1549168959",
		},
		{
			ID:        2,
			Mail:      "reymarcos51@gmail.com",
			Nombre:    "Marcos Rey",
			Direccion: "Amenedo 622",
			Numero:    "1549168959",
		},
	}
	core.WriteData("clientes_2.json", listClientes)

	var clienteTest = core.Cliente{
		ID:        1,
		Mail:      "marmarreyer@gmail.com",
		Nombre:    "Marcos",
		Direccion: "Amenedo 622",
		Numero:    "1549168959",
	}
	core.WriteData("marcos.json", clienteTest)

	// archivo.readFromFile("asd.json", cliente)
	// fmt.Println(archivo.addLine())
	// csvFile, err := os.Open("emp2.csv")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(csvFile)
	// defer csvFile.Close()

	// csvLines, err := csv.NewReader(csvFile).ReadAll()
	// if err != nil {
	// 	fmt.Println(csvLines)
	// }
	// for other, line := range csvLines {
	// 	fmt.Println(other)
	// 	emp := empData{
	// 		Name: line[0],
	// 		Age:  line[1],
	// 		City: line[2],
	// 	}
	// 	fmt.Println(emp.Name + " " + emp.Age + " " + emp.City)
	// }
	var administrador = Administrador{Usuario{"Jose", "Luis"}, "Computos"}

	fmt.Println(administrador.getDatosPersonales())
	fmt.Println(administrador.getDatosCompletos())
	kong := Gorila{}
	DarDeComer(kong)

	products := []csvmodule.Producto{{
		Producto: core.Producto{
			ID: 15, CantidadUnidad: 50,
			Nombre: "Marcos", Descripcion: "Rey", Medidas: "15*45", Empresa: "MarcosSA", Codigo: "1234",
			Precio:     154,
			Categorias: []string{"marcos"},
			Activo:     false,
		},
	}}
	prod := csvmodule.Producto{
		core.Producto{
			ID: 15, CantidadUnidad: 50,
			Nombre: "Marcosasdsa", Descripcion: "Redsady", Medidas: "15*454", Empresa: "MarcosSA2", Codigo: "2459",
			Precio:     154,
			Categorias: []string{"sad", "sad", "154"},
			Activo:     false,
		},
	}
	products = append(products, prod)
	csvmodule.GuardarProductos(products)

	fact := csvmodule.Factura{
		core.Factura{
			ID:          1,
			ClienteID:   2,
			Fecha:       time.Now(),
			PrecioTotal: 12,
		},
	}

	var facts []csvmodule.Factura

	factProd := csvmodule.FacturaProducto{
		core.FacturaProducto{
			ID:         1,
			ProductoID: prod.ID,
			FacturaID:  fact.ID,
			Precio:     prod.Precio,
			Cantidad:   5,
		},
	}
	fact.AddFacturaProduct(factProd.FacturaProducto)
	facts = append(facts, fact)

	factProd.AddFact(fact.Factura)
	factProd.AddProd(prod.Producto)

	csvmodule.GuardarFacturas(facts)

	ahora := time.Now()
	fmt.Printf("Ahora %v", ahora)

	cliente := []csvmodule.Cliente{{
		core.Cliente{
			ID:        0,
			Mail:      "marcos@smallsforsmalls.com",
			Nombre:    "Marcos Rey",
			Direccion: "Amenedo 622",
			Numero:    "1549168959",
		},
	}}
	csvmodule.GuardarClientes(cliente)
	csvproductos := csvmodule.ReadProductos()
	fmt.Println("%v", csvproductos)
}

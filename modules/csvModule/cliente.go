package csvmodule

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/marcos514/julianBE/modules/core"
)

//Cliente manejo de productos en los archivos CSVs
type Cliente struct {
	core.Cliente
}

//GuardarProductos Guardar una lista de Productos en un csv
func GuardarClientes(lc []Cliente) {
	csvfile, err := os.Create("./store/clientes.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvfile)
	lengthCliente := len(lc)
	for i := 0; i < lengthCliente; i++ {
		c := lc[i]
		if i == 0 {
			csvwriter.Write(c.GetPublicFields())
		}
		csvwriter.Write(c.GetPublicValues())

	}
	csvwriter.Flush()
	csvfile.Close()
	fmt.Println("This is a Save")

}

func (c *Cliente) GetPublicFields() []string {
	return c.Cliente.GetPublicFields()
}

func (c *Cliente) GetPublicValues() []string {
	return c.Cliente.GetPublicValues()
}

func ReadClientes() []Cliente {
	csvfile, err := os.Open("./store/clientes.csv")
	if err != nil {
		log.Fatalf("failed open file: %s", err)
	}
	var reader = csv.NewReader(csvfile)
	reader.Comma = ','
	var clientes []Cliente
	reader.Read()
	var cli Cliente
	for {
		err := Unmarshal(reader, &cli.Cliente)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		clientes = append(clientes, cli)
	}
	csvfile.Close()
	return clientes
}

func MapClientes(lc []Cliente) map[int]Cliente {
	mapClientes := make(map[int]Cliente)
	for i := 0; i < len(lc); i++ {
		c := lc[i]
		mapClientes[c.ID] = c
	}
	return mapClientes
}

func AgregarCliente(c Cliente) []Cliente {
	lc := ReadClientes()
	lastClienteId := lc[len(lc)-1].ID
	c.ID = lastClienteId + 1
	lc = append(lc, c)
	GuardarClientes(lc)
	return lc
}

func ActualizarCliente(c Cliente) []Cliente {
	lc := ReadClientes()
	cindex := c.IndexClienteEnLista(lc)
	if cindex == -1 {
		lastFacturaId := lc[len(lc)-1].ID
		c.ID = lastFacturaId
		lc = append(lc, c)
	} else {
		lc[cindex] = c
	}
	GuardarClientes(lc)
	return lc
}

func (c *Cliente) IndexClienteEnLista(lc []Cliente) int {
	index := -1
	for i := 0; i < len(lc); i++ {
		caux := lc[i]
		if c.Mail == caux.Mail || c.ID == caux.ID {
			index = i
			break
		}
	}
	return index
}

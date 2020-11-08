package csvmodule

import (
	"encoding/csv"
	"fmt"
	"julian_project/modules/core"
	"log"
	"os"
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

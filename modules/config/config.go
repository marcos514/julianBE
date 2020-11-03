// Package config implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package config

import (
	"julian_project/modules/archivo"
)

//Config todos los datos del sistema
type Config struct {
	ProximoCliente  int `json:"proximo_cliente"`
	ProximoProducto int `json:"proximo_producto"`
	ProximoFactura  int `json:"ultima_factura"`
}

// ProximoCliente returns its argument string reversed rune-wise left to right.
func ProximoCliente() int {
	c := Config{}
	archivo.ReadFromFile("config.json", &c)
	var ret = c.ProximoCliente
	return ret
}

// NuevoCliente returns its argument string reversed rune-wise left to right.
func (c Config) actualizarCliete() int {
	var ret = c.ProximoCliente
	return ret
}

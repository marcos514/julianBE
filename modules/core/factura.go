// Package core implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package core

import "time"

//"time"

//Factura this is a factura
type Factura struct {
	ID          int       `json:"id"`
	ClienteID   int       `json:"cliente"`
	Fecha       time.Time `json:"fecha"`
	PrecioTotal float32   `json:"precio_total"`
	productos   []FacturaProducto
}

//FacturaProducto los productos de la factura
type FacturaProducto struct {
	ID         int `json:"id"`
	producto   Producto
	ProductoID int `json:"producto_id"`
	FacturaID  int `json:"factura"`
	factura    Factura
	Precio     float32 `json:"precio"`
	Cantidad   int     `json:"cantidad"`
}

type FacturaProductoInterface interface {
	ThisIsFactProd()
}

//GetPublicFields Guardar Product
func (f *Factura) GetPublicFields() []string {
	return GetPublicFields(f)
}

//GetPublicValues Guardar Product
func (f *Factura) GetPublicValues() []string {
	return GetPublicValues(f)
}

func (f *Factura) GetFacturaProducto() []FacturaProducto {
	return f.productos
}

func (fp *FacturaProducto) GetProduct() Producto {
	return fp.producto
}

func (fp *FacturaProducto) GetFactura() Factura {
	return fp.factura
}

//GetPublicFields Guardar Product
func (fp *FacturaProducto) GetPublicFields() []string {
	return GetPublicFields(fp)
}

//GetPublicValues Guardar Product
func (fp *FacturaProducto) GetPublicValues() []string {
	return GetPublicValues(fp)
}

func (f *Factura) AddFacturaProduct(fp FacturaProducto) {
	f.productos = append(f.productos, fp)
}

func (fp *FacturaProducto) AddFact(f Factura) {
	fp.factura = f
}

func (fp *FacturaProducto) AddProd(p Producto) {
	fp.producto = p
}

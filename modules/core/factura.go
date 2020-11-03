// Package core implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package core

//"time"

//Factura this is a factura
type Factura struct {
	ID          int     `json:"id"`
	ClienteID   int     `json:"cliente"`
	Fecha       string  `json:"fecha"`
	PrecioTotal float32 `json:"precio_total"`
}

//FacturaProducto los productos de la factura
type FacturaProducto struct {
	ID         int `json:"id"`
	producto   Producto
	ProductoID int     `json:"producto_id"`
	FacturaID  Factura `json:"factura"`
	factura    Factura
	Precio     float32 `json:"precio"`
	Cantidad   int     `json:"cantidad"`
}

//GetFields Guardar Product
func (f *Factura) GetFields() []string {
	return GetFields(f)
}

//GetValues Guardar Product
func (f *Factura) GetValues() []string {
	return GetValues(f)
}

//GetFields Guardar Product
func (p *FacturaProducto) GetFields() []string {
	return GetFields(p)
}

//GetValues Guardar Product
func (p *FacturaProducto) GetValues() []string {
	return GetValues(p)
}

// Package core implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package core

//Producto this is a product
type Producto struct {
	ID, CantidadUnidad                            int
	Nombre, Descripcion, Medidas, Empresa, Codigo string
	Precio                                        float32
	Categorias                                    []string
	Activo                                        bool
}

// //ProductoInterface para usar las funciones en todos los lados
// type ProductoInterface []interface {
// 	GuardarProducto()
// }

// //ProductosInterface para usar las funciones en todos los lados
// type ProductosInterface []interface {
// 	GuardarProductos()
// }

//GuardarProducto Guardar Product
func (p Producto) GuardarProducto() []string {
	return p.Categorias
}

//GetFields Guardar Product
func (p *Producto) GetPublicFields() []string {
	return GetPublicFields(p)
}

//GetPublicValues Guardar Product
func (p *Producto) GetPublicValues() []string {
	return GetPublicValues(p)
}

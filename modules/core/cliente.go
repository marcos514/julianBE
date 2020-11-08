// Package core implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package core

//Cliente this is a client
type Cliente struct {
	ID        int
	Mail      string
	Nombre    string
	Direccion string
	Numero    string
}

// CrearCliente returns its argument string reversed rune-wise left to right.
func CrearCliente(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//GetPublicFields Guardar Product
func (p *Cliente) GetPublicFields() []string {
	return GetPublicFields(p)
}

//GetPublicValues Guardar Product
func (p *Cliente) GetPublicValues() []string {
	return GetPublicValues(p)
}

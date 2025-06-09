package dhp

import "fmt"

// VarExport печатает значение как Go-код
func VarExport(value interface{}) {
	fmt.Printf("%#v\n", value)
}

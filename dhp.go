package dhp

import (
	"fmt"
	"time"
	"github.com/davecgh/go-spew/spew"
)

// -----[ PRINT HELPERS ]-----
// VarExport печатает значение как Go-код
func VarExport(value interface{}) {
	fmt.Printf("%#v\n", value)
}

// -----[ PRINT HELPERS ]-----
// Sleep - более лаконичная обертка 
func Sleep(countSeconds int) {
	time.Sleep(time.Second * time.Duration(countSeconds))
}

// -----[ PRINT HELPERS ]-----
// Dump - более лаконичная обертка 
func Dump(value interface{}) {
	spew.Dump(value)
}

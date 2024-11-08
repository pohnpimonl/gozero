package pkg

import "fmt"

type Personal struct {
	Name    string // First capital letter mean export can access in public
	Age     int
	address string // cannot access via public
}

// func (p Personal) address() {
// 	fmt.Println("Bangkok")
// }

func (p Personal) Myaddress() {
	fmt.Println("Bangkok")
}

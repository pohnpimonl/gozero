package main

import "fmt"

type insuranceType string

var UL insuranceType = "Unit Link"
var UN insuranceType = "Universal Life"
var WL insuranceType = "Whole Life"

func (i insuranceType) ShowFullDisplay() {
	fmt.Printf("This is full display of insurance: %s\n", i)
}

var unitLink = "Unit Link"

// type insuranceList []insuranceType

// func (i *insuranceList) initilizeInsurance() {
// 	*i = append(*i, UL)
// 	*i = append(*i, UN)
// 	*i = append(*i, WL)
// }

// func (i insuranceList) Display() {
// 	for _, v := range i {
// 		fmt.Printf("insurance %s\n", v)
// 	}
// }

type errorState int

func (e errorState) Error() string {
	return fmt.Sprintf("error state is %d, mean user error", e)
}

func main() {

	// fmt.Printf("insurance type: %+v type:%T\n", UL, UL)
	// fmt.Printf("string type: %+v type:%T\n", unitLink, unitLink)

	// s := isInsurance()
	// fmt.Printf("insurance: %s", s)

	/* --- add method to type --- */
	// var is insuranceList

	// is.initilizeInsurance()

	// is.Display()

	/* --- implement interface error  --- */
	// e := retErrState()
	// fmt.Println(e.Error())

	/* --- add method to insurance type ---- */
	// UL.ShowFullDisplay()

}

// func retErrState() error {
// 	e := errorState(5)
// 	return e
// }

// func isInsurance() insuranceType {
// 	return UL
// }

// func isInsurance() string {
// 	return string(UL)
// }

package main

import "fmt"

func main() {

	/*
		in Go, maps are not safe for concurrent use by default.
		If multiple goroutines attempt to read from and write to a map simultaneously,
		this can lead to race conditions or runtime panics. Specifically,
		the Go runtime will panic with the error
		fatal error: concurrent map read and map write when it detects concurrent access to a map.
	*/

	/* ---------- initlize map with make ------- */
	// m := make(map[string]string)

	// m["a"] = "first"
	// m["b"] = "second"
	// m["c"] = "third"

	// fmt.Printf("maps m: %+v\n", m)

	// for k, v := range m {
	// 	fmt.Printf("key: %s, value: %s\n", k, v)
	// }

	/*----- delete maps -------*/
	// delete(m, "b")

	// fmt.Println("after delete key b")

	// for k, v := range m {
	// 	fmt.Printf("key: %s, value: %s\n", k, v)
	// }

	/*------- clear maps ------*/

	// clear(m)

	// fmt.Println("after clear m")
	// for k, v := range m {
	// 	fmt.Printf("key: %s, value: %s\n", k, v)
	// }

	// fmt.Printf("after clear maps m: %+v\n", m)

	/* --- initialize with out make  ----- */

	// var m map[string]string
	// m["a"] = "first"

	// fmt.Printf("maps m: %+v\n", m)

	// m := map[string]string{
	// 	"a": "first",
	// 	"b": "second",
	// 	"c": "third",
	// }

	/* ---- check value in map ----*/
	// if v, ok := m["a"]; ok {
	// 	fmt.Printf("value is: %s", v)
	// }

	// if _, ok := m["a"]; ok {
	// 	fmt.Printf("key a exist")
	// }

	// _, ok := m["a"]
	// if ok {
	// 	fmt.Printf("key a exist")
	// }

	// fmt.Printf("maps m: %+v\n", m)

	/* ---------- compare map ---------*/
	// m1 := map[string]string{
	// 	"a": "first",
	// 	"b": "second",
	// 	"c": "third",
	// }

	// m2 := map[string]string{
	// 	"a": "first",
	// 	"b": "second",
	// 	"c": "third",
	// }

	// fmt.Printf("m1 == m2: %t\n", maps.Equal(m1, m2))

	// m1 := map[string]string{
	// 	"a": "first",
	// 	"b": "second",
	// 	"c": "third",
	// }

	// m2 := map[string]string{
	// 	"a": "first",
	// 	"b": "second",
	// 	"c": "cccc",
	// }

	// fmt.Printf("m1 == m2: %t\n", maps.Equal(m1, m2))

	// m := make(map[string]string)
	// m["a"] = "a"
	// m["b"] = "b"

	// fmt.Printf("maps before edit %+v\n", m)

	// editMap(m)

	// fmt.Printf("maps after edit %+v\n", m)

}

func editMap(m map[string]string) {
	m["a"] = "this new a"
}

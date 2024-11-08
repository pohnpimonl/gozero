package main

import (
	"fmt"
	"runtime"
)

func main() {

	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// 	if i == 10 {
	// 		myRecovery()
	// 	}
	// }

	testDefer()

}

func myPanic() {
	panic("just from myPanic")
}

func myRecovery() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from panic: ", r)
		}
	}()

	panic("just from myRecovery")
}

func myRecoveryWithTrace() {

	defer func() {
		if r := recover(); r != nil {
			var buf [4096]byte
			v := runtime.Stack(buf[:], false)
			fmt.Println("recover from panic: ", r)
			fmt.Printf("stack trace %v\n", string(buf[:v]))
		}
	}()

	panic("just from myRecovery")
}

func testDefer() int {
	fmt.Println("Function called")

	// defer statements are executed after the return statement is evaluated
	// but before the result is returned to the caller
	defer fmt.Println("Deferred function executed")

	fmt.Println("Before return")

	return 10
}

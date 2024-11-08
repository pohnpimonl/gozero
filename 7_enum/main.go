package main

import (
	"fmt"
)

/* enum in go it more like constant */

// Define a custom type for the enum
type Day int

// Using iota to define enum values
const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	today := Monday

	switch today {
	case Monday:
		fmt.Println("Start of the workweek")
	case Friday:
		fmt.Println("End of the workweek")
	default:
		fmt.Println("It's a regular day")
	}
}

// public enum Day {
//     MONDAY("Start of the week"),
//     FRIDAY("Almost weekend"),
//     SUNDAY("Weekend");

//     private String description;

//     // Constructor
//     Day(String description) {
//         this.description = description;
//     }

//     public String getDescription() {
//         return description;
//     }
// }

// public class EnumExample {
//     public static void main(String[] args) {
//         System.out.println(Day.MONDAY.getDescription()); // Start of the week
//     }
// }

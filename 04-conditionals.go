package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Exploring Conditionals if and switch.
// Boolean comparison operators ==, !=, <, <=, >, >=, &&, ||
// Remember: The curly braces should be on the same line as the if statement.
// Compiler inserts ; as needed.
// if <boolean_expression> {
//    <run this code if condition is satisfied.
// } else if <boolean_expression> {
//    <another block of code>
// } else {
//    <another block of code>
// }

// if also accepts simple statements before the condition evaluation.
// The variables declared here are scoped to the if statement. If supports nesting.
// if <simple statements> ; <boolean_expression> {
//     <block of code>
// }

// switch keyword - both the simple statement and expression are optional,
// but if you have a simple statement, it should be followed by a ;
// irrespective of the expression.
// Note: Go doesn't need a break after completing execution for an if.
// swtich <simple statement>; <expression> {
// case <expression> : <code>
// case <expression> : <code>
// case <expression> : <code>
// default : <code>
// }

func main() {
	fmt.Println("Exploring Conditional Flow")
	// firstRank := 39
	// secondRank := 614

	// if firstRank < secondRank {
	// 	fmt.Println("First book is doing better than the second.")
	// } else if firstRank > secondRank {
	// 	fmt.Println("Second book is doing better than the first.")
	// } else {
	// 	fmt.Println("Both courses have same rating.")
	// }

	// If statement above has been rewritten with simple statements.
	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("First book is doing better than the second.")
	} else if secondRank < firstRank {
		fmt.Println("Second book is doing better than the first.")
	} else {
		fmt.Println("Both courses have same rating.")
	}

	// switch statement example:
	// implicit breaks!
	// If you want to fallthrough, go ahead and add the fallthrough keyword!
	// fallthrough works only on a case by case basis,
	// if you need to fallthrough all cases, you need need to add this keyword to respective case.
	switch course := "docker"; course {
	case "linux":
		fmt.Println("You chose linux!")
	case "docker":
		fmt.Println("You chose docker!")
		fallthrough // This keyword lets control fall through to next case and execute that!
	case "windows":
		fmt.Println("You chose windows!")
	default:
		fmt.Println("Sorry we couldn't find your selection!!")
	}

	// Note case values can be comma separated. Say a list of values.
	// e.g.
	// case 0, 2, 4, 6, 8:
	// case 1, 3, 5, 7, 9:
	// A more realistic switch example.
	number := random()
	switch result := number % 2; result {
	case 0:
		fmt.Println("Even Number: ", number)
	case 1:
		fmt.Println("Odd Number: ", number)
	}

	// Role of if in Error Handling
	// It is idiomatic to return an error as the last return from functions and methods.
	// nil is used to indicate success.
	// func testConn(target string) (rspTime float64 err error){
	//     if err != nil {
	//         <error handling code>
	//     }
	// }

	// IMPORTANT NOTE: Every function should consider having the last parameter as error value returned. It is a good practice!
	// Also consider checking for if err !== nil, for every function.

	_, err := os.Open("D:\\Srihari\\Test.txt")
	if err != nil {
		fmt.Println("Error opening file! Please check if the file exists.", err)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}

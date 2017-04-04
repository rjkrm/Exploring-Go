package main

import (
	"fmt"
	"time"
)

// Go just has a for keyword for looping. It is multi-purpose.
// Infinite loops, loop based on boolean expressions and loop over a range of values.
// Note: placement of curly brace is important.
// Syntax - blank expression is infinite loop.
// for <expression> {
//     <code>
// }

// index := 0
// for index < 10 {
//
//     index++
// }

// for index := 0; index < count; index++ {
//     <code>
// }

// for <variable> := range <list_of_values> {
//     <code>
// }

func main() {
	fmt.Println("Loops in Go!")
	// for pre;condition;post {
	// }
	// Pre statement is executed before the first iteration.
	// Post statement is executed for every iteration.
	fmt.Println("Demo for")
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Time completed!")
			break
		}

		if timer%2 == 0 {
			continue
		}

		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	// for range
	fmt.Println("Demo for ... range")
	courses := []string{"B.E", "M.B.A", "B.Tech"}
	// We don't care about the index values hence we assign them to underscore!
	for _, i := range courses {
		fmt.Println(i)
	}

	// If you need the index and item, then consider the below example!
	for index, item := range courses {
		fmt.Println("Index ", index, "has course ", item)
	}

	// We can have nested for loops.
	// Break and Continue
	// Break will break out of the immediate for loop.
	// Say you want to break out of innermost to outmost loop in a nested if (say 3 levels)
	// Use break to a label.
	// If you define a label you need to use it!
	// for var1 := range vars1 {
	//     breakPoint:
	//     for var2 := range vars2 {
	//         for var3 := range vars3 {
	//             break breakPoint
	//         }
	//     }
	// }

}

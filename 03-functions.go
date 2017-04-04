package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello Functions! This is really important")
	module := "learning functions"
	author := "srihari sridharan"
	fmt.Println(converter(module, author))

	bestFinishScores := bestLeagueFinishes(12, 10, 13, 13, 17, 14, 16, 7, 5, 2)
	fmt.Println("Best finish scores: ", bestFinishScores)
}

// Enable us to write clean, tidy and modular code.
// Go Function Syntax
// The curly braces should be Java or JS style on the same line. This is a must in Go!
// func functio_name(parameter data_type) return_type {
//     <code>
//     return <result>
// }

// Note: functions can return more than one value. The syntax is as shown below!
// We can use any variable name for parameters and return types.
// The variables module and author can be called s1 and s2 and return variables
// can be called str1 and str2, doesn't matter.
func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)
	return module, author
}

// Variadic Functions - you really don't know how many values are going to be passed in.
// The classic example is fmt.Println.
// Ellipses ... followed by data type.
func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}

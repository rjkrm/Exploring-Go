package main

import(
	"fmt"
)

func main(){
	fmt.Println("Structs in go!")
	// Custom datatypes
	// Collection of named fields each with own data types.
	// type Circle struct{
	// 	radius float64
	// 	diameter float64
	// 	circumference float64
	// }
	// Go doesn't have object, classes and inheritance, instead it is
	// more functional, we can achieve OO kind of design 
	// using modules and functions.

	type CourseMetadata struct{
		Name string
		Author string
		Level string
		Rating float64
	}

	// var learningGo CourseMetadata
	// learningGo := new (CourseMetadata)
	// The construction using new keyword gives a pointer.
	learningGo := CourseMetadata{
		Name : "Learning Go",
		Author: "Srihari",
		Level: "Beginner",
		Rating: 4.5, // you need the trailing comma, not as in JS or C#.
	}

	fmt.Println(learningGo)
	fmt.Println("Author Name: ", learningGo.Author)
	learningGo.Author = "Srihari Sridharan"
	fmt.Println("Updated Author Name:", learningGo.Author)


}
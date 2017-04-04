package main

import (
	"fmt"     // Standard i/o.
	"os"      // Provides operating system related info!
	"reflect" // For reflection of types
)

// Declaring variables at the package level use the var keyword.
// Variable names must start with _ or a letter
// var <variable_name> <type>
// or as below - Variables of same type on same line
// var (
//     <variable_name1>, <variable_name2> <type1>
//     <variable_name3> <type2>
// )
// var (
// 	name, course string
// 	module       float64
// )
// or as below - Variables on separate lines.
// var (
//     <variable_name1> <type1>
//     <variable_name2> <type2>
//     <variable_name3> <type3>
// )
// var (
// 	name   string
// 	course string
// 	module float64
// )

// Variables without any value are initialized as zero value
// func main() {
// 	fmt.Println("Name is set to: ", name)
// 	fmt.Println("Course is set to:", course)
// 	fmt.Println("Module is set to:", module)
// }

// const PI = 3.14 // This is how you define a constant.

// Go can infer types from values being set!
// Variables declared at the package level are global in scope and are available across the package.
var (
	// os.Getenv(key) the data from environment variables and in this case it gets the logged on user.
	name, course, module = os.Getenv("USERNAME"), "Learning Go", 1.1
)

// Use the reflect package to determine the type at runtime.
func main() {
	fmt.Println("Name is set to: ", name, "and is of type ", reflect.TypeOf(name))
	fmt.Println("Course is set to:", course, "and is of type ", reflect.TypeOf(course))
	fmt.Println("Module is set to:", module, "and is of type ", reflect.TypeOf(module))

	// Go is very particular on types as it is type safe! We cannot add an integer and float.
	// BTW: To declare a variable and set type inside a function use the := operator. := is the short assignment.
	// Can be used only within the function.
	// Further assignments can be done using = operator.
	a := 13.4
	b := 23
	// c := a + b // This line will result in an error as Go doesn't allow adding two different types.
	// .\learning-go.go:54: invalid operation: a + b (mismatched types float64 and int)
	// REMEMBER: Go will never allow unused variables within functions. However, at the module
	// level you can have unused variables / members!

	// In order to cast the variable use type(value_to_cast)
	c := a + float64(b) // Prints 36.4
	// or
	// c := int(a) + b // Prints 36
	fmt.Println("The total is ", c)

	// Pointers
	// & before a vaiable name denotes its address!
	// To declare a pointer to a variable:
	pointerToName := &name
	fmt.Println("Address of 'name' : ", pointerToName)
	fmt.Println("Address of 'name' : ", &name)
	fmt.Println("Value from address of a vaiable: ", *pointerToName)
	fmt.Println("Value from address of a vaiable: ", *(&name))

	// Go passes arguments by value and not by reference.
	// If a value is passed into a function, it gets a copy and only that copy gets modified.
	// Pass by value - Here the vaible name is passed by value.
	fmt.Println("Original course:", course)
	changeCourseByValue(course)
	fmt.Println("Changed course name after pass by value:", course)
	changeCourseByReference(&course)
	fmt.Println("Changed course name after pass by reference:", course)

	// Constants - immutable. const keyword. No shorthand := - not available for constants.
	fmt.Println("Operating System Environment Variables ", os.Environ()) // This dumps all into one huge line.

	// Prints one per line.
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}

// func <function_name> (arg type) <return_type>
// This kind of mimics the behavior of pass by reference.
func changeCourseByValue(course string) string {
	course = "Exploring Go in depth"
	return course // it is mandatory at the end of the function
}

// The '*' before the string data type indicates a pointer and it expects the memory address.
func changeCourseByReference(course *string) string {
	*course = "This is a new course!"
	return *course
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Arrays and Slices Demo")
	// Array vs Slices
	// Numbered lists of single type! YOU CANNOT HAVE ITEMS OF DIFFERENT TYPE!
	// Slices win hands-down!

	/*
	   Array
	   Numbered list of single type
	   Static or fixed length

	   Slices
	   Numbered list of single type.
	   Can be resized.
	   Passed by reference as they are references!
	   Reference contiguous portion of an array.
	   Length of the slice cannot be more than the array that is backing it.
	   Slice data is always stored in the array in the background.

	   Caveat need to be built on top of arrays!
	   All slices are built on top of arrays.
	   Slices are references, are pointers to real values.
	   Slices are cheap. All the slice has is type, length and capacity! make(<type>,<length>,<capacity>)
	*/

	// The [] brackets before string denotes that we are creating a slice of strings.
	// The hidden array in the background gets initialized with zero values. In this case for strings it is going to be empty strings.

	myCourses := make([]string, 5, 10)

	// len() gives length and cap() gives the capacity.
	fmt.Println("The length of the slice myCourses is ", len(myCourses), " and its capacity is ", cap(myCourses))

	// If we omit the capacity parameter it will have the same capacity as the length.
	myNewCourses := make([]string, 5)
	fmt.Println("The length of the slice myNewCourses is ", len(myNewCourses), " and its capacity is ", cap(myNewCourses))

	// Shorthand notation:
	myLatestCourses := []string{"B.E", "B.Tech", "M.B.A"}
	fmt.Println("The length of the slice myLatestCourses is ", len(myLatestCourses), " and its capacity is ", cap(myLatestCourses))

	// Access a specific element using [] brackets notation - 0 based index
	fmt.Println("Latest course ", myLatestCourses[1])
	// assign values using same bracket notation
	myLatestCourses[1] = "M.C.A"
	fmt.Println("Latest course ", myLatestCourses[1])

	// To Print all elements
	fmt.Println(myLatestCourses)

	// To create a new slice from an existing slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// [2:5] left index is inclusive and right index is exclusive.
	// [:5] means 0 to 4
	// [2:] means 2 to end of array
	selectedNumbers := numbers[2:5]
	//  Prints [3 4 5]
	fmt.Println(selectedNumbers)

	// append()
	// You can go ahead and keep on appending to the slice until it reaches the capacity of the array.
	// Once it reaches the capacity of the array, Go doubles the array based on the size of the underlying array! start at 4 -> 8 -> 16 -> 32
	mySlice := make([]int, 1, 4)
	fmt.Println("The length of the slice mySlice is ", len(mySlice), " and its capacity is ", cap(mySlice))
	mySlice[0] = 1
	fmt.Println("Capacity is ", cap(mySlice), "Length is ", len(mySlice), " Contents ", mySlice)
	for i := 2; i <= 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Capacity is ", cap(mySlice), "Length is ", len(mySlice), " Contents ", mySlice)
		time.Sleep(1 * time.Millisecond)
	}
	/*
		The length of the slice mySlice is  1  and its capacity is  4
		Capacity is  4 Length is  1  Contents  [1]
		Capacity is  4 Length is  2  Contents  [1 2]
		Capacity is  4 Length is  3  Contents  [1 2 3]
		Capacity is  4 Length is  4  Contents  [1 2 3 4]
		Capacity is  8 Length is  5  Contents  [1 2 3 4 5]
		Capacity is  8 Length is  6  Contents  [1 2 3 4 5 6]
		Capacity is  8 Length is  7  Contents  [1 2 3 4 5 6 7]
		Capacity is  8 Length is  8  Contents  [1 2 3 4 5 6 7 8]
		Capacity is  16 Length is  9  Contents  [1 2 3 4 5 6 7 8 9]
		Capacity is  16 Length is  10  Contents  [1 2 3 4 5 6 7 8 9 10]
		Capacity is  16 Length is  11  Contents  [1 2 3 4 5 6 7 8 9 10 11]
		Capacity is  16 Length is  12  Contents  [1 2 3 4 5 6 7 8 9 10 11 12]
		Capacity is  16 Length is  13  Contents  [1 2 3 4 5 6 7 8 9 10 11 12 13]
		Capacity is  16 Length is  14  Contents  [1 2 3 4 5 6 7 8 9 10 11 12 13 14]
		Capacity is  16 Length is  15  Contents  [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
		Capacity is  16 Length is  16  Contents  [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
		Capacity is  32 Length is  17  Contents  [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17]
	*/

	// Miscellaneous items.
	// Iterate using for range loops.
	for index, item := range mySlice {
		fmt.Println("Item in index ", index, " is ", item)
	}

	// Append a slice to a slice - remember same type!
	newSlice := []int{18, 19, 20, 21, 22}
	mySlice = append(mySlice, newSlice...) //... at the end gives the elements of the slice.
	fmt.Println(mySlice)
}

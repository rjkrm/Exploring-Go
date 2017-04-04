package main

import (
	"fmt"
)

func main () {
	fmt.Println("Exploring Maps in Go!")
	// Similar to arrays and slices.
	// Unordered lists
	// Like hashmaps and dictionaries in other languages
	// Maps are reference, similar to slices, passed by reference.

	// Define maps
    // map[keyType]valueType
	// keyType should be comparable types!
	// Keys got to be unique! You know this!
	scoreMap := make(map[string]int)
    scoreMap["ManchesterUnited"] = 10
	scoreMap["NewCastle"] = 4

	// Can be declared using Composite Literal form.
	recentScoreMap := map[string]int {
		"NewCastle" : 5,
		"ManchesterUnited": 12, // you need the trailing comma, not as in JS or C#.
	}

    // Dump to console.
	fmt.Println(scoreMap)
	fmt.Println(recentScoreMap)
    
	testMap := map[string]int {"A":1,"B":2,"C":3,"D":4,"E":5}
	printMap(testMap)
	fmt.Println("Modified value of B read from map after executing printMap()  ", testMap["B"])

	// Updating and deleting items from a map!
	// Update
	testMap["A"] = 25
	testMap["F"] = 200 // add the key and the value to the map!
    fmt.Println()
	fmt.Println("After insert and update!")
	fmt.Println()	
	printMap(testMap)

	// To delete use the delete(map,key) function.
	delete(testMap,"F")
	fmt.Println()
	fmt.Println("After delete!")
	fmt.Println()
	printMap(testMap)

	// Maps are reference types 
	// Passed by reference 
	// Unsafe for concurrency
	// Cheap to be passed into functions!
	// Specify size for large maps make(map[<keyType>]<valueType>, size) - can improve performance!
}

func printMap (testMap map[string]int) {
	// Maps are passed by reference
	testMap["B"] = 100
	fmt.Println("Modified B inside the printMap() function!")
	for key, value := range testMap {
		// Values are not printed in the same order as they were added.
		// Map doesn't guarantee order.
		fmt.Println("Team ", key, "\tScore ", value)
	}
	return
}
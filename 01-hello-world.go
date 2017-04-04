// Since this is a simple console application, it has package name as main.
package main

// Import the format package to print to console.
import (
	"fmt"
)

// Main method is the main entry point.
// Also, note that Go will export functions in a module to other modules only when they are Pascal Cased!
func main() {
	// Invoke Println to write to console.
	fmt.Println("Hello World")
}

// Package car to create a vehicle
package car

import "fmt"

// engine is a function that returns a bool
func Engine() bool {
	var engineState bool = true 
	fmt.Printf("%v", "Is engine on: ")
	return engineState
}

// --------------------------------------------
// Author: Kirubel Temesgen (C00260396)
// Date: 07/12/2024
// Project: Wa-Tor Simulation
// Description:
// Implementation of the Wa-Tor simulation to demonstrate understanding
// of Go concurrency and threading.
// Issues:
// None
// --------------------------------------------

package main

import "fmt"

// Entity interface represents any entity that can exist on the grid (e.g., Fish, Shark).
type Entity interface {
	Symbol() string // Returns the string representation of the entity (e.g. colored symbol).
}

// Fish struct represents a fish entity with a breeding counter.
type Fish struct {
	BreedCounter int // Tracks the number of steps since the fish last reproduced.
}

// Symbol returns the colored representation of a fish ("F") in green.
func (f *Fish) Symbol() string {
	return fmt.Sprintf("\033[32mF\033[0m")
}

// Shark struct represents a shark entity with a breeding counter and energy level.
type Shark struct {
	BreedCounter int // Tracks the number of steps since the shark last reproduced.
	Energy       int // Tracks the shark's energy level (decreases each step without food).
}

// Symbol returns the colored representation of a shark ("S") in red.
func (s *Shark) Symbol() string {
	return fmt.Sprintf("\033[31mS\033[0m")
}

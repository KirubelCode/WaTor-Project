/**
 * @file main.go
 * @brief Entry point for the Wa-Tor simulation.
 * @details This file contains the main logic for initialising and running the simulation,
 * including setting parameters for grid size, breeding times, and shark starvation energy.
 */
// Next implement use of threads
package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @brief Main function for the Wa-Tor simulation.
 * @details Initialises the simulation grid, sets parameters for fish and shark behaviour,
 * and iteratively simulates movement and interactions over a defined number of steps.
 */
func main() {
	start := time.Now()              ///< Record the start time
	rand.Seed(time.Now().UnixNano()) ///< Ensures random number generators are always random

	gridSize := 100   ///< (15x15) Grid Size
	fishBreed := 3    ///< Fish breed after 3 chronons
	sharkBreed := 4   ///< Sharks breed after 5 chronons
	starveEnergy := 4 ///< Sharks die if they dont eat within 4 chronons

	grid := NewGrid(gridSize)
	grid.Initialise(100, 100) ///< Initialise 20 fish and 10 sharks on the grid

	for step := 0; step < 50; step++ {
		fmt.Printf("Step %d:\n", step)
		grid.Print()
		numFish, numSharks := grid.CountEntities()
		fmt.Printf("Fish: %d, Sharks: %d\n\n", numFish, numSharks)

		grid.MoveEntitiesConcurrent(fishBreed, sharkBreed, starveEnergy) ///< Concurrently update breeding and starvation time
	}

	// Final summary
	fmt.Println("Simulation Ended.")
	numFish, numSharks := grid.CountEntities()
	fmt.Printf("Final Fish: %d, Final Sharks: %d\n", numFish, numSharks) ///< Print amount of sharks and fish
	end := time.Now()                                                    ///< Record the end time
	fmt.Printf("Execution Time: %v\n", end.Sub(start))                   ///< Calculate and print elapsed time
}

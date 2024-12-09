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

/**
 * @file main.go
 * @brief Entry point for the Wa-Tor simulation.
 * @details This file contains the main logic for initialising and running the simulation,
 * including setting parameters for grid size, breeding times, and shark starvation energy.
 */
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

	// Default parameters
	numShark := 100   ///< Initial number of sharks
	numFish := 100    ///< Initial number of fish
	fishBreed := 3    ///< Fish breed after 3 chronons
	sharkBreed := 3   ///< Sharks breed after 3 chronons
	starveEnergy := 4 ///< Sharks die if they don’t eat within 4 chronons
	gridSize := 100   ///< Grid size (50x50 by default)
	threads := 10     ///< Default number of threads for concurrency

	// Check if command-line arguments are provided
	if len(os.Args) == 8 {
		numShark, _ = strconv.Atoi(os.Args[1])
		numFish, _ = strconv.Atoi(os.Args[2])
		fishBreed, _ = strconv.Atoi(os.Args[3])
		sharkBreed, _ = strconv.Atoi(os.Args[4])
		starveEnergy, _ = strconv.Atoi(os.Args[5])
		gridSize, _ = strconv.Atoi(os.Args[6])
		threads, _ = strconv.Atoi(os.Args[7])
	} else if len(os.Args) != 1 { // Print usage only if arguments are partially supplied
		fmt.Println("Usage: go run main.go <NumShark> <NumFish> <FishBreed> <SharkBreed> <Starve> <GridSize> <Threads>")
		return
	}

	grid := NewGrid(gridSize)
	grid.Initialize(numFish, numShark) ///< Initialise the grid with sharks and fish

	// Simulation loop
	for step := 0; step < 50; step++ {
		fmt.Printf("Step %d:\n", step)
		grid.Print()                                               ///< Print the current state of the grid
		numFish, numSharks := grid.CountEntities()                 ///< Count the number of fish and sharks
		fmt.Printf("Fish: %d, Sharks: %d\n\n", numFish, numSharks) ///< Print the counts

		grid.MoveEntitiesWithThreads(fishBreed, sharkBreed, starveEnergy, threads) ///< Concurrently update grid state using threads
	}

	// Final summary
	fmt.Println("Simulation Ended.")
	numFish, numSharks := grid.CountEntities()
	fmt.Printf("Final Fish: %d, Final Sharks: %d\n", numFish, numSharks) ///< Print final counts

	end := time.Now()                                  ///< Record the end time
	fmt.Printf("Execution Time: %v\n", end.Sub(start)) ///< Calculate and print elapsed time
}

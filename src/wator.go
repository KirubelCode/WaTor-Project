/*!
 * @file main.go
 * @brief Wa-Tor simulation main file.
 * @author Kirubel Temesgen
 * @date 2024-12-03
 *
 * This file contains the main logic for initialising the simulation grid,
 * including concurrent population of fish and sharks, and visualising the grid.
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*!
 * @brief Size of the simulation grid.
 */
const GridSize = 20

/*!
 * @struct Fish
 * @brief Represents a fish in the simulation.
 *
 * @var Fish::Age
 * Age of the fish in chronons.
 *
 * @var Fish::BreedAge
 * Number of chronons required before the fish can reproduce.
 */
type Fish struct {
	Age      int
	BreedAge int
}

/*!
 * @struct Shark
 * @brief Represents a shark in the simulation.
 *
 * @var Shark::Age
 * Age of the shark in chronons.
 *
 * @var Shark::BreedAge
 * Number of chronons required before the shark can reproduce.
 *
 * @var Shark::Energy
 * Energy level of the shark, decreases each chronon.
 */
type Shark struct {
	Age      int
	BreedAge int
	Energy   int
}

/*!
 * @struct Cell
 * @brief Represents a single cell in the grid.
 *
 * @var Cell::Fish
 * Pointer to a Fish struct if a fish occupies this cell, or nil otherwise.
 *
 * @var Cell::Shark
 * Pointer to a Shark struct if a shark occupies this cell, or nil otherwise.
 */
type Cell struct {
	Fish  *Fish
	Shark *Shark
}

/*!
 * @typedef Grid
 * @brief A 2D slice representing the simulation grid.
 */
type Grid [][]Cell

/*!
 * @brief Initializes the simulation grid with fish and sharks concurrently.
 *
 * @param size The size of the grid (NxN).
 * @param numFish Number of fish to place in the grid.
 * @param numShark Number of sharks to place in the grid.
 * @return A 2D grid populated with fish and sharks.
 */
func initializeGridConcurrently(size, numFish, numShark int) Grid {
	grid := make(Grid, size) // Create the grid
	for i := range grid {
		grid[i] = make([]Cell, size) // Initialise each row of the grid
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Create a thread-safe random generator
	var wg sync.WaitGroup                                // WaitGroup to synchronise goroutines

	// Goroutine to populate the grid with fish
	wg.Add(1)
	go func() {
		defer wg.Done() // Mark this goroutine as done when finished
		for i := 0; i < numFish; i++ {
			for {
				// Randomly choose a cell in the grid
				x, y := r.Intn(size), r.Intn(size)
				// Check if the cell is empty (no fish or shark present)
				if grid[x][y].Fish == nil && grid[x][y].Shark == nil {
					grid[x][y].Fish = &Fish{Age: 0, BreedAge: 3} // Add a fish
					break                                        // Move on to the next fish
				}
			}
		}
	}()

	// Goroutine to populate the grid with sharks
	wg.Add(1)
	go func() {
		defer wg.Done() // Mark this goroutine as done when finished
		for i := 0; i < numShark; i++ {
			for {
				// Randomly choose a cell in the grid
				x, y := r.Intn(size), r.Intn(size)
				// Check if the cell is empty (no fish or shark present)
				if grid[x][y].Fish == nil && grid[x][y].Shark == nil {
					grid[x][y].Shark = &Shark{Age: 0, BreedAge: 5, Energy: 5} // Add a shark
					break                                                     // Move on to the next shark
				}
			}
		}
	}()

	wg.Wait() // Wait for all goroutines to finish
	return grid
}

/*!
 * @brief Prints the current state of the grid.
 *
 * This function visualizes the grid using the following symbols:
 * - 'F' for fish
 * - 'S' for sharks
 * - '.' for empty cells
 *
 * @param grid The simulation grid to print.
 */
func printGrid(grid Grid) {
	for _, row := range grid {
		for _, cell := range row {
			if cell.Fish != nil {
				fmt.Print("F ") // Display "F" for fish
			} else if cell.Shark != nil {
				fmt.Print("S ") // Display "S" for sharks
			} else {
				fmt.Print(". ") // Display "." for empty cells
			}
		}
		fmt.Println() // Move to the next row
	}
}

/*!
 * @brief Moves a fish to a random adjacent unoccupied cell.
 *
 * @param grid The simulation grid.
 * @param x The x-coordinate of the fish.
 * @param y The y-coordinate of the fish.
 * @return A boolean indicating whether the fish moved.
 */
func moveFish(grid Grid, x, y int) bool {
	size := len(grid)
	// Define possible directions: (dx, dy)
	directions := []struct {
		dx, dy int
	}{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // North, South, West, East
	}

	// Collect all valid adjacent unoccupied cells
	var validMoves []struct{ nx, ny int }
	for _, d := range directions {
		nx, ny := (x+d.dx+size)%size, (y+d.dy+size)%size           // Wrap around for toroidal grid
		if grid[nx][ny].Fish == nil && grid[nx][ny].Shark == nil { // Check if cell is unoccupied
			validMoves = append(validMoves, struct{ nx, ny int }{nx, ny})
		}
	}

	// If no valid moves, return false
	if len(validMoves) == 0 {
		return false
	}

	// Randomly select a valid move
	move := validMoves[rand.Intn(len(validMoves))]

	// Move the fish
	grid[move.nx][move.ny].Fish = grid[x][y].Fish // Place fish in the new cell
	grid[x][y].Fish = nil                         // Clear the old cell

	return true
}

/*!
 * @brief Updates the state of all fish on the grid by moving them.
 *
 * @param grid The simulation grid.
 */
func updateFish(grid Grid) {
	size := len(grid)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if grid[x][y].Fish != nil {
				moveFish(grid, x, y)
			}
		}
	}
}

func main() {
	// Initialise the grid with 10 fish and 5 sharks
	grid := initializeGridConcurrently(GridSize, 10, 5)

	// Display the initial state of the grid
	fmt.Println("Initial State:")
	printGrid(grid)

	// Update fish movement
	fmt.Println("\nAfter Fish Movement:")
	updateFish(grid)
	printGrid(grid)
}

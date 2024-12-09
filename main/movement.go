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
 * @file movement.go
 * @brief Handles movement and interactions of fish and sharks on the grid.
 * @details Implements concurrent movement using threads and WaitGroups for grid sections,
 * ensuring synchronization while processing fish and sharks in parallel.
 */
package main

import (
	"math/rand"
	"sync"
)

/**
 * @brief Moves fish and sharks concurrently in the grid using threads.
 * @details Divides the grid into sections handled by separate threads for parallel processing.
 * @param fishBreed Number of chronons before fish can reproduce.
 * @param sharkBreed Number of chronons before sharks can reproduce.
 * @param starveEnergy Maximum energy level before sharks die of starvation.
 * @param threads Number of threads to use for concurrent processing.
 */
func (g *Grid) MoveEntitiesWithThreads(fishBreed, sharkBreed, starveEnergy, threads int) {
	newGrid := NewGrid(g.Size) ///< Create a new grid for updated positions

	rowsPerThread := g.Size / threads ///< Divide rows among threads
	var wg sync.WaitGroup             ///< WaitGroup to synchronise goroutines

	// Launch threads to process sections of the grid
	for i := 0; i < threads; i++ {
		startRow := i * rowsPerThread
		endRow := startRow + rowsPerThread
		if i == threads-1 {
			endRow = g.Size // Ensure the last thread handles all remaining rows
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			g.processSection(newGrid, start, end, fishBreed, sharkBreed, starveEnergy)
		}(startRow, endRow)
	}

	wg.Wait()               ///< Block until all threads complete
	g.Cells = newGrid.Cells ///< Update the main grid with the new positions
}

/**
 * @brief Processes a section of the grid for movement and interactions.
 * @details Handles fish and shark movement in a specific section of the grid.
 * @param newGrid The new grid for updated positions.
 * @param startRow The starting row for this section.
 * @param endRow The ending row for this section.
 * @param fishBreed Number of chronons before fish can reproduce.
 * @param sharkBreed Number of chronons before sharks can reproduce.
 * @param starveEnergy Maximum energy level before sharks die of starvation.
 */
func (g *Grid) processSection(newGrid *Grid, startRow, endRow, fishBreed, sharkBreed, starveEnergy int) {
	for x := startRow; x < endRow; x++ {
		for y := 0; y < g.Size; y++ {
			if fish, ok := g.Cells[x][y].(*Fish); ok {
				g.processFish(newGrid, fish, x, y, fishBreed)
			} else if shark, ok := g.Cells[x][y].(*Shark); ok {
				g.processShark(newGrid, shark, x, y, sharkBreed, starveEnergy)
			}
		}
	}
}

/**
 * @brief Handles movement and reproduction of fish.
 * @details Updates fish position and reproduces based on breeding counter.
 * @param newGrid The new grid for updated positions.
 * @param fish The fish entity to process.
 * @param x The current x-coordinate of the fish.
 * @param y The current y-coordinate of the fish.
 * @param fishBreed Number of chronons before fish can reproduce.
 */
func (g *Grid) processFish(newGrid *Grid, fish *Fish, x, y, fishBreed int) {
	newX, newY := g.findEmptyAdjacent(x, y)
	if newX != -1 && newY != -1 {
		newGrid.Cells[newX][newY] = fish ///< Move fish to the new position
	} else {
		newGrid.Cells[x][y] = fish ///< Fish stays in its current position
	}
	fish.BreedCounter++
	if fish.BreedCounter >= fishBreed {
		newGrid.Cells[x][y] = &Fish{} ///< Leave a new fish in the current position
		fish.BreedCounter = 0         ///< Reset breeding counter
	}
}

/**
 * @brief Handles movement, reproduction, and starvation of sharks.
 * @details Sharks move to eat fish or to adjacent empty cells and handle reproduction and energy depletion.
 * @param newGrid The new grid for updated positions.
 * @param shark The shark entity to process.
 * @param x The current x-coordinate of the shark.
 * @param y The current y-coordinate of the shark.
 * @param sharkBreed Number of chronons before sharks can reproduce.
 * @param starveEnergy Maximum energy level before sharks die of starvation.
 */
func (g *Grid) processShark(newGrid *Grid, shark *Shark, x, y, sharkBreed, starveEnergy int) {
	shark.Energy-- ///< Sharks lose energy each step
	if shark.Energy <= 0 {
		return ///< Shark dies if energy reaches 0
	}

	newX, newY := g.findNearestFish(x, y)
	if newX != -1 && newY != -1 {
		newGrid.Cells[newX][newY] = shark ///< Move shark to eat fish
		shark.Energy = starveEnergy       ///< Reset energy after eating
	} else {
		newX, newY = g.findEmptyAdjacent(x, y)
		if newX != -1 && newY != -1 {
			newGrid.Cells[newX][newY] = shark ///< Move shark to an empty cell
		} else {
			newGrid.Cells[x][y] = shark ///< Shark stays in its current position
		}
	}

	shark.BreedCounter++
	if shark.BreedCounter >= sharkBreed {
		newGrid.Cells[x][y] = &Shark{Energy: starveEnergy} ///< Reproduce a new shark
		shark.BreedCounter = 0                             ///< Reset breeding counter
	}
}

/**
 * @brief Finds an adjacent empty cell for movement.
 * @details Searches the four directions (North, South, West, East) for empty cells.
 * @param x The x-coordinate of the current cell.
 * @param y The y-coordinate of the current cell.
 * @return Coordinates of an empty cell, or (-1, -1) if none are available.
 */
func (g *Grid) findEmptyAdjacent(x, y int) (int, int) {
	directions := []struct{ dx, dy int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // North, South, West, East
	}
	rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] }) // Randomise directions

	for _, dir := range directions {
		newX := (x + dir.dx + g.Size) % g.Size
		newY := (y + dir.dy + g.Size) % g.Size
		if g.Cells[newX][newY] == nil {
			return newX, newY
		}
	}
	return -1, -1 ///< No empty adjacent cells found
}

/**
 * @brief Finds the nearest adjacent fish for a shark to eat.
 * @details Searches the four cardinal directions for fish.
 * @param x The x-coordinate of the current cell.
 * @param y The y-coordinate of the current cell.
 * @return Coordinates of the nearest fish, or (-1, -1) if none are found.
 */
func (g *Grid) findNearestFish(x, y int) (int, int) {
	directions := []struct{ dx, dy int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // North, South, West, East
	}
	rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] }) // Randomise directions

	for _, dir := range directions {
		newX := (x + dir.dx + g.Size) % g.Size        ///< Wrap around toroidal grid horizontally
		newY := (y + dir.dy + g.Size) % g.Size        ///< Wrap around toroidal grid vertically
		if _, ok := g.Cells[newX][newY].(*Fish); ok { ///< Check if the cell contains a fish
			return newX, newY
		}
	}
	return -1, -1 ///< No fish found in adjacent cells
}

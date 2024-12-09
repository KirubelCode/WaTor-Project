/**
 * @file movement.go
 * @brief Handles movement and interactions of fish and sharks on the grid.
 * @details Implements concurrent movement using goroutines, WaitGroups, and channels to process shark and fish interactions.
 */

package main

import (
	"math/rand"
	"sync"
)

/**
 * @brief Moves fish and sharks concurrently in the grid.
 * @details Fish are moved first, followed by sharks, whose movements are managed concurrently using goroutines.
 * @param fishBreed Number of chronons before fish can reproduce.
 * @param sharkBreed Number of chronons before sharks can reproduce.
 * @param starveEnergy Maximum energy level before sharks die of starvation.
 */
func (g *Grid) MoveEntitiesConcurrent(fishBreed, sharkBreed, starveEnergy int) {
	newGrid := NewGrid(g.Size)

	// Move fish
	g.moveFish(newGrid, fishBreed)

	// Move sharks concurrently
	var wg sync.WaitGroup                         ///< WaitGroup to manage concurrent goroutines
	moveChannel := make(chan Move, g.Size*g.Size) ///< Buffered channel to collect shark moves safely

	wg.Add(1) // Add one goroutine to the WaitGroup
	go func() {
		defer wg.Done() // Mark goroutine as done when it finishes
		g.moveSharksConcurrent(newGrid, sharkBreed, starveEnergy, moveChannel)
	}()

	// Wait for all goroutines to finish
	wg.Wait()          ///< Block until all goroutines in the WaitGroup are done
	close(moveChannel) ///< Close the channel to signal no more moves will be sent

	// Apply all moves to the new grid
	for move := range moveChannel { // Retrieve moves from the channel
		newGrid.Cells[move.X][move.Y] = move.Entity
	}

	// Update the main grid
	g.Cells = newGrid.Cells
}

/**
 * @struct Move
 * @brief Represents a movement or action of an entity.
 */
type Move struct {
	X, Y   int    ///< Coordinates where the entity will move.
	Entity Entity ///< The entity being moved (fish or shark).
}

/**
 * @brief Concurrently moves sharks on the grid.
 * @details Sharks prioritise hunting fish and handle starvation and breeding conditions. All movements are sent through a channel.
 * @param newGrid The new grid for updated positions.
 * @param sharkBreed Number of chronons before sharks can reproduce.
 * @param starveEnergy Maximum energy level before sharks die of starvation.
 * @param moveChannel Channel to send shark movements safely across goroutines.
 */
func (g *Grid) moveSharksConcurrent(newGrid *Grid, sharkBreed, starveEnergy int, moveChannel chan<- Move) {
	for x := 0; x < g.Size; x++ {
		for y := 0; y < g.Size; y++ {
			if shark, ok := g.Cells[x][y].(*Shark); ok {
				// Sharks lose energy every step
				shark.Energy--

				// If energy is 0 or less, shark dies
				if shark.Energy <= 0 {
					continue
				}

				// Prioritise hunting fish
				newX, newY := g.findNearestFish(x, y)
				if newX != -1 && newY != -1 {
					// Eat fish and reset energy
					moveChannel <- Move{X: newX, Y: newY, Entity: shark} ///< Send move to the channel
					shark.Energy = starveEnergy
				} else {
					// No fish nearby, move to an empty cell
					newX, newY = g.findEmptyAdjacent(x, y)
					if newX != -1 && newY != -1 {
						moveChannel <- Move{X: newX, Y: newY, Entity: shark} ///< Send move to the channel
					} else {
						// Stay in place if no empty space is available
						moveChannel <- Move{X: x, Y: y, Entity: shark} ///< Send move to the channel
					}
				}

				// Increment breeding counter
				shark.BreedCounter++
				if shark.BreedCounter >= sharkBreed {
					// Shark reproduces
					moveChannel <- Move{X: x, Y: y, Entity: &Shark{Energy: starveEnergy}} ///< Send new shark to the channel
					shark.BreedCounter = 0
				}
			}
		}
	}
}

/**
 * @brief Moves fish on the grid.
 * @details Fish move to adjacent empty cells and handle reproduction based on their breed counter.
 * @param newGrid The new grid for updated positions.
 * @param fishBreed Number of chronons before fish can reproduce.
 */
func (g *Grid) moveFish(newGrid *Grid, fishBreed int) {
	for x := 0; x < g.Size; x++ {
		for y := 0; y < g.Size; y++ {
			if fish, ok := g.Cells[x][y].(*Fish); ok {
				// Find an empty adjacent cell
				newX, newY := g.findEmptyAdjacent(x, y)
				if newX != -1 && newY != -1 {
					// Move fish to the new position
					newGrid.Cells[newX][newY] = fish
				} else {
					// Fish stays in its current position
					newGrid.Cells[x][y] = fish
				}
				fish.BreedCounter++
				if fish.BreedCounter >= fishBreed { // Check if fish can reproduce
					newGrid.Cells[x][y] = &Fish{} // Leave a new fish in the current position
					fish.BreedCounter = 0         // Reset breeding counter
				}
			}
		}
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
	// Randomise direction order
	rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })
	// Wrap around toroidal grid
	for _, dir := range directions {
		newX := (x + dir.dx + g.Size) % g.Size
		newY := (y + dir.dy + g.Size) % g.Size
		if g.Cells[newX][newY] == nil { // Check if the cell is empty
			return newX, newY
		}
	}
	return -1, -1 // No empty adjacent cells found
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
		// Wrap around toroidal grid
		newX := (x + dir.dx + g.Size) % g.Size
		newY := (y + dir.dy + g.Size) % g.Size
		if _, ok := g.Cells[newX][newY].(*Fish); ok { // Check if the cell contains a fish
			return newX, newY
		}
	}
	return -1, -1 // No fish found in adjacent cells
}

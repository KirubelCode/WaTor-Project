/**
 * @file grid.go
 * @brief Defines the simulation grid where entities (fish and sharks) interact.
 * @details Implements grid creation, entity placement, counting, and visualisation.
 */

package main

import (
	"fmt"
	"math/rand"
)

/**
 * @struct Grid
 * @brief Represents the simulation grid.
 * @details The grid holds all entities (fish and sharks) and tracks their positions.
 */
type Grid struct {
	Size  int        ///< Dimensions of the grid
	Cells [][]Entity ///< Holds entities at each grid position
}

/**
 * @brief Creates a new Grid of the specified size with empty cells.
 * @param size The dimensions of the grid (size x size).
 * @return A pointer to the newly created Grid.
 */
func NewGrid(size int) *Grid {
	cells := make([][]Entity, size)
	for i := range cells {
		cells[i] = make([]Entity, size)
	}
	return &Grid{Size: size, Cells: cells}
}

/**
 * @brief Initialises and populates the grid with a specified number of fish and sharks.
 * @param numFish The number of fish to add to the grid.
 * @param numSharks The number of sharks to add to the grid.
 */
func (g *Grid) Initialise(numFish, numSharks int) {
	for i := 0; i < numFish; i++ {
		g.addEntity(&Fish{}) ///< Add fish to random positions
	}
	for i := 0; i < numSharks; i++ {
		g.addEntity(&Shark{Energy: 4}) ///< Add sharks with initial energy to random positions
	}
}

/**
 * @brief Places a fish or shark in a random unoccupied cell on the grid.
 * @param e The entity (fish or shark) to place on the grid.
 */
func (g *Grid) addEntity(e Entity) {
	for {
		x, y := rand.Intn(g.Size), rand.Intn(g.Size) ///< Randomly select grid position
		if g.Cells[x][y] == nil {                    ///< Place entity only if cell is empty
			g.Cells[x][y] = e
			break
		}
	}
}

/**
 * @brief Counts the number of fish and sharks currently on the grid.
 * @return A tuple (numFish, numSharks) representing the counts of each entity type.
 */
func (g *Grid) CountEntities() (numFish, numSharks int) {
	for x := 0; x < g.Size; x++ {
		for y := 0; y < g.Size; y++ {
			if _, ok := g.Cells[x][y].(*Fish); ok {
				numFish++ ///< Increment fish count
			}
			if _, ok := g.Cells[x][y].(*Shark); ok {
				numSharks++ ///< Increment shark count
			}
		}
	}
	return
}

/**
 * @brief Displays the current state of the grid with borders for clarity.
 */
func (g *Grid) Print() {
	fmt.Println("+---------------------+")
	for _, row := range g.Cells {
		fmt.Print("| ")
		for _, cell := range row {
			if cell == nil {
				fmt.Print(". ") ///< Print "." for empty cells
			} else {
				fmt.Print(cell.Symbol(), " ") ///< Print the symbol of the entity in the cell
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+---------------------+")
}

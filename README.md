Wa-Tor Simulation – Multi-Threaded Predator-Prey Model
This project implements the Wa-Tor predator-prey simulation using Go and demonstrates a strong understanding of concurrency. The simulation models a toroidal ecosystem where fish and sharks move, breed, and survive based on simple rules. It uses multi-threading with Go's goroutines and synchronisation tools to efficiently handle updates across the grid in parallel, speeding up execution on multi-core systems.

- Inspired by: https://en.wikipedia.org/wiki/Wa-Tor
------

Overview
The Wa-Tor simulation models a predator-prey system where fish move and breed, while sharks hunt fish and die if they starve. The world is a toroidal grid, meaning creatures that move off one edge reappear on the other. The key focus of this project is concurrency: the simulation updates parts of the grid concurrently using multiple threads, making it faster and more efficient.

------

Key Features
- Concurrency: Parallelizes the movement and actions of fish and sharks using Go’s goroutines.

- Customisable Parameters: Configure the grid size, number of fish/sharks, breeding intervals, and number of threads.

- Toroidal World: Grid wraps around, ensuring continuous movement.

- Performance Benchmarking: Logs execution time for different configurations to evaluate the impact of concurrency.

-----

Installation & Usage
- Prerequisites: Install Go (latest version recommended) on your system.

Clone the Repository:
- git clone https://github.com/KirubelCode/WaTor-Project.git
- cd WaTor-Project

Run the Simulation:
- To run with default settings (100 fish, 100 sharks, 100x100 grid, 10 threads): go run main.go

- To customise parameters (e.g., 100 fish, 100 sharks, 8 threads):
go run main.go 100 100 3 3 4 100 8

Parameters:
- NumFish: Number of fish

- NumSharks: Number of sharks

- FishBreed: Fish breeding interval

- SharkBreed: Shark breeding interval

- Starve: Starvation time for sharks

- GridSize: Grid dimensions (N x N)

- Threads: Number of threads to use for concurrency

-----

Technologies Used
Go (Golang): Core logic and concurrency management.

Standard Library: For synchronisation (e.g., sync.WaitGroup).

-----

Future Improvements

- Extended Features: Add new creatures or behaviors to increase complexity.

- Code Optimisation: Refactor code for better readability and efficiency.

- Testing: Add a comprehensive suite of tests, including concurrency-specific ones.

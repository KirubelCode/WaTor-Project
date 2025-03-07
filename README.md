Wa-Tor Simulation – Multi-Threaded Predator-Prey Model
Produced by: Kirubel Temesgen
College ID: C00260396
Date: 07/12/2024

Overview
This project implements the Wa-Tor simulation, a classic predator-prey model, using Go concurrency and threading. The simulation models interactions between fish and sharks on a toroidal grid, where fish breed over time and sharks hunt for fish to survive.

This project demonstrates:

Efficient parallel processing using Go routines and WaitGroups.
Dynamic workload distribution across multiple threads.
Optimised synchronisation techniques to handle concurrent grid updates.
Features

Multi-Threaded Execution – Utilises concurrent processing for entity movement.
Customisable Simulation Parameters – Users can specify grid size, entity count, breeding time, and starvation thresholds.
Toroidal Grid Representation – Allows seamless movement across edges.
Performance Benchmarking – Measures execution time for different thread counts.

Installation

1. Prerequisites
Go (Golang) installed (latest version recommended).
Command-line access (Windows, macOS, or Linux).

2. Clone the Repository
Run the following commands:
git clone https://github.com/KirubelCode/WaTor-Project.git

3. Compile and Run the Simulation
To run with default settings:
go run main.go

To specify parameters:
go run main.go <NumShark> <NumFish> <FishBreed> <SharkBreed> <Starve> <GridSize> <Threads>

Example:
go run main.go 100 100 3 3 4 100 8
(Runs the simulation with 100 fish, 100 sharks, 100x100 grid, and 8 threads).

Simulation Parameters
Parameter	Description	Default Value
NumShark	Number of sharks	100
NumFish	Number of fish	100
FishBreed	Breeding time for fish (chronons)	3
SharkBreed	Breeding time for sharks (chronons)	3
Starve	Starvation threshold for sharks	4
GridSize	Grid dimensions (NxN)	100
Threads	Number of concurrent threads	10

How It Works
Grid Initialisation – The simulation starts with a random distribution of fish and sharks.
Parallel Entity Movement – Fish and sharks move concurrently, processed by separate threads.
Shark Behaviour – Sharks hunt for fish, consume energy, and starve if no food is available.
Fish Reproduction – Fish breed at a predefined rate if space is available.
Thread Synchronisation – The grid is divided into sections, and each thread updates a portion before synchronising with the main grid.
Performance Logging – Execution time is measured to analyse concurrency efficiency.

Performance Observations
Grid Size	Fish	Sharks	Steps	Threads	Execution Time
50x50	50	50	50	1	6.337s
50x50	50	50	50	2	6.668s
100x100	100	100	50	4	1.80s
100x100	100	100	50	8	2.725s
100x100	100	100	50	10	2.01s

Key Findings
Low thread counts (1-2) offer minimal speedup due to thread creation overhead.
4 threads show optimal performance on a dual-core processor (maximising concurrency).
Higher thread counts (8-10) introduce synchronisation delays, limiting performance gains.

File Structure
Wa-Tor-Simulation
│── main.go # Simulation entry point
│── movement.go # Concurrency-based entity movement logic
│── grid.go # Grid structure and initialisation
│── shark.go # Shark behaviour implementation
│── fish.go # Fish behaviour implementation
│── README.md # Project documentation

Future Enhancements
Dynamic Load Balancing – Adaptive thread allocation based on real-time workload.
Graphical Visualisation – Adding a GUI for real-time simulation playback.
Machine Learning Integration – Training models to predict ecosystem changes.

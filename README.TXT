Wa-Tor Simulation Project

1. Overview

The Wa-Tor simulation models sharks and fish on a toroidal grid. Implemented in Go, it demonstrates concurrency and threading for simulation.

2. Project Structure

NonThreadedVersion/: Basic implementation without threading. or --> use Old-State branch wa-tor file in parent directory.

Main Branch
main/: Implementation with threading.

3. Parameters

The simulation accepts the following parameters:

NumShark: Starting shark population.

NumFish: Starting fish population.

FishBreed: Chronons for fish reproduction.

SharkBreed: Chronons for shark reproduction.

Starve: Time before sharks die without food.

GridSize: Dimensions of the grid.

Threads: Number of threads for concurrency (threaded version only).

Can be added via terminal i.e:

Or

Within the code in the top of main.go

4. Usage

Non-Threaded Version:

Non-threaded Version - /wa-tor/main (Old-State branch)
go run main.go <NumShark> <NumFish> <FishBreed> <SharkBreed> <Starve> <GridSize>

Threaded Version:

Threaded Version - main/ (main branch)
go run main.go <NumShark> <NumFish> <FishBreed> <SharkBreed> <Starve> <GridSize> <Threads>

5. Results

Speedup measurements with 1, 2, 4, and 8 threads.

Results are documented in Wa-tor Speedup Results/ with tables.

6. Requirements

Go Language: Installed via sudo apt install golang.

Linux Environment: Tested using WSL and Ubuntu.

7. Author

Kirubel Temesgen

8. License

This project is licensed under the MIT License.
Refrences found in Wa-tor Speedup Results/


# Percolation Problem - Practice Project

## Description

This is a practice project implementing the Percolation problem using the path compression Weighted-Union-Find data structure in Go. The Percolation problem is a classic problem from computational physics that models a system of randomly blocked sites in a grid.

The project includes a `Percolation` interface and a `UnionFind` interface, both implemented in Go. The `Percolation` interface includes methods for opening a site, checking if a site is open or full, and checking if the system percolates. The `UnionFind` interface includes methods for union operation, checking if two sites are connected, and finding the root of a site.

## Usage

```bash
go run main.go path_to_input_file
```

The input file should be a `.txt` file where the first line is a valid integer representing the grid size, and every line after the grid length must include two integers separated by a whitespace representing the row and column of a site to be opened.
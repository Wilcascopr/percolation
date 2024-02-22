package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"percolation/percolation"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("a file path should be specified")
		return
	}
	fileTest := args[1]
	fileExt := filepath.Ext(fileTest)
	if fileExt != ".txt" {
		fmt.Println("input file must be txt")
		return
	}
	file, err := os.Open(fileTest)
	if err != nil {
			fmt.Printf("failed to open file: %v", err)
			return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	sizeStr := scanner.Text()
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		fmt.Println("first line of file must be a valid integer")
		return
	}
	pr, err := percolation.NewPercolation(size)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)
			if len(parts) < 2 {
				fmt.Println("every line after the grid length must include two integers separated by a whitesapce")
				return
			}
			row, err := strconv.Atoi(parts[0])
			col, err2 := strconv.Atoi(parts[1])
			if err != nil || err2 != nil {
				fmt.Println("the two elements in the line must be valid integers")
				return
			}
			pr.Open(row, col)
	}
	pr.PrintGrid()
}

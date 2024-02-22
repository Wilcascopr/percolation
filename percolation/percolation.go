package percolation

import (
	"fmt"
	"percolation/unionfind"
)

type Percolation interface {
	validate(p int) error
	Open(row int, col int) error
	IsOpen(row int, col int) (bool, error)
}

type percolation struct {
	uf unionfind.UnionFind
	full unionfind.UnionFind
	size int
	row_length int
	sites []bool
	open_sites int
	start int
	end int
}

func NewPercolation(n int) (*percolation, error) {
	if n < 1 {
		return nil, fmt.Errorf("n must be an integer greater than 0")
	}
	pr := &percolation{}
	pr.size = n * n + 2
	pr.start = 0
	pr.end = n*n + 1
	pr.row_length = n
	pr.sites = make([]bool, n*n)
	pr.open_sites = 0
	pr.uf = unionfind.NewUnionFind(pr.size)
	pr.full = unionfind.NewUnionFind(pr.size - 1)
	for i := pr.start + 1; i <= n; i++ {
		pr.uf.Union(pr.start, i)
		pr.full.Union(pr.start, i)
		pr.uf.Union(pr.end, pr.end - i)
	}
	return pr, nil
}

func (pr *percolation) flat(row int, col int) int {
	return (row - 1) * pr.row_length + col
}

func (pr *percolation) validateFlat(p int) error {
	if p > pr.size - 2 || p < 1 {
		return fmt.Errorf("point out of bounds")
	}
	return nil
}

func (pr *percolation) validate(row int, col int) error {
	if row < 1 || row > pr.row_length || col < 1 || col > pr.row_length {
		return fmt.Errorf("point out of bounds")
	}
	return nil
} 

func (pr *percolation) unionNeighbour(p int, n int, checkRow bool) {
	if err := pr.validateFlat(n); err != nil || !pr.sites[n - 1]  {
		return
	}
	if checkRow {
		if (n - 1) / pr.row_length != (p - 1) / pr.row_length {
			return
		}
	}
	pr.uf.Union(p, n)
	pr.full.Union(p, n)
} 


func (pr *percolation) Open(row int, col int) error {
	open, err := pr.IsOpen(row, col)
	if err != nil {
		return err
	}
	if open {
		return nil
	}
	p := pr.flat(row, col)
	pr.sites[p - 1] = true
	pr.open_sites++
	pr.neighboring(p)
	return nil
}

func (pr *percolation) IsOpen(row int, col int) (bool, error) {
	if err := pr.validate(row, col); err != nil {
		return false, err
	}
	p := pr.flat(row, col)
	return pr.sites[p - 1], nil
}

func (pr *percolation) neighboring(p int) {
	pr.unionNeighbour(p, p - pr.row_length, false)
	pr.unionNeighbour(p, p + pr.row_length, false)
	pr.unionNeighbour(p, p - 1, true)
	pr.unionNeighbour(p, p + 1, true)
}

func (pr *percolation) IsFull(row int, col int) (bool, error) {
	if open, err := pr.IsOpen(row, col); err != nil {
		return false, err
	} else if !open {
		return open, nil
	}
	p := pr.flat(row, col)
	return pr.full.Connected(pr.start, p)
}

func (pr *percolation) NumberOfOpenSites() int {
	return pr.open_sites
}

func (pr *percolation) Percolates() bool {
	res, _ := pr.uf.Connected(pr.start, pr.end)
	return res
}

func (pr *percolation) PrintGrid() {
	for i, v := range pr.sites {
		if v {
			fmt.Print("\x1b[37m■\x1b[0m")
		} else {
			fmt.Print("\x1b[31m■\x1b[0m")
		}
		if (i + 1) % pr.row_length == 0 {
			fmt.Println()
		}
	}
}
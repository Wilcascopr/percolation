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
	size int
	row_length int
	sites []bool
	open_sites int
	start int
	end int
}

func NewPercolation(n int) *percolation {
	pr := &percolation{}
	pr.size = n * n + 2
	pr.start = 0
	pr.end = n*n + 1
	pr.row_length = n
	pr.sites = make([]bool, n*n)
	pr.open_sites = 0
	pr.uf = unionfind.NewUnionFind(pr.size)
	for i := pr.start + 1; i <= n; i++ {
		pr.uf.Union(pr.start, i)
		pr.uf.Union(pr.end, pr.end - i)
	}
	return pr
}

func (pr *percolation) flat(row int, col int) int {
	return (row - 1) * pr.row_length + col
}

func (pr *percolation) validateFlat(p int) error {
	if p > pr.size - 1 || p < 1 {
		return fmt.Errorf("point out of bounds")
	}
	return nil
}

func (pr *percolation) validate(row int, col int) error {
	if row < 1 || row >= pr.row_length || col < 1 || col >= pr.row_length {
		return fmt.Errorf("point out of bounds")
	}
	return nil
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
	pr.sites[p - 2] = true
	pr.open_sites++
	pr.neighboring(p)
	return nil
}

func (pr *percolation) IsOpen(row int, col int) (bool, error) {
	if err := pr.validate(row, col); err != nil {
		return false, err
	}
	p := pr.flat(row, col)
	return pr.sites[p - 2], nil
}

func (pr *percolation) neighboring(p int) {
	neighbours := [4]int{ p - pr.row_length, p + pr.row_length, p - 1, p + 1}
	for _, n := range neighbours {
		if err := pr.validateFlat(n); err != nil {
			continue
		}
		pr.uf.Union(p, n)
	}
}

func (pr *percolation) IsFull(row int, col int) (bool, error) {
	if open, err := pr.IsOpen(row, col); err != nil {
		return false, err
	} else if !open {
		return open, nil
	}
	p := pr.flat(row, col)
	return pr.uf.Connected(pr.start, p)
}

func (pr *percolation) NumberOfOpenSites() int {
	return pr.open_sites
}

func (pr *percolation) Percolates() bool {
	res, _ := pr.uf.Connected(pr.start, pr.end)
	return res
}
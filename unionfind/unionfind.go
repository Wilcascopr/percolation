package unionfind

import "fmt"

type UnionFind interface {
	Union(p int, q int)
	Connected(p int, q int) bool
	Find(p int) int
	Count() int
	validate(p int)
}

type unionFind struct {
	nodes []int
	size   []int
	count  int
}

func NewUnionFind(size int) *unionFind {
	uf := &unionFind{}
	uf.nodes = make([]int, size)
	uf.size = make([]int, size)
	for i := 0; i < size; i++ {
		uf.nodes[i] = i
		uf.size[i] = 1
	}
	uf.count = 0
	return uf
}

func (uf *unionFind) Union(p int, q int) error {
	rootP, errP := uf.Find(p)
	rootQ, errQ := uf.Find(q)

	if errP != nil || errQ != nil {
		return fmt.Errorf("error: %v, %v", errP, errQ)
	}

	if rootP == rootQ {
		return nil;
	}

	if uf.size[rootP] > uf.size[rootQ]{
		uf.size[rootP] += uf.size[rootQ]
		uf.nodes[rootQ] = rootP
	} else {
		uf.size[rootQ] += uf.size[rootP]
		uf.nodes[rootP] = rootQ
	}

	uf.count--

	return nil
}

func (uf *unionFind) Connected(p int, q int) (bool, error) {
	rootP, errP := uf.Find(p)
	rootQ, errQ := uf.Find(q)

	if errP != nil || errQ != nil {
		return false, fmt.Errorf("error: %v, %v", errP, errQ)
	}

	return rootP == rootQ, nil
}

func (uf *unionFind) Find(p int) (int, error) {
	if err := uf.validate(p); err != nil {
		return -1, err
	}
	root := p
	for root != uf.nodes[root] {
		root = uf.nodes[root]
	}
	for p != root {
		newp := uf.nodes[p]
		uf.nodes[p] = root
		p = newp
	}
	return root, nil
}

func (uf *unionFind) Count() int {
	return uf.count
}

func (uf *unionFind) validate(p int) error {
	n := len(uf.nodes)
	if p < 0 || p >= n {
		err := fmt.Sprint("index ", p, " is not between 0 and ", n-1)
		return fmt.Errorf(err)
	}
	return nil
}
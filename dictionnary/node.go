package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value  uint64
	childs map[byte]*Node
}

func NewNode(value uint64) *Node {
	return &Node{
		value:  value,
		childs: make(map[byte]*Node),
	}
}

func (n *Node) GetValueToPathMap() (map[uint64][]byte, error) {
	values := map[uint64][]byte{}
	for symbol, node := range n.childs {
		if _, ok := values[node.value]; ok {
			return nil, errors.New("Duplicate value in tree")
		}
		values[node.value] = []byte{symbol}
	}
	return values, nil
}

type Searcher struct {
	count   uint64
	path    []byte
	root    *Node
	current *Node
}

func NewSearcher() *Searcher {
	root := NewNode(0)
	return &Searcher{
		count:   0,
		root:    root,
		current: root,
	}
}

func (s *Searcher) Reset() {
	s.path = []byte{}
	s.current = s.root
}

func (s *Searcher) Contains(symbol byte) bool {
	_, ok := s.current.childs[symbol]
	return ok
}

func (s *Searcher) Next(symbol byte) bool {
	if next, ok := s.current.childs[symbol]; ok {
		s.path = append(s.path, symbol)
		s.current = next
		return true
	} else {
		return false
	}
}

func (s *Searcher) AddNode(symbol byte) (uint64, error) {
	debugNewEntry(append(s.path, symbol), s.count)

	if _, ok := s.current.childs[symbol]; ok {
		return 0, errors.New("Child node already exists for this symbol")
	}

	s.current.childs[symbol] = NewNode(s.count)
	s.count++

	if !s.Next(symbol) {
		panic(fmt.Sprintf("Missing node in tree"))
	}

	return s.count - 1, nil
}

func (s *Searcher) GetValue() (uint64, error) {
	if s.current == s.root {
		return 0, errors.New("Root node doesn't contain a value")
	}

	return s.current.value, nil
}

func (s *Searcher) GetCurrentNode() *Node {
	return s.current
}

func (s *Searcher) GetCurrentPath() []byte {
	return s.path
}

package skip

import (
	"errors"
	"math/rand"
)

type node struct {
	forwards []*node
	data     Interface
}

// Interface must be satisfied for data will be stored in skiplist.
// Any type satisfies Interface can be store in skiplist.
type Interface interface {
	// Compare data of reciever with argument
	// Return 0 if equal
	// Return 1/-1 if greater/less
	Compare(Interface) int
}

// Skiplist represents a probabilistic skip list(PSL).
type Skiplist struct {
	header, null    *node
	maxLevel, level int
	possibility     float64
	size            int
}

const (
	// DefaultMaxLevel is default maximum level, can be changed by call SetMaxLevel()
	DefaultMaxLevel = 1 << 6
	// DefaultPossibility is default possibility, can be changed by call SetP()
	DefaultPossibility = 0.5
)

// New return a new initialized Skiplist
func New() *Skiplist {
	header := &node{
		forwards: make([]*node, DefaultMaxLevel+1),
		data:     nil,
	}
	null := new(node)
	for i := 0; i <= DefaultMaxLevel; i++ {
		header.forwards[i] = null
	}
	return &Skiplist{
		header:      header,
		null:        null,
		maxLevel:    DefaultMaxLevel,
		possibility: DefaultPossibility,
		level:       0,
		size:        0,
	}
}

// regard s.null as +Inf
func (s *Skiplist) compare(n *node, data Interface) int {
	if n == s.null {
		return 1
	}
	return n.data.Compare(data)
}

func newNode(data Interface, level int) *node {
	level++
	return &node{
		data:     data,
		forwards: make([]*node, level),
	}
}

func (s *Skiplist) randLevel() int {
	var l int
	for rand.Float64() < s.possibility && l < s.maxLevel {
		l++
	}
	return l
}

// SetMaxLevel set the maximum level of skiplist.
// New maxLevel must greater than former.
func (s *Skiplist) SetMaxLevel(l int) error {
	if l < s.maxLevel {
		return errors.New("new level must greater than former level")
	}
	newForwards := make([]*node, l)
	copy(newForwards, s.header.forwards)
	s.maxLevel = l
	s.header.forwards = newForwards
	return nil
}

// SetP set possibility of skiplist, which is used to increase the level of new node
func (s *Skiplist) SetP(p float64) {
	s.possibility = p
}

// Search return data if found, otherwise return non-nil error
func (s *Skiplist) Search(data Interface) (Interface, error) {
	x := s.header
	for i := s.level; i >= 0; i-- {
		for s.compare(x.forwards[i], data) < 0 {
			x = x.forwards[i]
		}
	}
	// now x.data < data <= x.forwards[0].data
	x = x.forwards[0]
	if s.compare(x, data) == 0 {
		return x.data, nil
	}
	return nil, errors.New("not found")
}

// Insert insert data into Skiplist
func (s *Skiplist) Insert(data Interface) {
	update := make([]*node, s.maxLevel+1)
	x := s.header

	for i := s.level; i >= 0; i-- {
		for s.compare(x.forwards[i], data) < 0 {
			x = x.forwards[i]
		}
		update[i] = x
	}
	x = x.forwards[0]

	if s.compare(x, data) == 0 {
		x.data = data
	} else {
		level := s.randLevel()
		if level > s.level {
			for i := s.level + 1; i <= level; i++ {
				update[i] = s.header
			}
			s.level = level
		}

		x = newNode(data, level)
		for i := 0; i <= level; i++ {
			x.forwards[i] = update[i].forwards[i]
			update[i].forwards[i] = x
		}
	}
}

// Delete remove node with given data from skiplist
func (s *Skiplist) Delete(data Interface) (Interface, error) {
	update := make([]*node, s.maxLevel+1)
	x := s.header

	for i := s.level; i >= 0; i-- {
		for s.compare(x.forwards[i], data) < 0 {
			x = x.forwards[i]
		}
		update[i] = x
	}
	x = x.forwards[0]
	if s.compare(x, data) != 0 {
		return nil, errors.New("can't find data to be deleted")
	}
	for i := 0; i <= s.level; i++ {
		if update[i].forwards[i] != x {
			break
		}
		update[i].forwards[i] = x.forwards[i]
	}

	for s.level > 0 && s.header.forwards[s.level] == s.null {
		s.level--
	}
	return x.data, nil
}

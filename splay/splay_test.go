package splay

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
)

type Int int

const Count = 1 << 19

func (a Int) Compare(b Interface) int {
	y := b.(Int)
	if a > y {
		return 1
	} else if a < y {
		return -1
	}
	return 0
}

func TestNew(t *testing.T) {
	st := New()
	for i := 0; i < (Count); i++ {
		I := Int(i)
		st.Insert(I)
	}

	for i := 0; i < Count; i++ {
		I := Int(i)
		if _, err := st.Search(I); err != nil {
			fmt.Println(i)
			t.Fail()
		}
	}

	var deleted [Count]bool
	for i := 0; i < Count>>1; i++ {
		I := Int(rand.Intn(Count))
		if !deleted[int(I)] {
			_, err := st.Delete(I)
			if err != nil {
				log.Println(int(I), err.Error())
				t.Fail()
			} else {
				deleted[int(I)] = true
			}
		}
	}
}

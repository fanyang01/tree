/*
package llrb implement the Left Leaning Red-Black Tree based on following work:
http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
Note that Robert Sedgewick's paper didn't describe how to perform deletion on Top-down 2-3-4 LLRB, following question on stackoverflow gave the answer:
http://stackoverflow.com/questions/11336167/what-additional-rotation-is-required-for-deletion-from-a-top-down-2-3-4-left-lea
*/
package llrb

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

func TestBU23(t *testing.T) {
	lt := New(BU23)
	for i := 0; i < (Count); i++ {
		I := Int(i)
		lt.Insert(I)
	}

	for i := 0; i < Count; i++ {
		I := Int(i)
		if _, err := lt.Search(I); err != nil {
			fmt.Println(i)
			t.Fail()
		}
	}

	var deleted [Count]bool
	for i := 0; i < Count>>1; i++ {
		I := Int(rand.Intn(Count))
		if !deleted[int(I)] {
			_, err := lt.Delete(I)
			if err != nil {
				log.Println(int(I), err.Error())
				t.Fail()
			} else {
				deleted[int(I)] = true
			}
		}
	}
}
func TestTD234(t *testing.T) {
	lt := New(TD234)
	for i := 0; i < (Count); i++ {
		I := Int(i)
		lt.Insert(I)
	}

	for i := 0; i < Count; i++ {
		I := Int(i)
		if _, err := lt.Search(I); err != nil {
			fmt.Println(i)
			t.Fail()
		}
	}

	var deleted [Count]bool
	for i := 0; i < Count>>1; i++ {
		I := Int(rand.Intn(Count))
		if !deleted[int(I)] {
			_, err := lt.Delete(I)
			if err != nil {
				log.Println(int(I), err.Error())
				t.Fail()
			} else {
				deleted[int(I)] = true
			}
		}
	}
}

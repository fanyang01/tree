package llrb

import (
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
	tr := New(BU23)
	for i := 0; i < (Count); i++ {
		I := Int(i)
		tr.Insert(I)
	}

	for i := 0; i < Count; i++ {
		I := Int(i)
		if _, err := tr.Search(I); err != nil {
			t.Fail()
		}
	}

	var deleted [Count]bool
	for i := 0; i < Count>>1; i++ {
		I := Int(rand.Intn(Count))
		if !deleted[int(I)] {
			_, err := tr.Delete(I)
			if err != nil {
				t.Fail()
			} else {
				deleted[int(I)] = true
			}
		}
	}
}
func TestTD234(t *testing.T) {
	tr := New(TD234)
	for i := 0; i < (Count); i++ {
		I := Int(i)
		tr.Insert(I)
	}

	for i := 0; i < Count; i++ {
		I := Int(i)
		if _, err := tr.Search(I); err != nil {
			t.Fail()
		}
	}

	var deleted [Count]bool
	for i := 0; i < Count>>1; i++ {
		I := Int(rand.Intn(Count))
		if !deleted[int(I)] {
			_, err := tr.Delete(I)
			if err != nil {
				t.Fail()
			} else {
				deleted[int(I)] = true
			}
		}
	}
}

func BenchmarkBU23Insert(b *testing.B) {
	tr := New(BU23)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		I := Int(i)
		tr.Insert(I)
	}
}

func BenchmarkBU23Search(b *testing.B) {
	tr := New(BU23)
	for i := 0; i < b.N; i++ {
		I := Int(i)
		tr.Insert(I)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		I := Int(i)
		tr.Search(I)
	}
}
func BenchmarkBU23Delete(b *testing.B) {
	tr := New(BU23)
	for i := 0; i < b.N; i++ {
		I := Int(i)
		tr.Insert(I)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		I := Int(rand.Intn(Count))
		tr.Delete(I)
	}
}

func BenchmarkTD234Insert(b *testing.B) {
	tr := New(TD234)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		I := Int(i)
		tr.Insert(I)
	}
}
func BenchmarkTD234Search(b *testing.B) {
	tr := New(TD234)
	for i := 0; i < b.N; i++ {
		I := Int(i)
		tr.Insert(I)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		I := Int(i)
		tr.Search(I)
	}
}
func BenchmarkTD234Delete(b *testing.B) {
	tr := New(TD234)
	for i := 0; i < b.N; i++ {
		I := Int(i)
		tr.Insert(I)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		I := Int(rand.Intn(Count))
		tr.Delete(I)
	}
}

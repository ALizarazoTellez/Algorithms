package quicksort_test

import (
	"math/rand/v2"
	"slices"
	"testing"

	"local/quicksort"

	"github.com/google/go-cmp/cmp"
)

func TestQuickSort(t *testing.T) {
	testdata := []struct {
		desc string
		data []int
	}{
		{"basic", []int{4, 2, 6, 8, 3, 1, 5, 7, 9, 0}},
		{"right max", []int{4, 2, 6, 8, 3, 1, 5, 7, 0, 9}},
		{"ordered", []int{1, 2, 3, 4, 5}},
		{"reversed", []int{5, 4, 3, 2, 1}},
	}

	for _, test := range testdata {
		t.Run(test.desc, func(t *testing.T) {
			items := test.data
			t.Log("Slice:", items)

			want := make([]int, len(items))
			copy(want, items)
			slices.Sort(want)

			got := make([]int, len(items))
			copy(got, items)
			quicksort.QuickSort(got)

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("Want (-) | Got (+): %s", diff)
			}
		})
	}
}

func BenchmarkQuickSort_Custom(b *testing.B) {
	slice := makeSlice()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		quicksort.QuickSort(slice)
	}
}

func BenchmarkQuickSort_Stdlib(b *testing.B) {
	slice := makeSlice()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slices.Sort(slice)
	}
}

func makeSlice() []int {
	rand := rand.NewPCG(0, 0).Uint64

	slice := make([]int, 100)
	for i := range slice {
		slice[i] = int(rand())
	}

	return slice
}

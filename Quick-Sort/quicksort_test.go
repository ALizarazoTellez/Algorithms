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
		{"large", makeSlice()},
	}

	for _, test := range testdata {
		t.Run(test.desc, func(t *testing.T) {
			t.Log("Items:", test.data)

			want := duplicateSlice(t, test.data)
			slices.Sort(want)

			got1 := duplicateSlice(t, test.data)
			quicksort.Custom1_QuickSort(got1)

			got2 := duplicateSlice(t, test.data)
			quicksort.AI1_QuickSort(got2)

			compareSlices(t, want, got1)
			compareSlices(t, want, got2)
		})
	}
}

func compareSlices(t *testing.T, want, got []int) {
	t.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Want (-) | Got (+): %s", diff)
	}
}

func duplicateSlice(t *testing.T, slice []int) []int {
	t.Helper()

	dup := make([]int, len(slice))
	copy(dup, slice)

	return dup
}

func BenchmarkCustom1_QuickSort(b *testing.B) {
	slice := makeSlice()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		quicksort.Custom1_QuickSort(slice)
	}
}

func BenchmarkAI1_QuickSort(b *testing.B) {
	slice := makeSlice()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		quicksort.AI1_QuickSort(slice)
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

package quicksort

import "math/rand/v2"

// Custom1_QuickSort sorts a slice in-place.
func Custom1_QuickSort(slice []int) {
	if len(slice) <= 1 {
		return
	}

	pivot := slice[len(slice)-1]

	middle := 0
	for index := 0; index < len(slice)-1; index++ {
		value := slice[index]

		if value <= pivot {
			slice[middle], slice[index] = slice[index], slice[middle]
			middle++
		}
	}

	slice[middle], slice[len(slice)-1] = slice[len(slice)-1], slice[middle]

	Custom1_QuickSort(slice[:middle])
	Custom1_QuickSort(slice[middle:])
}

func AI1_QuickSort(slice []int) {
	quickSortHelper(slice, 0, len(slice)-1)
}

func quickSortHelper(slice []int, low, high int) {
	if low < high {
		pivotIndex := partition(slice, low, high)
		quickSortHelper(slice, low, pivotIndex-1)
		quickSortHelper(slice, pivotIndex+1, high)
	}
}

func partition(slice []int, low, high int) int {
	pivotIndex := rand.IntN(high-low+1) + low
	slice[pivotIndex], slice[high] = slice[high], slice[pivotIndex]
	pivot := slice[high]

	i := low - 1
	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

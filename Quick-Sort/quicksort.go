package quicksort

// QuickSort sorts a slice in-place.
func QuickSort(slice []int) {
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

	QuickSort(slice[:middle])
	QuickSort(slice[middle:])
}

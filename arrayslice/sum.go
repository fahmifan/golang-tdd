package arrayslice

// Sum takes slice of integer and sum the items in array
func Sum(nums []int) (sums int) {
	for _, i := range nums {
		sums += i
	}

	return sums
}

// SumAll takes many slice of integer and sum each of the slice
func SumAll(numToSums ...[]int) (sums []int) {
	for _, nums := range numToSums {
		sums = append(sums, Sum(nums))
	}

	return sums
}

// SumAllTails sum all tails of each slices
// tails are items that apart of the first (head)
func SumAllTails(numToSums ...[]int) (sums []int) {
	for _, nums := range numToSums {
		// check for empty slice
		if len(nums) == 0 {
			sums = append(sums, 0)
			continue
		}

		// take a slice from index 1 to the end
		tail := nums[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}

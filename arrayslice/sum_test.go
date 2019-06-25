package arrayslice

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := Sum(nums)
		want := 15

		if got != want {
			t.Errorf("want %v, got %v, given %v", got, want, nums)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum of many collection", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", got, want)
		}
	}

	t.Run("sum all tails in slices of slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}

		checkSums(t, got, want)
	})

	t.Run("sum all tails in slices of slice", func(t *testing.T) {
		got := SumAllTails([]int{3, 2, 1}, []int{7, 5, 6})
		want := []int{3, 11}

		checkSums(t, got, want)
	})

	t.Run("safely sums empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

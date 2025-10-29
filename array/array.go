package array

func Sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

/*
// iteration 1
func SumAll(arraysOfNums ...[]int) []int {
	sum := make([]int, len(arraysOfNums))
	for i, nums := range arraysOfNums {
		sum[i] = Sum(nums)
	}
	return sum
}
*/

// iteration 2
func SumAll(arraysOfNums ...[]int) []int {
	var sum []int
	for _, nums := range arraysOfNums {
		sum = append(sum, Sum(nums))
	}
	return sum
}

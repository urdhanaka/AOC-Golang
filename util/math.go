package util

func MaxInt(nums ...int) int {
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func SumIntArrays(nums []int) (sum int) {
	sum = 0

	for _, num := range nums {
		sum += num
	}

	return
}

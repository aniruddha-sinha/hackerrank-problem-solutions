package simplearraysum

func simpleArraySum(ar []int32) int32 {
	var sum int32
	for _, elem := range ar {
		sum += elem
	}
	return sum
}

func SimpleArraySumHelper() int32 {
	arr := []int32{1, 2, 4, 8, 16, 32, 64, 128}
	return simpleArraySum(arr)
}

package averybigsum

func aVeryBigSum(ar []int64) int64 {
	var sum int64
	for _, v := range ar {
		sum += v
	}
	return sum
}

func AVeryBigSumProblemHelper() int64 {
	ar := []int64{1000000001, 1000000002, 1000000003, 10000000055, 10000000079}
	return aVeryBigSum(ar)
}

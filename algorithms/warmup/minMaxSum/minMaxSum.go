package minmaxsum

import "log"

func deleteAtIndex(slice []int32, index int32) []int32 {
	var resultantSlice []int32 = make([]int32, 0)
	resultantSlice = append(resultantSlice, slice[:index]...)
	resultantSlice = append(resultantSlice, slice[index+1:]...)
	return resultantSlice
}

func arrSum(arr []int32) int64 {
	var sum int64
	for _, v := range arr {
		sum += int64(v)
	}
	return sum
}

func getMin(slice []int64) int64 {
	var min, i, length int64 = int64(slice[0]), 0, int64(len(slice))
	for i = 1; i < length; i++ {
		if int64(slice[i]) < int64(min) {
			min = int64(slice[i])
		}
	}
	return min
}

func getMax(slice []int64) int64 {
	var max, i, length int64 = int64(slice[0]), 0, int64(len(slice))
	for i = 1; i < length; i++ {
		if int64(slice[i]) > int64(max) {
			max = int64(slice[i])
		}
	}
	return max
}

func MinMaxSum(arr []int32) {
	var resultantSumArrayContainer []int64 = make([]int64, 0)
	for i := range arr {
		resArr := deleteAtIndex(arr, int32(i))
		resultantSumArrayContainer = append(resultantSumArrayContainer, int64(arrSum(resArr)))
	}

	min := getMin(resultantSumArrayContainer)
	max := getMax(resultantSumArrayContainer)

	log.Println(min, " ", max)
}

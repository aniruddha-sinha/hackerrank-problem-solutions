package minmaxsum

import "log"

func deleteAtIndex(slice []int32, index int32) []int32 {
	var resultantSlice []int32 = make([]int32, 0)
	resultantSlice = append(resultantSlice, slice[:index]...)
	resultantSlice = append(resultantSlice, slice[index+1:]...)
	return resultantSlice
}

func arrSum(arr []int32) int32 {
	var sum int32
	for _, v := range arr {
		sum += v
	}
	return sum
}

func getMin(slice []int32) int32 {
	var min, i, length int32 = slice[0], 0, int32(len(slice))
	for i = 1; i < length; i++ {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func getMax(slice []int32) int32 {
	var max, i, length int32 = slice[0], 0, int32(len(slice))
	for i = 1; i < length; i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MinMaxSum(arr []int32) {
	var resultantSumArrayContainer []int32 = make([]int32, 0)
	for i, _ := range arr {
		resArr := deleteAtIndex(arr, int32(i))
		resultantSumArrayContainer = append(resultantSumArrayContainer, arrSum(resArr))
	}

	min := getMin(resultantSumArrayContainer)
	max := getMax(resultantSumArrayContainer)

	log.Println(min, " ", max)
}

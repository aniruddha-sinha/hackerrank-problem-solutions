package plusminus

import "fmt"

func plusMinus(arr []int32) (float32, float32, float32) {
	len := float32(len(arr))
	var negCounter, posCounter, zeroCounter int32
	for _, val := range arr {
		if val < 0 {
			negCounter++
		} else if val > 0 {
			posCounter++
		} else {
			zeroCounter++
		}
	}

	return float32(posCounter) / len, float32(negCounter) / len, float32(zeroCounter) / len
}

func PlusMinusFunctionHelper() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	posRatio, negRatio, zeroRatio := plusMinus(arr)
	fmt.Printf("%.6f\n", posRatio)
	fmt.Printf("%.6f\n", negRatio)
	fmt.Printf("%.6f\n", zeroRatio)
}

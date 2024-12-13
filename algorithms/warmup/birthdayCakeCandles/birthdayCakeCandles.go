package birthdaycakecandles

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var tallestCandleCount int32
	var tallestCandleValue int32 = candles[0]

	for _, value := range candles {
		if value > tallestCandleValue {
			tallestCandleValue = value
		} else if value == tallestCandleValue {
			tallestCandleCount++
		}
	}

	return tallestCandleCount
}

func BirthdayCakeCandlesHelper(candles []int32) int32 {
	//candles := []int32{3, 2, 1, 3}
	tallestCandleCount := birthdayCakeCandles(candles)
	return tallestCandleCount
}

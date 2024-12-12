package main

import (
	"log"

	averybigsum "github.com/aniruddha-sinha/hackerrank/aVeryBigSum"
	birthdaycakecandles "github.com/aniruddha-sinha/hackerrank/birthdayCakeCandles"
	comparethetriplets "github.com/aniruddha-sinha/hackerrank/compareTheTriplets"
	diagonaldifference "github.com/aniruddha-sinha/hackerrank/diagonalDifference"
	minmaxsum "github.com/aniruddha-sinha/hackerrank/minMaxSum"
	plusminus "github.com/aniruddha-sinha/hackerrank/plusMinus"
	simplearraysum "github.com/aniruddha-sinha/hackerrank/simpleArraySum"
	solvemefirst "github.com/aniruddha-sinha/hackerrank/solveMeFirst"
	staircaseproblem "github.com/aniruddha-sinha/hackerrank/staircaseProblem"
	timeconversion "github.com/aniruddha-sinha/hackerrank/timeConversion"
)

func main() {
	log.Println("Check the question in the problem folder : <problem>.md")
	log.Println("Solve Me First Problem")
	log.Println(solvemefirst.SolveMeFirstHelperFun())
	log.Println("Simple Array Sum")
	log.Println(simplearraysum.SimpleArraySumHelper())
	log.Println("Compare the triplets")
	log.Println(comparethetriplets.CompareTripletsHelper())
	log.Println("A Very Big Sum")
	log.Println(averybigsum.AVeryBigSumProblemHelper())
	log.Println("Diagonal Difference in a Square Matrix")
	log.Println(diagonaldifference.DiagonalDifferenceFunctionHelper())
	log.Println("Plus Minus")
	plusminus.PlusMinusFunctionHelper()
	log.Println("Staircase Problem")
	staircaseproblem.StaircaseProblem(int32(6))
	log.Println("Min Max Sum")
	minmaxsum.MinMaxSum([]int32{254961783, 604179258, 462517083, 967304281, 860273491})
	log.Println("Min Max Sum second test case")
	minmaxsum.MinMaxSum([]int32{1, 2, 3, 4, 5})
	log.Println("Min Max Sum third test case")
	minmaxsum.MinMaxSum([]int32{942381765, 627450398, 954173620, 583762094, 236817490})
	log.Println("Birthday cakes and candles problem")
	log.Println(birthdaycakecandles.BirthdayCakeCandlesHelper([]int32{44, 53, 31, 27, 77, 60, 66, 77, 26, 36}))
	log.Println("Birthday cakes and candles problem use case 2 :: 18 90 90 13 90 75 90 8 90 43")
	log.Println(birthdaycakecandles.BirthdayCakeCandlesHelper([]int32{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}))
	log.Println("Time Conversion")
	log.Println(timeconversion.TimeConversionHelper("07:05:45PM"))
	log.Println("Time Conversion second test case")
	log.Println(timeconversion.TimeConversionHelper("12:40:22AM"))
	log.Println("Time Conversion third test case")
	log.Println(timeconversion.TimeConversionHelper("06:40:03AM"))
}

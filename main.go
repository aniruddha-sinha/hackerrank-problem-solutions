package main

import (
	"log"

	averybigsum "github.com/aniruddha-sinha/hackerrank/aVeryBigSum"
	comparethetriplets "github.com/aniruddha-sinha/hackerrank/compareTheTriplets"
	diagonaldifference "github.com/aniruddha-sinha/hackerrank/diagonalDifference"
	minmaxsum "github.com/aniruddha-sinha/hackerrank/minMaxSum"
	plusminus "github.com/aniruddha-sinha/hackerrank/plusMinus"
	simplearraysum "github.com/aniruddha-sinha/hackerrank/simpleArraySum"
	solvemefirst "github.com/aniruddha-sinha/hackerrank/solveMeFirst"
	staircaseproblem "github.com/aniruddha-sinha/hackerrank/staircaseProblem"
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
	//942381765 627450398 954173620 583762094 236817490
	log.Println("Min Max Sum third test case")
	minmaxsum.MinMaxSum([]int32{942381765, 627450398, 954173620, 583762094, 236817490})
}

package solvemefirst

func solveMeFirst(a, b uint32) uint32 {
	return a + b
}

func SolveMeFirstHelperFun() uint32 {
	a, b := uint32(2), uint32(3)
	return solveMeFirst(a, b)
}

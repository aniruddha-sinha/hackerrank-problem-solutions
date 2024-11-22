package comparethetriplets

func compareTriplets(a []int32, b []int32) []int32 {
	var alicePoint, bobPoint int32
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alicePoint++
		} else if b[i] > a[i] {
			bobPoint++
		}
	}
	return []int32{alicePoint, bobPoint}
}

func CompareTripletsHelper() []int32 {
	alice := []int32{-1, 2, 4}
	bob := []int32{2, 0, 10}
	res := compareTriplets(alice, bob)
	return res
}

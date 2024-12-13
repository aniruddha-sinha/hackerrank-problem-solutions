package diagonaldifference

func diagonalDifference(arr [][]int32) int32 {
	var leftDiag, rightDiag, absDiagDiff int32
	for rowIndex, rowValue := range arr {
		for columnIndex, columnValue := range rowValue {
			if rowIndex == columnIndex {
				leftDiag += columnValue
			}

			if rowIndex+columnIndex == len(rowValue)-1 {
				rightDiag += columnValue
			}
		}
	}

	if leftDiag > rightDiag {
		absDiagDiff = leftDiag - rightDiag
	} else {
		absDiagDiff = rightDiag - leftDiag
	}
	return absDiagDiff
}

func DiagonalDifferenceFunctionHelper() int32 {
	twoDArr := [][]int32{
		{1, 2, 3, 44},
		{5, 16, 772, 8},
		{9, 107, 11, 12},
		{134, 14, 15, 21},
	}

	return diagonalDifference(twoDArr)
}

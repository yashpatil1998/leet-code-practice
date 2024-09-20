// https://leetcode.com/problems/pascals-triangle

package EasyAlgoAndStrings

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		result[i-1] = make([]int, i)
	}
	result[0][0] = 1
	if numRows == 1 {
		return result
	}
	for i := 1; i < numRows; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}
	return result
}

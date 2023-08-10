// https://leetcode.com/problems/search-a-2d-matrix

package MediumAlgo

func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        if target <= matrix[i][n-1] && target >= matrix[i][0] {
            for j := 0; j < n; j++ {
                if target == matrix[i][j] {
                    return true
                }
            }
        } else {
            continue
        }
    }
    return false
}
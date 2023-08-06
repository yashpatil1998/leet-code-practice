// https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle

package MediumAlgo

func countPoints(points [][]int, queries [][]int) []int {
    var ans []int
    for i := 0; i < len(queries); i++ {
        validPoints := 0
        for j := 0; j < len(points); j++ {
            if dist(points[j], queries[i][:2]) <= (queries[i][2] * queries[i][2]) {
                validPoints++
            }
        }
        ans = append(ans, validPoints)
    }
    return ans
}

func dist(xy1,xy2 []int) int {
    return ((xy2[1] - xy1[1]) * (xy2[1] - xy1[1])) + ((xy2[0] - xy1[0]) * (xy2[0] - xy1[0]))
}

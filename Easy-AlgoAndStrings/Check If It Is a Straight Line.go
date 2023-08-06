// https://leetcode.com/problems/check-if-it-is-a-straight-line

package EasyAlgoAndStrings

func checkStraightLine(coordinates [][]int) bool {
    if len(coordinates) <= 2 {
        return true
    }
    var slope float64
    if coordinates[0][0] == coordinates[1][0] {
        for i := 3; i < len(coordinates); i++ {
            if coordinates[i][0] != coordinates[0][0] {
                return false
            }
        }
        return true
    }
    slope = calcSlope(coordinates[0], coordinates[1])
    yint := calcYInt(slope, coordinates[0])
    for i := 2; i < len(coordinates); i++ {
        if !isOnLine(coordinates[i], slope, yint) {
            return false
        }
    }
    return true
}

func calcSlope(xy1, xy2 []int) float64 {
    return (float64)(((float64)(xy2[1]) - (float64)(xy1[1])) / ((float64)(xy2[0]) - (float64)(xy1[0])))
}

func calcYInt(slope float64, xy []int) float64 {
    yint :=  (float64)(xy[1]) - (slope * (float64)(xy[0]))
    return yint
}

func isOnLine(xy []int, slope float64, yint float64) bool {
    return (float64)(xy[1]) == (slope * (float64)(xy[0])) + yint
}
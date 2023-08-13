// https://leetcode.com/problems/minimum-lines-to-represent-a-line-chart

package MediumAlgo

import "sort"

func minimumLines(stockPrices [][]int) int {
	if len(stockPrices) == 2 {
		return 1
	}
	if len(stockPrices) == 1 {
		return 0
	}
	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})
	lineCount := 1
	for i := 2; i < len(stockPrices); i++ {
		if !isLine(stockPrices[i-2 : i+1]) {
			lineCount++
		}
	}
	return lineCount
}

func isLine(xy [][]int) bool {
	return 0 == (xy[0][0]*(xy[1][1]-xy[2][1]) + xy[1][0]*(xy[2][1]-xy[0][1]) + xy[2][0]*(xy[0][1]-xy[1][1]))
}

// https://leetcode.com/problems/calculate-score-after-performing-instructions/

package MediumAlgo

func calculateScore(instructions []string, values []int) int64 {
	n := len(values)
	sum := int64(0)
	visitedI := make(map[int]bool)
	for i := 0; i < n && i >= 0 && !visitedI[i]; {
		// fmt.Println("i: ", i, ", values[i]: ", values[i], ", instructions[i]: ", instructions[i], ", visitedI[i]", visitedI[i])
		visitedI[i] = true
		if instructions[i] == "add" {
			sum += int64(values[i])
			i++
		} else {
			i = i + values[i]
		}
	}

	return sum
}

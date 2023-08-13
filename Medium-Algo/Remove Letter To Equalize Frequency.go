// https://leetcode.com/problems/remove-letter-to-equalize-frequency

package MediumAlgo

import "sort"

func equalFrequency(word string) bool {

	// example abcc

	freq := make(map[string]int)
	for _, v := range word {
		singleChar := string(v)
		freq[singleChar]++
	}

	// freq = [a:1 b:1 c:2]
	if len(freq) == 1 {
		return true
	}

	countFreq := make(map[int]int) // key = frequency, value = count of appearances of keys in freq map
	for _, count := range freq {
		countFreq[count]++
	}
	// countFreq = [1:2 2:1]

	counts := []int{}
	for v := range countFreq {
		counts = append(counts, v)
	}
	// counts = [2 1]

	if len(counts) == 1 && counts[0] == 1 {
		return true
	}

	if len(counts) == 2 {
		sort.Ints(counts)

		// counts = [1 2]
		if counts[1]-counts[0] == 1 && countFreq[counts[1]] == 1 { // if one cahracter has occured one more than other characters
			return true
		}

		if len(countFreq) == 2 && counts[0] == 1 && countFreq[1] == 1 {
			return true
		}
	}

	return false
}

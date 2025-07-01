// https://leetcode.com/problems/find-the-original-typed-string-i

package EasyAlgoAndStrings

func possibleStringCount(word string) int {
	freq := make(map[int]int)
	counter := 0
	freq[counter]++
	for i := 1; i < len(word); i++ {
		if word[i] != word[i-1] {
			counter++
		}
		freq[counter]++
	}
	if len(freq) == len(word) {
		return 1
	}

	ans := 1

	for _, f := range freq {
		ans += f - 1
	}

	return ans
}

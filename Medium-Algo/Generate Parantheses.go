package MediumAlgo

func generateParenthesis(n int) []string {
	ans := make([]string, 0)

	genCombos(n, n, n, "", &ans)

	return ans
}

func genCombos(n, open, clos int, currCombo string, ans *[]string) {
	if len(currCombo) == 2*n {
		*ans = append(*ans, currCombo)
		return
	}

	if open > 0 {
		genCombos(n, open-1, clos, currCombo+"(", ans)
	}
	if clos > open {
		genCombos(n, open, clos-1, currCombo+")", ans)
	}
}

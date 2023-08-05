// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

package EasyAlgoAndStrings

import (
    "strings"
)

func strStr(haystack string, needle string) int {
    return strings.Index(haystack, needle)
}
package str

import "strings"

func lengthOfLastWord(s string) int {

	arr := strings.Split(strings.TrimRight(s, " "), " ")
	if len(arr) == 0 {
		return 0
	}
	return len(arr[len(arr)-1])
}

package str

import (
	"strconv"
)

func addBinary(a string, b string) string {

	v1, _ := strconv.ParseInt(a, 2, 64)
	v2, _ := strconv.ParseInt(b, 2, 64)
	v := v1 + v2
	return strconv.FormatInt(v, 2)
}

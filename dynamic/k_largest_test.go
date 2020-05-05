package dynamic

import (
	"fmt"
	"testing"
)

func Test_aaa(t *testing.T) {

	k := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(k.Add(3))
	fmt.Println(k.Add(5))
	fmt.Println(k.Add(10))
	fmt.Println(k.Add(9))
	fmt.Println(k.Add(4))
}

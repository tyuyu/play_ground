package stack

import (
	"unicode"
)

//实现一个基本的计算器来计算一个简单的字符串表达式的值。
//
//字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
//
//示例 1:
//
//输入: "3+2*2"
//输出: 7
//示例 2:
//
//输入: " 3/2 "
//输出: 1
//示例 3:
//
//输入: " 3+5 / 2 "
//输出: 5
//说明：
//
//你可以假设所给定的表达式都是有效的。
//请不要使用内置的库函数 eval。

func calculate(s string) int {

	num := 0
	st := make([]int, 0)
	sign := byte('+')
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) { //如果是多位数字
			num = num*10 + int(s[i]-'0')
		}
		//如果当前字符是运算符，或者是最后一个字符（一定是数字）
		if s[i] != ' ' && !unicode.IsDigit(rune(s[i])) || i == len(s)-1 {
			switch sign {
			case '+':
			case '-':
				num = -num
			case '*':
				num = st[len(st)-1] * num
				st = st[:len(st)-1]
			case '/':
				num = st[len(st)-1] / num
				st = st[:len(st)-1]
			}
			sign = s[i]
			st = append(st, num)
			num = 0
		}
	}
	for _, v := range st {
		num += v
	}
	return num

}

package str

import "strings"

//6. Z 字形变换
//将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
//
//比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
//
//L   C   I   R
//E T O E S I I G
//E   D   H   N
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
//
//请你实现这个将字符串进行指定行数变换的函数：
//
//string convert(string s, int numRows);
//示例 1:
//
//输入: s = "LEETCODEISHIRING", numRows = 3
//输出: "LCIRETOESIIGEDHN"
//示例 2:
//
//输入: s = "LEETCODEISHIRING", numRows = 4
//输出: "LDREOEIIECIHNTSG"
//解释:
//
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	//按numRows+numRows-2的数量分块
	n := numRows*2 - 2
	rows := make([]string, numRows)
	for i := 0; i < len(s); i++ {
		ii := i % n
		if ii > numRows-1 {
			ii = n - ii
		}
		rows[ii] = rows[ii] + string(s[i])
	}
	return strings.Join(rows, "")
}

//执行结果：
//通过
//显示详情
//执行用时 :
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗 :
//5.1 MB
//, 在所有 Go 提交中击败了
//25.00%
//的用户

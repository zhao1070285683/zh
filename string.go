package main

//todo  最长回文子串
func longestPalindrome(s string) string {
	return ""
}

//字符串： 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
/*
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

示例 1：
输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
示例 3：

输入：s = "A", numRows = 1
输出："A"

func convert(s string, numRows int) string {
	arr := make([][]uint8, numRows)
	for k := 0; k < len(s); {
		for i := 0; i < numRows && k < len(s); i++ {
			arr[i] = append(arr[i], s[k])
			k++
		}

		for i := numRows - 2; i > 0 && k < len(s); i-- {
			arr[i] = append(arr[i], s[k])
			k++
		}
	}

	var res []uint8
	for i := range arr {
		fmt.Println(string(arr[i]))
		res = append(res, arr[i]...)
	}

	return string(res)

}

*/

//字符串：查找字符串数组中的最长公共前缀。
/*
如果不存在公共前缀，返回空字符串 ""。



示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。


提示：

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成


func longest(strArr []string) string {
	if len(strArr) == 0 {
		return ""
	}
	var min = len(strArr[0])
	for _, str := range strArr {
		if len(str) < min {
			min = len(str)
		}
	}

	var res int
	for i := 0; i < min; i++ {
		var notEqual bool
		for j := 0; j < len(strArr)-1; j++ {
			if strArr[j][i] != strArr[j+1][i] {
				notEqual = true
				break
			}
		}
		if notEqual {
			res = i
			break
		}
	}

	fmt.Println(res)

	return strArr[0][:res]
}
*/


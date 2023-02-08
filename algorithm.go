package main

/*
1 二分法
2 排序：冒泡、简单选择、直接插入、快排、堆、希尔、归并、
3 动态规划
4 递归
5 双指针
*/

func dupstring(s string) string {
	var start, end int
	var max int
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			var dup bool
			for k := i; k < j; k++ {
				if s[j] == s[k] {
					dup = true
					break
				}
			}
			if dup {
				break
			}

			if max < (j-i+1) {
				start = i
				end = j
				max = j - i + 1
			}

		}
	}

	return s[start : end+1]
}

package main

func main() {

}

/*
1 线性表：数组、栈 队列
2 链表：单链表、双向链表
3 字符串： 朴素算法
4 树：前/中/后序遍历、二叉树、二叉搜索树
5 散列表：哈希和解决冲突
*/

// 最小栈
/*type stack struct {
	s   []int
	min []int
	len int
}

func new() *stack {
	return &stack{}
}

func (st *stack) push(v int) {
	st.s = append(st.s, v)
	if len(st.min) == 0 || st.min[len(st.min)-1] > v {
		st.min = append(st.min, v)
	} else {
		st.min = append(st.min, st.min[len(st.min)-1])
	}
	st.len++
}
func (st *stack) pop() int {
	if st.len == 0 {
		return -1
	}
	res := st.s[st.len-1]
	st.s = st.s[:st.len-1]
	st.min = st.min[:st.len-1]
	st.len--
	return res
}
func (st *stack) minVal() int {
	if st.len == 0 {
		return -1
	}
	return st.min[st.len-1]
}*/

//两个队列模拟栈
/*type stack struct {
	queueA []int
	queueB []int
	len    int
}

func newStack() stack {
	return stack{}
}
func (s *stack) push(v int) {
	if len(s.queueA) == 0 && len(s.queueB) == 0 {
		s.queueA = append(s.queueA, v)
		return
	}
	if len(s.queueA) > 0 {
		s.queueA = append(s.queueA, v)
	} else if len(s.queueB) > 0 {
		s.queueB = append(s.queueB, v)
	}

}
func (s *stack) pop() int {
	if len(s.queueA) > 0 {
		for i := 0; i < len(s.queueA)-1; i++ {
			s.queueB = append(s.queueB, s.queueA[i])
		}
		s.len--
		res := s.queueA[len(s.queueA)-1]
		s.queueA = nil
		return res
	} else if len(s.queueB) > 0 {
		for i := 0; i < len(s.queueB)-1; i++ {
			s.queueA = append(s.queueA, s.queueB[i])
		}
		s.len--
		res := s.queueB[len(s.queueB)-1]
		s.queueB = nil
		return res
	}

	return -1
}*/

//冒泡 两两交换
func bubble(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	return arr
}

// 最长回文子串 babad 正着读和反着读都是一样的。
func longestPalindrome() string {
	return ""
}

/* 深度优先、广度优先 */

type treeNode struct {
	Data  int
	Left  *treeNode
	Right *treeNode
}

/*子树
深度优先搜索？
*/
func isTreeSubstructure(root *treeNode, sub *treeNode) bool {
	return false
}

/*x:=[]int{4:44,55,66,1:77,88}
fmt.Println(len(x),x[2])
a,b:=f([]string{"test"})
fmt.Println(a==nil,b==nil)*/
func f(x []string) (_, _ []string) {
	_, _ = x, x
	return
}

// 1. golang强类型type自定义类型和原生的是等价的，但是用相同原生类型定义出的两个自定义类型，他们之间是不相等的。
func typeConversion() {
	/*type test []string
	var t test
	var s []string
	s=t
	type zh []string
	var z zh
	t=z

	func() []string {
		var t test
		return t
	}()*/
}

/* 涉及知识点
类型转换： string,[]byte{} 互转

转换方式两种：
1。 强制转换 不安全的指针转换 详情待补充。。。
2。 标准转换 string([]byte) 这种会有内存分配，并且性能没有强制转换好。因为string是只读的
*/

//fmt.Println(replace(" hello  world ", "%20", ' '))
func replace(src, target string, replace byte) string {
	if src == "" {
		return src
	}

	targetByte := []byte(target)

	res := make([]byte, 0, 10)
	for i := range src {
		if src[i] == replace {
			res = append(res, targetByte...)
		} else {
			res = append(res, src[i])
		}
	}

	return string(res)
}

/*arr := make([][]int, 4)
fmt.Println(len(arr), arr)
for i := 0; i < 4; i++ {
	arr[i] = make([]int, 4)
	fmt.Println(len(arr[i]), arr[i])
}
arr[0] = []int{1, 2, 8, 9}
arr[1] = []int{2, 4, 9, 12}
arr[2] = []int{4, 7, 10, 13}
arr[3] = []int{6, 8, 11, 15}

fmt.Println(FindNumFrom2dArray(arr, 16))*/

func FindNumFrom2dArray(arr [][]int, number int) bool {
	i := len(arr) - 1
	j := 0

	for ; i >= 0; i-- {
		for ; j <= len(arr[i])-1; {
			if arr[i][j] > number {
				break
			} else if arr[i][j] < number {
				j++
			} else {
				return true
			}
		}
		if j > len(arr[i])-1 {
			break
		}
	}

	return false
}

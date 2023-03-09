package main

import "fmt"

func main() {
	head := createLinked()
	for s := head; s != nil; s = s.Next {
		fmt.Println(s.val)
	}

}

/*
1 线性表：数组、栈 队列
2 链表：单链表、双向链表
3 字符串： 朴素算法
4 树：前/中/后序遍历、二叉树、二叉搜索树
5 散列表：哈希和解决冲突
*/
type node struct {
	val   int
	left  *node
	right *node
}

//前序遍历
func preOrder(root *node) {
	var stack []*node
	stack = append(stack, root)
	for ; len(stack) != 0; {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		println(top.val)

		if top.right != nil {
			stack = append(stack, top.right)
		}

		if top.left != nil {
			stack = append(stack, top.left)
		}
	}
}

func preOrder2(root *node) {
	if root == nil {
		return
	}
	println(root.val)
	preOrder2(root.left)
	preOrder2(root.right)
}

//链表逆置
/*func reverse(head *Node) *Node {
	if head == nil {
		return head
	}
	p := head
	q := p.Next
	for ; q != nil; {
		s := q.Next
		q.Next = p
		p = q
		q = s
	}
	head.Next = nil

	return p
}*/

//1. 找出两个链表的交点
/*
160\. Intersection of Two Linked Lists (Easy)

[Leetcode](https://leetcode.com/problems/intersection-of-two-linked-lists/description/) / [力扣](https://leetcode-cn.com/problems/intersection-of-two-linked-lists/description/)

例如以下示例中 A 和 B 两个链表相交于 c1：

```html
A:          a1 → a2
                    ↘
                      c1 → c2 → c3
                    ↗
B:    b1 → b2 → b3
```

但是不会出现以下相交的情况，因为每个节点只有一个 next 指针，也就只能有一个后继节点，而以下示例中节点 c 有两个后继节点。

```html
A:          a1 → a2       d1 → d2
                    ↘  ↗
                      c
                    ↗  ↘
B:    b1 → b2 → b3        e1 → e2
```



要求时间复杂度为 O(N)，空间复杂度为 O(1)。如果不存在交点则返回 null。

设 A 的长度为 a + c，B 的长度为 b + c，其中 c 为尾部公共部分长度，可知 a + c + b = b + c + a。

当访问 A 链表的指针访问到链表尾部时，令它从链表 B 的头部开始访问链表 B；同样地，当访问 B 链表的指针访问到链表尾部时，令它从链表 A 的头部开始访问链表 A。这样就能控制访问 A 和 B 两个链表的指针能同时访问到交点。

如果不存在交点，那么 a + b = b + a，以下实现代码中 l1 和 l2 会同时为 null，从而退出循环。
*/

type Node struct {
	val  int
	Next *Node
}

func createLinked() *Node {
	/*head:=&Node{}*/
	head := new(Node)
	p := head
	for i := 0; i < 10; i++ {
		p.val = i
		p.Next = new(Node)
		p = p.Next
	}
	return head
}

/*func IntersectionNode(a, b *Node) *Node {

}*/

//将数组中的零挪到后面
/*func moveZero(arr []int) []int {
for i, j := 0, len(arr)-1; i < j; {
	if arr[i] == 0 && arr[j] != 0 {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
		j--
	} else if arr[i] == 0 && arr[j] == 0 {
		j--
	} else {
		i++
	}
}

var count int
for i := 0; i < len(arr) && count < len(arr); {
fmt.Println(arr, i)
if arr[i] == 0 {
arr = append(arr[:i], arr[i+1:len(arr)]...)
arr = append(arr, 0)
} else {
i++
}
count++
}

return arr
}*/

// 队列：滑动窗口中的最大值
/*给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。

示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]

func windows(arr []int, size int) []int {
	var i, j int
	var res []int

	j = i + size - 1
	if j > len(arr)-1 {
		j = len(arr) - 1
	}
	for i := 0; i < len(arr) && j < len(arr); i++ {
		var max int
		fmt.Println(arr[i],res)
		if i == 0 || (len(res) > 0 && arr[i-1] == res[len(res)-1]) {
			for k := i; k <= j; k++ {
				if arr[k] > max {
					max = arr[k]
				}
			}
		} else {
			if arr[j] > res[len(res)-1] {
				max = arr[j]
			} else if len(res) > 0 {
				max = res[len(res)-1]
			}
		}
		res = append(res, max)
		j++
	}

	return res
}
*/

// 栈：给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水
/*。
示例 1：

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

func sum(height []int) int {
	var res int
	for i := 0; i < len(height); {
		var flag bool
		for j := i + 1; j < len(height); j++ {
			if height[j] >= height[i] {
				for k := i + 1; k < j; k++ {
					if height[i]-height[k] > 0 {
						res += height[i] - height[k]
					}
				}
				i = j
				flag = true
				break
			}
		}

		if !flag {
			i++
			flag = false
		}

	}

	return res
}
*/

//用栈实现括号匹配
/*func match(s string) bool {
	var stack []int32
	for _, ss := range s {
		if ss == '{' {
			stack = append(stack, ss)
		} else if ss == '}' {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}*/

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

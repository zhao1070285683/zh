package main

func main() {

}

type treeNode struct {
	Data int
	Left *treeNode
	Right *treeNode
}

/*树的子结构
深度优先搜索？
*/
func isTreeSubstructure(root *treeNode, sub *treeNode) bool{
	return false
}

// 1. golang强类型type自定义类型和原生的是等价的，但是用相同原生类型定义出的两个自定义类型，他们之间是不相等的。
func typeConversion(){
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

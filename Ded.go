package main

import "fmt"

func main() {
	var n int
	fmt.Println("输入切片初始元素个数：")
	fmt.Scanf("%d", &n)
	var s = []int{}
	var mp = make(map[int]bool)
	fmt.Println("输入初始元素:")
	for i := 1; i <= n; i++ {
		var x int
		fmt.Scanf("%d", &x)
		s = append(s, x)
	}
	for i := 0; i < len(s); i++ {
		if mp[s[i]] {
			s = append(s[:i], s[i+1:]...)
			i--
		} else {
			mp[s[i]] = true
		}
	}
	fmt.Println("去重后的切片为：", s)
}

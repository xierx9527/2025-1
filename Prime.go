package main

import "fmt"

func isP(x int) bool {
	if x == 2 {
		return true
	} else {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
	}
	return true
}
func Prime(n int) []int {
	var res = []int{}
	for i := 2; i <= n; i++ {
		if isP(i) {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var a []int = Prime(n)
	var b []int = Prime(100)
	fmt.Printf("小于%d的质数有 %d\n", n, a)
	fmt.Println("小于100的质数有", b)
}

package main

import "fmt"

func merge(a []int, b []int, n int, m int) []int {
	i := n - 1
	j := m - 1
	k := n + m - 1
	a = append(a, make([]int, m)...)
	for i >= 0 && j >= 0 {
		if a[i] < b[j] {
			a[k] = b[j]
			j--
		} else if b[j] < a[i] {
			a[k] = a[i]
			i--
		}
		k--
	}
	for i >= 0 {
		a[k] = a[i]
		i--
		k--
	}
	for j >= 0 {
		a[k] = b[j]
		j--
		k--
	}
	return a
}
func main() {
	a := []int{1, 2, 4, 8, 9}
	b := []int{3, 5, 7}
	fmt.Println(a)
	fmt.Println(b)
	a = merge(a, b, len(a), len(b))
	fmt.Println(a)
}

package main

import "fmt"

func main() {
	m := []int{0,1,2,2,3,3,5}
	count := deleteduplicate(m)
	for i := 0;i < count; i++ {
		fmt.Println(m[i])
	}
}

func deleteduplicate(m []int) int {
	len := len(m)
	count := 0
	for i := 1; i < len; i++ {
		if m[count] != m[i] {
			count++
			m[count] = m[i]
		}
	}
	return count + 1
}
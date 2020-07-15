package main 

import "fmt"

func main() {
	arr:= [5]int {20,3,7,1,15}
	len:= len(arr)
	for i:=1;i<len;i++ {
		for j:=0;j<len-i;j++ {
			if arr[j]>arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr)
}
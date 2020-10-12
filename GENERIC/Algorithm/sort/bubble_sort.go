package main

import "fmt"

func bubbleSort(arr []int){
	len := len(arr)
	isSwap := false
	for i:=0;i<len-1;i++{
		isSwap = false
		for j:= 0; j<len-i-1;j++{
			if arr[j] > arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j] //swap
				isSwap = true
			}
		}
		if isSwap == false{
			break
		}
	}
	fmt.Println("\narray after insertion sort :")
	for i:=0; i<len;i++{
		fmt.Printf("%d ",arr[i])
	}
}

func main(){
	sample := []int{55,11,22,44,33}
	bubbleSort(sample)

	sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	bubbleSort(sample)
}


package main

import "fmt"

func insertionSort(arr []int){
	for i:=1;i<len(arr);i++{
		for j:= 0; j<i;j++{
			if arr[j] > arr[i]{
				arr[j],arr[i] = arr[i],arr[j] //swap
			}
		}
	}
	fmt.Println("\narray after insertion sort :")
	for i:=0; i<len(arr);i++{
		fmt.Printf("%d ",arr[i])
	}
}

func main(){
	sample := []int{22,55,33,44,11}
	insertionSort(sample)

	sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	insertionSort(sample)
}
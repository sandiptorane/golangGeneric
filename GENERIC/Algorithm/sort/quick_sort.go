package main

import (
	"fmt"
)

func quicksort(arr []int)[]int{
	len := len(arr)
	if len<2{
		return arr
	}
	low ,high := 0,len-1
	pivot := high
	for i,_:=range arr{
		if arr[i]< arr[pivot]{
			arr[i],arr[low] =arr[low],arr[i]
			low++
		}
	}
	arr[low], arr[pivot] = arr[pivot], arr[low]

	quicksort(arr[:low])
	quicksort(arr[low+1:])

	return arr
}

func main(){
	slice := []int{55,36,45,70,2,10}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
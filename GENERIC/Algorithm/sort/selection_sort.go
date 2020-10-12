package main

import "fmt"

func selectionSort(arr []int){
	for i:=0;i<len(arr)-1;i++{
		minindex := i
		for j:=i+1; j< len(arr);j++{
			if arr[j] < arr[minindex]{
				arr[j],arr[minindex] = arr[minindex],arr[j] //swap
			}
		}
	}
	/*fmt.Println("\narray after insertion sort :")
	for i:=0; i<len(arr);i++{
		fmt.Printf("%d ",arr[i])
	}*/
}

func main(){
	sample := []int{22,55,33,44,11}
	selectionSort(sample)
	fmt.Println("\narray after insertion sort :")
	for i:=0; i<len(sample);i++{
		fmt.Printf("%d ",sample[i])
	}

	sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	selectionSort(sample)
	fmt.Println("\narray after insertion sort :")
	for i:=0; i<len(sample);i++{
		fmt.Printf("%d ",sample[i])
	}
}

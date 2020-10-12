package main

import "fmt"

func binarySearch(arr []int,low int,high int,key int)bool{
     if high>=low {
		 mid := (low + high) / 2
		 if arr[mid] == key {
			 return true
		 }
		 if arr[mid] > key {
			 return binarySearch(arr, low, mid-1, key)
		 } else {
			 return binarySearch(arr, mid+1, high, key)
		 }
	 }
	return false
}
func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100,124}
	key := 3
	fmt.Println(key,"in array is",binarySearch( items,0,len(items),key))
}

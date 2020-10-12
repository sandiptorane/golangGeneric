package main

import (
	"fmt"
)

func linearSearch(arr []int,value int)bool{
	for i,_:= range arr{
		if arr[i]== value{
			return true
		}
	}
	return false
}

func main(){
	arr :=[]int{22,11,55,77,66,88}
	value := 55
	found := linearSearch(arr,value)
	if found {
		fmt.Println(value,"is found in array ")
	}else{
		fmt.Println(value,"is not found in array")
	}
}



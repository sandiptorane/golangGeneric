package main

import "fmt"

func mergesort(arr []int) []int{
	length := len(arr)
	if  length== 1{
		return arr
	}
	middle:= length/2
	left := make([]int, middle)
	right := make([]int,length-middle)
	for i:=0; i< length;i++{
		if i < middle {
			left[i] = arr[i]
		}else{
			right[i-middle] = arr[i]
		}
	}
  return  merge(mergesort(left),mergesort(right))
}
func merge(left ,right []int) (result []int){
		result = make([]int,len(left)+len(right))
		i:=0
		for len(left)>0 && len(right)>0{
			if left[0]< right[0]{
				result[i]=left[0]
				left = left[1:]
				i++
			}else{
				result[i] = right[0]
				right =  right[1:]
				i++
			}
		}

		for j:=0; j < len(left);j++{
			result[i]=left[j]
			i++
		}
		for j:=0;j<len(right);j++{
			result[i] = right[j]
			i++
		}
		return
}

func main(){
	array := []int{55,36,45,70,2,10}
	fmt.Println("before sort:",array)
	fmt.Println("after sort :",mergesort(array))

}

package main

import "fmt"

type maxheap struct{
	heapArray []int
	size int
	maxsize int
}

func newMaxHeap(maxsize int) *maxheap{
	maxheap := &maxheap{
		heapArray: []int{},
		size : 0,
		maxsize: maxsize,
	}
	return maxheap
}

func (m *maxheap)leaf(index int)bool{
	if index>=(m.size/2) && index <= m.size{
		return true
	}
	return false
}

func (m *maxheap)parent(index int)int{
	return (index-1)/2
}

func (m *maxheap)leftChild(index int)int{
	return (2*index)+1
}

func(m *maxheap)rightChild(index int)int{
	return  (2*index)+2
}

func (m *maxheap)insert(value int)error{
	if m.maxsize==m.size{
		return  fmt.Errorf("insertion not possible heap is full")
	}
	m.heapArray =append(m.heapArray,value)
	m.size++
	m.upHeapify(m.size-1)        //if value is less than parent
	return nil
}

func (m *maxheap)swap(first, second int){
	temp := m.heapArray[first]
	m.heapArray[first] = m.heapArray[second]
	m.heapArray[second]= temp
}

func (m *maxheap)upHeapify(index int){
	for m.heapArray[index]> m.heapArray[m.parent(index)]{
		m.swap(index,m.parent(index))
		index = m.parent(index)
	}
}

func (m *maxheap)downHeapify(current int){
	if m.leaf(current){
		return
	}
	largest := current
	leftChildIndex := m.leftChild(current)
	rightChildIndex := m.rightChild(current)
	//If current is smallest then return
	if leftChildIndex <  m.size && m.heapArray[leftChildIndex] > m.heapArray[largest]{
		largest= leftChildIndex
	}
	if rightChildIndex < m.size && m.heapArray[rightChildIndex] > m.heapArray[largest]{
		largest = rightChildIndex
	}
	if largest !=current{
		m.swap(largest,current)
		m.downHeapify(largest)
	}

	return
}

func (m *maxheap)buildMaxHeap(){
	for i:= ((m.size/2)-1); i>=0; i--{   // lastnode of previous of last  level is (m.size/2 -1)
		m.downHeapify(i)
	}
}

func (m *maxheap)remove() int{
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size-1]   //replace  root by last element
	m.heapArray= m.heapArray[:(m.size-1)]    //delete last array element
	m.size--
	m.downHeapify(0)
	return top
}
func main(){
	arr :=  []int{6, 5, 3, 7, 2, 8}
	maxheap := newMaxHeap(len(arr))
	for i:=0;i<len(arr);i++{
		maxheap.insert(arr[i])
	}

	maxheap.buildMaxHeap()
	fmt.Println("Removed element mininum root")
	for i := 0; i < len(arr); i++ {
		fmt.Println(maxheap.remove())
	}
}



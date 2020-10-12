package main

import "fmt"

type minheap struct{
	minheapArray []int
	size int
	maxsize int
}

func newMinHeap(maxsize int) *minheap{
	minheap := &minheap{
		minheapArray: []int{},
		size : 0,
		maxsize: maxsize,
	}
	return minheap
}

func (m *minheap)leaf(index int)bool{
	if index>=(m.size/2) && index <= m.size{
		return true
	}
	return false
}

func (m *minheap)parent(index int)int{
	return (index-1)/2
}

func (m *minheap)leftChild(index int)int{
	return (2*index)+1
}

func(m *minheap)rightChild(index int)int{
	return  (2*index)+2
}

func (m *minheap)insert(value int)error{
	if m.maxsize==m.size{
		return  fmt.Errorf("insertion not possible heap is full")
	}
	m.minheapArray =append(m.minheapArray,value)
	m.size++
	m.upHeapify(m.size-1)        //if value is less than parent
	return nil
}

func (m *minheap)swap(first, second int){
	temp := m.minheapArray[first]
	m.minheapArray[first] = m.minheapArray[second]
	m.minheapArray[second]= temp
}

func (m *minheap)upHeapify(index int){
	for m.minheapArray[index]< m.minheapArray[m.parent(index)]{
		m.swap(index,m.parent(index))
		index = m.parent(index)
	}
}

func (m *minheap)downHeapify(current int){
	if m.leaf(current){
		return
	}
	smallest := current
	leftChildIndex := m.leftChild(current)
	rightChildIndex := m.rightChild(current)
	//If current is smallest then return
	if leftChildIndex <  m.size && m.minheapArray[leftChildIndex] < m.minheapArray[smallest]{
		smallest = leftChildIndex
	}
	if rightChildIndex < m.size && m.minheapArray[rightChildIndex]< m.minheapArray[smallest]{
		smallest = rightChildIndex
	}
	if smallest !=current{
		m.swap(smallest,current)
		m.downHeapify(smallest)
	}

	return
}

func (m *minheap)buildMinHeap(){
	for i:= ((m.size/2)-1); i>=0; i--{   // lastnode of previous of last  level is (m.size/2 -1)
		m.downHeapify(i)
	}
}

func (m *minheap)remove() int{
	top := m.minheapArray[0]
	m.minheapArray[0] = m.minheapArray[m.size-1]   //replace  root by last element
	m.minheapArray= m.minheapArray[:(m.size-1)]    //delete last array element
	m.size--
	m.downHeapify(0)
	return top
}
func main(){
       arr :=  []int{6, 5, 3, 7, 2, 8}
       minheap := newMinHeap(len(arr))
       for i:=0;i<len(arr);i++{
       	 minheap.insert(arr[i])
	   }

	   minheap.buildMinHeap()
       fmt.Println("Removed element mininum root")
       for i := 0; i < len(arr); i++ {
			fmt.Println(minheap.remove())
       }
		fmt.Scanln()
}



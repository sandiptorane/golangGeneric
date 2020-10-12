//ascending oreder array sorting

package main

import "fmt"

type maxheap struct {
	heapArray []int
}

func newMaxHeap(arr []int) *maxheap{
	maxheap := &maxheap{
		heapArray: arr,
	}
	return maxheap
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

func (m *maxheap)swap(first, second int){
	temp := m.heapArray[first]
	m.heapArray[first] = m.heapArray[second]
	m.heapArray[second]= temp
}

func (m *maxheap)downHeapify(current int,size int){
	largest := current
	leftChildIndex := m.leftChild(current)
	rightChildIndex := m.rightChild(current)
	//If current is smallest then return
	if leftChildIndex <  size && m.heapArray[leftChildIndex] > m.heapArray[largest]{
		largest= leftChildIndex
	}
	if rightChildIndex < size && m.heapArray[rightChildIndex] > m.heapArray[largest]{
		largest = rightChildIndex
	}
	if largest !=current{
		m.swap(largest,current)
		m.downHeapify(largest,size)
	}

	return
}

func (m *maxheap)buildMaxHeap(size int){
	for i:=((size/2)-1); i>=0; i--{   // lastnode of previous of last  level is (m.size/2 -1)
		m.downHeapify(i,size)
	}
}

func (m *maxheap) sort(size int) {
	m.buildMaxHeap(size)
	for i := size - 1; i > 0; i-- {
		// Move current root to end
		m.swap(0, i)
		m.downHeapify(0, i)
	}
}

func (m *maxheap) print() {
	for _, val := range m.heapArray {
		fmt.Println(val)
	}
}

func main(){
	inputArray := []int{6, 5, 3, 7, 2, 8, -1}
	Heap := newMaxHeap(inputArray)
	Heap.sort(len(inputArray))
	Heap.print()
}
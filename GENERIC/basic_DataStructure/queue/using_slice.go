package main

import (
	"fmt"
	"sync"
)

type customQueue struct {
	queue []string
	lock sync.RWMutex
}

func (c *customQueue)enqueue(value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.queue = append(c.queue,value)
}

func (c *customQueue)dequeue() error{
	c.lock.Lock()
	defer c.lock.Unlock()
	if len(c.queue) >0{
		c.queue = c.queue[1:]       //dequeue front
		return nil
	}
	return fmt.Errorf("Error : Queue is empty")
}

func (c *customQueue)front() (string,error){
	c.lock.Lock()
	defer c.lock.Unlock()
	if len(c.queue) > 0{
		return c.queue[0],nil
	}
	return "",fmt.Errorf("peep Error : Queue is empty")
}

func (c *customQueue)size() int{
	return len(c.queue)
}

func (c *customQueue)Empty() bool{
	return len(c.queue)==0
}


func main(){
	myQueue := &customQueue{
		queue : make([]string,0),    //new queue created
	}
	if myQueue.Empty(){
		fmt.Println("Queue is empty")
	}
	fmt.Println("Enqueue : A")
	myQueue.enqueue("A")
	fmt.Println("Enqueue : B")
	myQueue.enqueue("B")
	fmt.Println("Enqueue : C")
	myQueue.enqueue("C")

	fmt.Println("Size of queue :",myQueue.size())

	for len(myQueue.queue)>0{
		frontvalue ,_ := myQueue.front()
		fmt.Println("front of queue :",frontvalue)
		fmt.Println("Dequeue :",frontvalue)
		myQueue.dequeue()
	}

	if myQueue.Empty(){
		fmt.Println("After dequeue queue is empty")   //
	}
}


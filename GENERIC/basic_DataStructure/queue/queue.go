package main

import(
	"container/list"    //this package used for queue
	"fmt"

)

type customQueue struct {
	queue *list.List
}

func (c *customQueue)enqueue(value string) {
	 c.queue.PushBack(value)
}

func (c *customQueue)dequeue() error{
	if c.queue.Len() >0{
		element := c.queue.Front()
		c.queue.Remove(element)
		return nil
	}
	return fmt.Errorf("Error : Queue is empty")
}

func (c *customQueue)front() (string,error){
	if c.queue.Len() > 0{
		if val,ok := c.queue.Front().Value.(string); ok{
			return val,nil
		}
	    return "",fmt.Errorf("peep Error: Queue datatype is incorrect")
	}
	return "",fmt.Errorf("peep Error : Queue is empty")
}

func (c *customQueue)size() int{
	return c.queue.Len()
}

func (c *customQueue)Empty() bool{
	return c.queue.Len()==0
}


func main(){
	myQueue := &customQueue{
		queue: list.New(),    //new queue created
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

	for myQueue.queue.Len()>0{
		frontvalue ,_ := myQueue.front()
		fmt.Println("front of queue :",frontvalue)
		fmt.Println("Dequeue :",frontvalue)
		myQueue.dequeue()
	}

	if myQueue.Empty(){
		fmt.Println("After dequeue queue is empty")   //
	}
}
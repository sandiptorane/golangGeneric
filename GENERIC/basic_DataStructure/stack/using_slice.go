package main

import (
	"fmt"
	"sync"
)

type customStack struct {
	stack []string
	lock sync.RWMutex
}



func (c *customStack)Push(value string) {
c.lock.Lock()
defer c.lock.Unlock()
c.stack = append(c.stack,value)
}

func (c *customStack)Pop() error{
	c.lock.Lock()
	defer c.lock.Unlock()
	length := len(c.stack)
	if  length>0{
		c.stack = c.stack[:length-1]       //pop top
		return nil
	}
	return fmt.Errorf("Error : Stack is empty")
}

func (c *customStack)Front() (string,error){
	c.lock.Lock()
	defer c.lock.Unlock()
	length :=len(c.stack)
	if length > 0{
		return c.stack[length-1],nil
	}
	return "",fmt.Errorf("peep Error : Stack is empty")
}
func (c *customStack)size() int{
	return len(c.stack)
}

func (c *customStack)Empty() bool{
	return len(c.stack)==0
}

func main(){
	myStack :=  &customStack{
		stack: make([]string,0),
	}
	if myStack.Empty(){
		fmt.Println("Stack is empty")
	}
	fmt.Println("Push : A")
	myStack.Push("A")
	fmt.Println("Push : B")
	myStack.Push("B")
	fmt.Println("Push  : C")
	myStack.Push("C")

	fmt.Println("Size of queue :",myStack.size())
	fmt.Println("Stack :",(myStack.stack))

	for len(myStack.stack)>0{
		frontvalue ,_ := myStack.Front()
		fmt.Println("Top of stack :",frontvalue)
		fmt.Println("Pop :",frontvalue)
		myStack.Pop()
	}

}
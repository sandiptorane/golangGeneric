package main

import (
	"container/list"
	"fmt"
)

type customStack struct {
	stack *list.List
}

func (c *customStack)Push(value string) {
	c.stack.PushFront(value)
}

func (c *customStack)Pop() error{
	if c.stack.Len() >0 {
		    value := c.stack.Front()
			c.stack.Remove(value)
			return nil

	}
	return  fmt.Errorf("Error : Stack is empty")
}

func (c *customStack)Front() (string,error){
	if c.stack.Len() > 0{
		if val,ok := c.stack.Front().Value.(string); ok{
			return val,nil
		}
		return "",fmt.Errorf("peep Error: Stack datatype is incorrect")
	}
	return "",fmt.Errorf("peep Error : Stack is empty")
}

func (c *customStack)size() int{
	return c.stack.Len()
}

func (c *customStack)Empty() bool{
	return c.stack.Len()==0
}

func main(){
	myStack :=  &customStack{
		stack: list.New(),
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
	fmt.Println("Stack :",*(myStack.stack))

	for myStack.stack.Len()>0{
		frontvalue ,_ := myStack.Front()
		fmt.Println("Top of stack :",frontvalue)
		fmt.Println("Pop :",frontvalue)
		myStack.Pop()
	}

}
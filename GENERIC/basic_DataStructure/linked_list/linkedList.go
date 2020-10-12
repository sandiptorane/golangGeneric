package main

import "fmt"

type node struct{
	data int
	next *node
}

type singleList struct {
	length int
	head *node
}

func initList() *singleList{
	return &singleList{}
}

func (s *singleList)AddFront(data int){
	tmpnode := &node{
		data : data,
	}
	if s.head == nil{
		s.head=tmpnode
	}else{
		tmpnode.next = s.head
		s.head =tmpnode
	}
	s.length++
	return
}

func (s *singleList)AddBack(data int){
	tmpnode := &node{
		data :data,
	}
	if s.head == nil{
		s.head = tmpnode
	}else{
		current := s.head
		for current.next!=nil{
			current = current.next
		}
		current.next = tmpnode
	}
	s.length++
	return
}

func (s *singleList)RemoveFront() error{
	if s.head == nil{
		return fmt.Errorf("List is empty")
	}else{
		s.head = s.head.next
	}
	s.length--
	return nil
}

func (s *singleList)RemoveBack() error {
	if s.head == nil {
		return fmt.Errorf("RemoveBack : List is empty")
	}
	if s.head.next == nil {
			s.head = nil
	} else{
			current := s.head
			for current.next.next != nil {
				current = current.next
			}
			current.next = nil
	}
	s.length--
	return nil
}

func (s *singleList)Traverse() error {
	if s.head == nil{
		return  fmt.Errorf("Traverse error: List is empty")
	}
	current := s.head
	for current!=nil{
		fmt.Println(current.data)
		current = current.next
	}
	return nil
}

func (s *singleList)Front() (int,error){
	if s.head == nil{
		return 0,fmt.Errorf("Front error: List is empty")
	}
	return s.head.data,nil
}

func (s *singleList)size() int{
	return s.length
}

func main(){
	singleList := initList()
	fmt.Println("Add Front : 11")
	singleList.AddFront(11)
	fmt.Println("Add Front : 22")
	singleList.AddFront(22)
	fmt.Println("Add Back : 33")
	singleList.AddBack(33)

	fmt.Println("Size of list",singleList.size())
	fmt.Println("Traversed list")
	err := singleList.Traverse()
	if  err!=nil{
		fmt.Println(err.Error())
	}

	fmt.Printf("RemoveFront\n")
	err = singleList.RemoveFront()
	if err != nil {
		fmt.Printf("RemoveFront Error: %s\n", err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	fmt.Println("Traversed list")
	err = singleList.Traverse()
	if  err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println("size of list :",singleList.length)

}
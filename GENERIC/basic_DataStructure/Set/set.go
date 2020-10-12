package main

import "fmt"

type customSet struct{
	container map[string]struct{}
}

//makeSet initialize the set
func makeSet() *customSet{
	return &customSet{
		container: make(map[string]struct{}),
	}
}

func (c *customSet)Exists(key string) bool{
	_,exists := c.container[key]
	return exists
}
func (c *customSet)Add(key string){
	c.container[key] = struct{}{}
}

func (c *customSet)Remove(key string)error{
	if !c.Exists(key){
		return fmt.Errorf("Error : key does not exists")
	}
	delete(c.container,key)
	return nil
}

func (c *customSet)Size() int{
	return len(c.container)
}

func  main(){
	set := makeSet()
	fmt.Println("Add : A")
	set.Add("A")
	fmt.Println("Add : B")
	set.Add("B")

	fmt.Println("size of set :",set.Size())
	fmt.Println("A Exists? :",set.Exists("A"))
	fmt.Println("B Exists? :",set.Exists("B"))
	fmt.Println("C Exists? :",set.Exists("C"))
	fmt.Println("Remove B")
	set.Remove("B")
	fmt.Println("B Exists? :",set.Exists("B"))

}
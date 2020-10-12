package main

import "fmt"

type bstnode struct {
	data int
	left *bstnode
	right *bstnode
}

type bst struct {
	root *bstnode
}

func (b *bst)reset(){
	b.root = nil
}

func (b *bst)insert(value int){
	b.insertRec(b.root,value)
}

func (b *bst)insertRec(node *bstnode,value int) *bstnode{
	if b.root ==nil {
		b.root = &bstnode{data : value,}
		return b.root
	}
	if node ==nil{
		node = &bstnode{data : value}
		return  node
	}
	if value <= node.data{
		node.left= b.insertRec(node.left,value)
	}
	if value > node.data{
		node.right=b.insertRec(node.right,value)
	}
	return node
}

func (b *bst)find(value int)error{
   node  := b.findRec(b.root,value)
   if node == nil{
   	return fmt.Errorf("value %d not  found",value)
   }
	return nil
}

func (b *bst)findRec(node *bstnode,value int) *bstnode{
	if node==nil {
		return nil
	}
	if value == node.data{
		return node
	}
	if value < node.data{
		return b.findRec(node.left,value)
	}
	return b.findRec(node.right,value)
}

func (b *bst)inorder(){
	b.inorderRec(b.root)
}

func (b *bst)inorderRec(node *bstnode){
	if node != nil{
		b.inorderRec(node.left)
		fmt.Println(node.data)
		b.inorderRec(node.right)
	}
}
func main(){
    bst := &bst{}
    values := []int{2,5,7,-1,-1,5,5}
    for _,value := range values{
    	bst.insert(value)
	}
	fmt.Println("Inorder :")
    bst.inorder()
    fmt.Println("finding :")
    for i:=0;i<6;i++{
    	err := bst.find(i)
    	if err==nil{
    		fmt.Println(i,"is found in the Tree")
		}else{
			fmt.Println(i,"is not found in the Tree")
		}
	}

	bst.reset()
    fmt.Println("BST reset done")

}

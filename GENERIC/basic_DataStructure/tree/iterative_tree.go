package main

import "fmt"

type bstnode struct {
	value int
	left  *bstnode
	right *bstnode
}

type bst struct {
	root *bstnode
}

func initList() *bst {
	return &bst{}
}

func (b *bst) reset() {
	b.root = nil
}

func (b *bst) insert(value int) {
	b.insertRec(b.root, value)
}

func (b *bst) insertRec(node *bstnode, value int) {
	if b.root == nil {
		b.root = &bstnode{
			value: value,
		}
	}
	if node == nil {
		return
	}
	//Find the terminalNode where to insert the new node
	var terminalNode *bstnode
	for node != nil {
		terminalNode = node
		if value <= node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
	if value <= terminalNode.value {
		terminalNode.left = &bstnode{value: value}
	} else {
		terminalNode.right = &bstnode{value: value}
	}
	return
}

func (b *bst) find(value int) error {
	node := b.findRec(b.root, value)
	if node == nil {
		return fmt.Errorf("Value: %d not found in tree", value)
	}
	return nil
}

func (b *bst) findRec(node *bstnode, value int) *bstnode {
	if node == nil {
		return nil
	}
	if node.value == value {
		return b.root
	}
	if value < node.value {
		return b.findRec(node.left, value)
	}
	return b.findRec(node.right, value)
}

func (b *bst) inorder() {
	b.inorderRec(b.root)
}

func (b *bst) inorderRec(node *bstnode) {
	if node != nil {
		b.inorderRec(node.left)
		fmt.Println(node.value)
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

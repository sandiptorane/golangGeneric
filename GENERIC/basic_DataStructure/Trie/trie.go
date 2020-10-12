package main

import (
	"fmt"
)

const ALPHABET_SIZE = 26

type trieNode struct {
	childrens [ALPHABET_SIZE]*trieNode
	wordEnd bool
}

type trie struct {
	root *trieNode
}

func initTrie() *trie{
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie)insert(word string){
	wordLength := len(word)
	current := t.root
	for i:=0; i<wordLength;i++{
		index:= word[i]-'a'
		if current.childrens[index]==nil{
			current.childrens[index]= &trieNode{}
		}
		current = current.childrens[index]
	}
	current.wordEnd =true
}

func (t *trie)find(word string)bool{
	wordLength := len(word)
	current := t.root
	for i:=0;i<wordLength;i++{
		index := word[i]-'a'
		if current.childrens[index]==nil{
			return false  //if particluar char is not present at childrens[index]
		}
		current =current.childrens[index] //go down to childrens
	}
	if current.wordEnd{
		return true
	}

	return false
}

func main(){
	trie :=initTrie()
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	for i:=0;i<len(words);i++{
		trie.insert(words[i])
	}
	wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}
	for i := 0; i < len(wordsToFind); i++ {
		found := trie.find(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
		} else {
			fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
		}
	}
}
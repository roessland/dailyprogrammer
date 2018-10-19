package main

import "fmt"
import "log"
import (
	"io/ioutil"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	Word     string
	Children map[string]*Node
}

func NewNode(word string) *Node {
	node := Node{
		Word:     word,
		Children: make(map[string]*Node),
	}
	return &node
}

type Graph map[string]*Node

func NewGraph() Graph {
	return make(Graph)
}

func (g Graph) GetNode(word string) *Node {
	node, ok := g[word]
	if ok {
		return node
	} else {
		g[word] = NewNode(word)
		return g[word]
	}
}

func (g Graph) AddChildren(word string, wlist map[string]bool) {
	node := g.GetNode(word)
	for i := 0; i < len(word); i++ {
		child := word[0:i] + word[i+1:len(word)]
		if !wlist[child] {
			continue
		}
		childNode := g.GetNode(child)
		node.Children[child] = childNode
	}
}

func (root *Node) PrintTree() {
	var printTree func(*Node, int)
	printTree = func(node *Node, indent int) {
		fmt.Println(strings.Repeat(" ", indent), node.Word)
		for _, child := range node.Children {
			printTree(child, indent+1)
		}
	}
	printTree(root, 0)
}

func (root *Node) Length() int {
	var getDepth func(*Node, int) int
	getDepth = func(node *Node, depth int) int {
		if len(node.Children) == 0 {
			return depth
		}
		maxDepth := 0
		for _, child := range node.Children {
			maxDepth = max(maxDepth, getDepth(child, 0))
		}
		return maxDepth + 1
	}
	return 1 + getDepth(root, 0)
}

func ReadWordList() map[string]bool {
	wlist := make(map[string]bool)
	buf, err := ioutil.ReadFile("enable1.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, word := range strings.Split(string(buf), "\n") {
		if word == "" {
			continue
		}
		wlist[word] = true
	}
	return wlist
}

func BuildGraph(wlist map[string]bool) Graph {
	g := NewGraph()
	for word, _ := range wlist {
		_ = g.GetNode(word)
		g.AddChildren(word, wlist)
	}
	return g
}

func main() {
	fmt.Println("vim-go")
	wlist := ReadWordList()
	graph := BuildGraph(wlist)
	fmt.Printf("%v\n", graph["gnash"])
	graph["gnash"].PrintTree()
	fmt.Println(4, graph["gnash"].Length())
	fmt.Println(9, graph["princesses"].Length())
	fmt.Println(5, graph["turntables"].Length())
	fmt.Println(1, graph["implosive"].Length())
	fmt.Println(2, graph["programmer"].Length())
	for _, node := range graph {
		l := node.Length()
		if l > 8 {
			fmt.Println(node.Word, node.Length())
		}
	}
	graph["complecting"].PrintTree()
}

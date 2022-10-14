package main

import (
	"fmt"
	"math/rand"
	"time"
)

type valueType int

type Tree struct {
	Left  *Tree
	Value valueType
	Right *Tree
}

func createTree(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, valueType(temp))
	}
	return t
}

func insert(t *Tree, v valueType) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func preOrderTraverse(t *Tree) {
	if t == nil {
		return
	}
	fmt.Print(t.Value, " ")
	preOrderTraverse(t.Left)
	preOrderTraverse(t.Right)
}

func midOrderTraverse(t *Tree) {
	if t == nil {
		return
	}
	midOrderTraverse(t.Left)
	fmt.Print(t.Value, " ")
	midOrderTraverse(t.Right)
}

func aftOrderTraverse(t *Tree) {
	if t == nil {
		return
	}
	aftOrderTraverse(t.Left)
	aftOrderTraverse(t.Right)
	fmt.Print(t.Value, " ")
}

func main() {
	tree := createTree(10)
	fmt.Println("PreOrder:")
	preOrderTraverse(tree)
	fmt.Println()
	fmt.Println("MidOrder:")
	midOrderTraverse(tree)
	fmt.Println()
	fmt.Println("AftOrder:")
	aftOrderTraverse(tree)
	fmt.Println()

	tree = insert(tree, -10)
	tree = insert(tree, -2)
	fmt.Println("PreOrder:")
	preOrderTraverse(tree)
	fmt.Println()
	fmt.Println("MidOrder:")
	midOrderTraverse(tree)
	fmt.Println()
	fmt.Println("AftOrder:")
	aftOrderTraverse(tree)
	fmt.Println()
}

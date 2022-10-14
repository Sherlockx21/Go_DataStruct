package main

import "fmt"

type valueType int

type Node struct {
	Value valueType
	Pre   *Node
	Next  *Node
}

var root = new(Node)

func addNode(t *Node, v valueType) int {
	if root == nil {
		root = &Node{v, nil, nil}
		return 0
	}

	if t.Value == v {
		fmt.Println("Node already exist", v)
		return -1
	}
}

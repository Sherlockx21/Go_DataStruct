package main

import "fmt"

const SIZE = 15

type valueType int
type keyType int

type node struct {
	Value valueType
	Next  *node
}

type hashTable struct {
	Table map[keyType]*node
	Size  int
}

func hashFunc(v valueType, size int) keyType {
	//doing something to match value with the key
	return keyType(int(v) % size)
}
func insert(hash *hashTable, v valueType) keyType {
	index := hashFunc(v, hash.Size)
	elem := node{Value: v, Next: hash.Table[index]}
	hash.Table[index] = &elem
	return index
}

func traverse(hash *hashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%v", t.Value)
				if t.Next != nil {
					fmt.Printf("->")
				}
				t = t.Next
			}
		}
		fmt.Println()
	}
}

func lookup(hash *hashTable, v valueType) bool {
	index := hashFunc(v, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == v {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func main() {
	table := make(map[keyType]*node, SIZE)
	hash := &hashTable{Table: table, Size: SIZE}
	fmt.Println("Size of table: ", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, valueType(i))
	}
	traverse(hash)

	fmt.Println(lookup(hash, 120))
	fmt.Println(lookup(hash, 15))
	fmt.Println(lookup(hash, 32646))

}

package main

import (
    "fmt"
)

type Node struct {
    values map[rune]*Node
    isLast bool
}

func (n *Node) add(value rune) *Node {
    newNode := &Node{values:nil, isLast:true}
    if (n.values == nil) {
        n.values = map[rune]*Node{value:newNode}
    } else {
        if (n.values[value] == nil) {
            n.values[value] = newNode
        } else {
            return n.values[value]
        }
    }
    n.isLast = false
    return newNode
}

func (n *Node) traverse(path string) {
    if !n.isLast {
        for key := range n.values {
            path = path + fmt.Sprintf("%c --> ", key)
            n.values[key].traverse(path)
            path = path + fmt.Sprintf("\b\b\b\b\b\b")
        }
    } else {
        fmt.Println(path,"\b\b\b\b(Final Node)")
        return
    }
}

func createTree(rootNode *Node, char_lst []rune) {
    currNode := rootNode
    for _, ele := range char_lst {
        fmt.Printf("Adding %c \n", ele)
        currNode = currNode.add(ele)
    }   
}

func main() {
    fmt.Println("start of program")
    char_lst_1 := []rune{'a', 'b', 'c'}
    char_lst_2 := []rune{'a', 'b', 'd'}
    char_lst_3 := []rune{'a', 'b', 'c', 'd'}
    char_lst_4 := []rune{'k', 'b', 'l'}
    char_lst_5 := []rune{'a', 'h', 'm'}
    rootNode := &Node{values:nil, isLast:true}
    createTree(rootNode, char_lst_1)
    fmt.Println("-------------------------------------------------")
    createTree(rootNode, char_lst_2)
    fmt.Println("-------------------------------------------------")
    createTree(rootNode, char_lst_3)
    fmt.Println("-------------------------------------------------")
    createTree(rootNode, char_lst_4)
    fmt.Println("-------------------------------------------------")
    createTree(rootNode, char_lst_5)
    fmt.Println("-------------------------------------------------")
    fmt.Println("root node", rootNode)
    rootNode.traverse("start --> ")
}

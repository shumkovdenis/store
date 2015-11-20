package main

//import "reflect"
//
//type Node struct {
//	Properties map[string]string
//	Nodes      []*Node
//	Count      int
//}
//
//type Tree struct {
//	nodes []*Node
//}
//
//func newNode(properties map[string]string) *Node {
//	return &Node{properties, []*Node{}, 1}
//}
//
//func (node *Node) Check(properties map[string]string) bool {
//	for v, k := range properties {
//		if node.Properties[k] != v {
//			return false
//		}
//	}
//	return true
//}
//
//func NewTree() *Tree {
//	return &Tree{[]*Node{}}
//}
//
//func (tree *Tree) Save(track *Track) {
//	var n Node
//	for _, node := range tree.nodes {
//		if node.Check(track.Properties) {
//			n = node
//			break
//		}
//	}
//
//	if n == nil {
//		n = NewNode(track.Properties)
//	}
//
//	n.Nodes = append(n.Nodes, NewNode(track.Properties))
//}
//
//func (tree *Tree) Count(track *Track) int {
//	c := 0
//	return c
//}

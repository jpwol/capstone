package engine

import "fmt"

type NodeId uint32

type Node struct {}

type Graph struct {
	Adjacency map[NodeId][]NodeId
}

func (g Graph) Iterate() {
	for i := range g.Adjacency {
		for _, v := range g.Adjacency[i] {
			fmt.Println(i,v)
		}
	}
}

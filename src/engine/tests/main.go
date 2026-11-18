package main

import "capstone/engine"

func main() {
	g := engine.Graph{
		Adjacency: make(map[engine.NodeId][]engine.NodeId),
	}

	g.Adjacency[1] = []engine.NodeId{2,3}
	g.Adjacency[2] = []engine.NodeId{3}
	g.Adjacency[4] = []engine.NodeId{2}
	g.Iterate()
}

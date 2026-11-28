package engine

import (
	"fmt"
)

type NodeId uint32

type Graph struct {
	nextID       NodeId
	nodes        map[NodeId]*TaskNode
	dependents   map[NodeId][]NodeId
	dependencies map[NodeId][]NodeId

	initialized bool
}

func NewGraph() *Graph {
	return &Graph{
		nextID:       0,
		nodes:        make(map[NodeId]*TaskNode),
		dependents:   make(map[NodeId][]NodeId),
		dependencies: make(map[NodeId][]NodeId),
		initialized: false,
	}
}

func (g *Graph) AddTask(task Task) NodeId {
	id := g.nextID
	g.nextID++

	g.nodes[id] = &TaskNode{
		ID:                    id,
		Task:                  task,
		State:                 Pending,
		remainingDependencies: 0,
	}

	g.dependents[id] = []NodeId{}
	g.dependencies[id] = []NodeId{}

	return id
}

// call this once graph construction is done
func (g *Graph) Initialize() ([]*TaskNode, error) {
	if g.initialized {
		return nil, fmt.Errorf("cannot initialize graph: already initialized")
	}

	if g.HasCycle() {
		return nil, fmt.Errorf("cannot initialize graph: cycle detected")
	}

	var ready []*TaskNode

	for _, node := range g.nodes {
		if node.State == Pending && node.remainingDependencies == 0 {
			node.State = Ready
			ready = append(ready, node)
		}
	}

	g.initialized = true
	return ready, nil
}

type Graph struct {
	nextID     NodeId
	nodes      map[NodeId]TaskNode
	adjacency  map[NodeId][]NodeId
	reverseAdj map[NodeId][]NodeId
}

func InitGraph() Graph {
	return Graph{
		nextID:     0,
		nodes:      make(map[NodeId]TaskNode),
		adjacency:  make(map[NodeId][]NodeId),
		reverseAdj: make(map[NodeId][]NodeId),
	}
}

func (g *Graph) AddTask(task Task) NodeId {
	id := g.nextID
	g.nextID++

	g.nodes[id] = TaskNode{
		ID:   id,
		Task: task,
		remainingDependencies: 0,
	}

	g.adjacency[id] = []NodeId{}
	g.reverseAdj[id] = []NodeId{}

	return id
}

// add cycle detection on from and to instead of checking whole graph
// on every edge creation
func (g *Graph) AddDependency(from, to NodeId) {
	g.adjacency[from] = append(g.adjacency[from], to)
	g.reverseAdj[to] = append(g.reverseAdj[to], from)

	if entry, ok := g.nodes[from]; ok {
		entry.remainingDependencies++

		g.nodes[from] = entry
	}
}

func (g Graph) GetTask(id NodeId) Task {
	return g.nodes[id].Task
}

func (g Graph) GetNeighbors(id NodeId) []NodeId {
	return []NodeId{}
}

// inspect functions are currently public, should probably be private
func (g Graph) Inspect() {
	fmt.Println("Adjacency List")
	fmt.Println("--------------")
	for i := range g.adjacency {
		fmt.Printf("%d:  { ", i)
		for _, v := range g.adjacency[i] {
			fmt.Printf("%d ", v)
		}
		fmt.Println("}")
	}
}

func (g Graph) InspectReverse() {
	fmt.Println("Reverse Adjacency List")
	fmt.Println("----------------------")
	for i := range g.reverseAdj {
		fmt.Printf("%d:  { ", i)
		for _, v := range g.reverseAdj[i] {
			fmt.Printf("%d ", v)
		}
		fmt.Println("}")
	}
}

func (g *Graph) HasCycle() bool {
	const (
		unvisited = iota
		visiting
		visited
	)

	state := make(map[NodeId]int)

	var dfs func(NodeId) bool

	dfs = func(id NodeId) bool {
		state[id] = visiting

		for _, neighbor := range g.adjacency[id] {
			if state[neighbor] == visiting {
				return true
			}

			if state[neighbor] == unvisited {
				if dfs(neighbor) {
					return true
				}
			}
		}

		state[id] = visited
		return false
	}

	for id := range g.nodes {
		if state[id] == unvisited {
			if dfs(id) {
				return true
			}
		}
	}

	return false
}

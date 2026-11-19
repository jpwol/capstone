package engine

import "fmt"

type NodeId uint32

type Node struct {
	ID   NodeId
	Task Task
	// could track execution state
	// and more if needed

	// should include reference counter here
}

type Task struct {
	Name string
	// other data will go here
	// such as `cmd`, etc
}

type Graph struct {
	nextID     NodeId
	nodes      map[NodeId]Node
	adjacency  map[NodeId][]NodeId
	reverseAdj map[NodeId][]NodeId
}

func InitGraph() Graph {
	return Graph{
		nextID:     0,
		nodes:      make(map[NodeId]Node),
		adjacency:  make(map[NodeId][]NodeId),
		reverseAdj: make(map[NodeId][]NodeId),
	}
}

func (g *Graph) AddTask(task Task) NodeId {
	id := g.nextID
	g.nextID++

	g.nodes[id] = Node{
		ID:   id,
		Task: task,
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
}

func (g Graph) GetTask(id NodeId) Task {

	return Task{}
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

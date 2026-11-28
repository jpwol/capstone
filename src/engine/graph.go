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

// TODO: add check to make sure same from and to are not added more than once
func (g *Graph) AddDependency(from, to NodeId) error {
	if _, ok := g.nodes[from]; !ok {
		return fmt.Errorf("node %d does not exist", from)
	}

	_, ok := g.nodes[to]
	if !ok {
		return fmt.Errorf("node %d does not exist", to)
	}

	g.dependents[from] = append(g.dependents[from], to)
	g.dependencies[to] = append(g.dependencies[to], from)

	g.nodes[to].remainingDependencies++

	return nil
}

func (g *Graph) StartTask(id NodeId) error {
	node, ok := g.nodes[id]
	if !ok {
		return fmt.Errorf("node %d does not exist", id)
	}

	if node.State != Ready {
		return fmt.Errorf("node %d is not ready", id)
	}

	node.State = Running
	return nil
}

func (g *Graph) CompleteTask(id NodeId) ([]*TaskNode, error) {
	node, ok := g.nodes[id]
	if !ok {
		return nil, fmt.Errorf("node %d does not exist", id)
	}

	if node.State != Running {
		return nil, fmt.Errorf("node %d is not running", id)
	}

	node.State = Completed

	var ready []*TaskNode

	for _, dependentID := range g.dependents[id] {
		dependent := g.nodes[dependentID]

		dependent.remainingDependencies--

		if dependent.remainingDependencies == 0 && dependent.State == Pending {
			dependent.State = Ready
			ready = append(ready, dependent)
		}
	}

	return ready, nil
}

func (g *Graph) FailTask(id NodeId, err error) error {
	node, ok := g.nodes[id]
	if !ok {
		return fmt.Errorf("node %d does not exist", id)
	}

	if node.State != Running {
		return fmt.Errorf("node %d is not running", id)
	}

	node.State = Failed
	return nil
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
	for i := range g.dependents {
		fmt.Printf("%d:  { ", i)
		for _, v := range g.dependents[i] {
			fmt.Printf("%d ", v)
		}
		fmt.Println("}")
	}
}

func (g Graph) InspectReverse() {
	fmt.Println("Reverse Adjacency List")
	fmt.Println("----------------------")
	for i := range g.dependencies {
		fmt.Printf("%d:  { ", i)
		for _, v := range g.dependencies[i] {
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

		for _, neighbor := range g.dependents[id] {
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

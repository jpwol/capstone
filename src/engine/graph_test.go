package engine

import "testing"

func TestAddTask(t *testing.T) {
	g := InitGraph()

	one := g.AddTask(Task{Name: "one"})
	two := g.AddTask(Task{Name: "two"})

	g.AddDependency(one, two)

	if len(g.adjacency[one]) != 1 {
		t.Fatal("Expected one outgoing edge")
	}

	if g.adjacency[one][0] != two {
		t.Fatal("Incorrect outgoing edge")
	}

	if len(g.reverseAdj[two]) != 1 {
		t.Fatal("Expected one reverse edge")
	}

	if g.reverseAdj[two][0] != one {
		t.Fatal("Incorrect reverse edge")
	}

	three := g.AddTask(Task{Name: "three"})

	g.AddDependency(one, three)
	g.AddDependency(two, three)

	if len(g.adjacency[one]) != 2 {
		t.Fatal("Expected two outgoing edges")
	}

	if g.adjacency[one][1] != three {
		t.Fatal("Incorrect outgoing edge")
	}
	
	if len(g.adjacency[two]) != 1 {
		t.Fatal("Expected one outgoing edge")
	}

	if g.adjacency[two][0] != three {
		t.Fatal("Incorrect outgoing edge")
	}

	if len(g.reverseAdj[three]) != 2 {
		t.Fatal("Expected one reverse edge")
	}

	if g.reverseAdj[three][0] != one {
		t.Fatal("Incorrect reverse edge")
	}

	g.Inspect()
	g.InspectReverse()

	if g.HasCycle() != false {
		t.Fatal("Cycle Detected")
	}
}


package engine

import "testing"

func TestAddTask(t *testing.T) {
	g := NewGraph()

	one := g.AddTask(Task{Name: "one"})
	two := g.AddTask(Task{Name: "two"})

	g.AddDependency(one, two)

	if len(g.dependents[one]) != 1 {
		t.Fatal("Expected one outgoing edge")
	}

	if g.dependents[one][0] != two {
		t.Fatal("Incorrect outgoing edge")
	}

	if len(g.dependencies[two]) != 1 {
		t.Fatal("Expected one reverse edge")
	}

	if g.dependencies[two][0] != one {
		t.Fatal("Incorrect reverse edge")
	}

	three := g.AddTask(Task{Name: "three"})

	g.AddDependency(one, three)
	g.AddDependency(two, three)

	if len(g.dependents[one]) != 2 {
		t.Fatal("Expected two outgoing edges")
	}

	if g.dependents[one][1] != three {
		t.Fatal("Incorrect outgoing edge")
	}
	
	if len(g.dependents[two]) != 1 {
		t.Fatal("Expected one outgoing edge")
	}

	if g.dependents[two][0] != three {
		t.Fatal("Incorrect outgoing edge")
	}

	if len(g.dependencies[three]) != 2 {
		t.Fatal("Expected one reverse edge")
	}

	if g.dependencies[three][0] != one {
		t.Fatal("Incorrect reverse edge")
	}

	g.Inspect()
	g.InspectReverse()
}

func TestCheckCycles(t *testing.T) {
	g := NewGraph()

	one := g.AddTask(Task{Name: "one"})
	two := g.AddTask(Task{Name: "two"})

	g.AddDependency(one, two)

	if g.HasCycle() != false {
		t.Fatal("Cycle Detected")
	}

	g.AddDependency(two, one)

	if g.HasCycle() != true {
		t.Fatal("Cycle Not Detected")
	}
}

func TestRemainingCounter(t *testing.T) {
	g := NewGraph()

	one := g.AddTask(Task{Name: "one"})
	two := g.AddTask(Task{Name: "two"})

	remaining := g.nodes[two].remainingDependencies
	if remaining != 0 {
		t.Fatal("Dependency Count Should Be 0, Got ", remaining)
	}

	g.AddDependency(one, two)

	remaining = g.nodes[two].remainingDependencies
	if remaining != 1 {
		t.Fatal("Dependency Count Should Be 1, Got ", remaining)
	}
}

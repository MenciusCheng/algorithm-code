package async_queue

import (
	"reflect"
	"testing"
)

/*
TestDAG_case1
A -> B -> C
*/
func TestDAG_case1(t *testing.T) {
	dag := NewDAG()
	dag.AddVertex("A", nil)
	dag.AddVertex("B", nil)
	dag.AddVertex("C", nil)
	dag.AddEdge("A", "B")
	dag.AddEdge("B", "C")

	want := []*Vertex{
		dag.Vertexes["A"],
		dag.Vertexes["B"],
		dag.Vertexes["C"],
	}
	got := dag.BFS()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BFS() = %v, want %v", got, want)
	}

	if got := dag.CanFinish(); got != true {
		t.Errorf("CanFinish() = %v, want %v", got, true)
	}

	// 构成有环图
	dag.AddEdge("C", "A")
	if got := dag.CanFinish(); got != false {
		t.Errorf("CanFinish() = %v, want %v", got, false)
	}
}

/*
TestDAG_case2
A -> B -> (C/D) -> E -> F
*/
func TestDAG_case2(t *testing.T) {
	dag := NewDAG()
	dag.AddVertex("A", nil)
	dag.AddVertex("B", nil)
	dag.AddVertex("C", nil)
	dag.AddVertex("D", nil)
	dag.AddVertex("E", nil)
	dag.AddVertex("F", nil)
	dag.AddEdge("A", "B")
	dag.AddEdge("B", "C")
	dag.AddEdge("B", "D")
	dag.AddEdge("C", "E")
	dag.AddEdge("D", "E")
	dag.AddEdge("E", "F")

	want1 := []*Vertex{
		dag.Vertexes["A"],
		dag.Vertexes["B"],
		dag.Vertexes["C"], // C 和 D 的先后顺序会调换
		dag.Vertexes["D"],
		dag.Vertexes["E"],
		dag.Vertexes["F"],
	}
	want2 := []*Vertex{
		dag.Vertexes["A"],
		dag.Vertexes["B"],
		dag.Vertexes["D"],
		dag.Vertexes["C"],
		dag.Vertexes["E"],
		dag.Vertexes["F"],
	}
	got := dag.BFS()
	if !reflect.DeepEqual(got, want1) && !reflect.DeepEqual(got, want2) {
		t.Errorf("BFS() = %v, want1 %v, want2 %v", got, want1, want2)
	}

	if got := dag.CanFinish(); got != true {
		t.Errorf("CanFinish() = %v, want %v", got, true)
	}

	// 构成有环图
	dag.AddEdge("F", "A")
	if got := dag.CanFinish(); got != false {
		t.Errorf("CanFinish() = %v, want %v", got, false)
	}
}

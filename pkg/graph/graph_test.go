package graph

import (
	"fmt"
	"testing"
)

func TestDfs(t *testing.T) {
	graph := NewGraph()
	graph.AddNode("a")
	graph.AddNode("b")
	graph.AddNode("c")
	graph.AddNode("d")
	graph.AddNode("e")

	graph.AddEdge("a", "b")
	graph.AddEdge("b", "c")
	graph.AddEdge("c", "d")
	graph.AddEdge("a", "d")
	graph.AddEdge("d", "e")
	//graph.AddEdge("e", "b")

	hasCycle := HasCycle(graph)
	fmt.Println(hasCycle)
}

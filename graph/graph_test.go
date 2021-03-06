package graph

import (
	"fmt"
	"testing"
)

//无向图
func TestNewUndirectedGraph(t *testing.T) {
	/**
	0 1,2,3,4
	1 0,4,7,9
	2 0
	3 0,5,6
	4 0,1,5
	5 3,4
	6 3
	7 1,8,10
	8 7
	9 1
	10 7
	*/

	undirectedGraph := &UndirectedGraph{}

	undirectedGraph.NewGraph(11)
	undirectedGraph.AddEdge(0, 1)
	undirectedGraph.AddEdge(0, 2)
	undirectedGraph.AddEdge(0, 3)
	undirectedGraph.AddEdge(0, 4)
	undirectedGraph.AddEdge(1, 4)
	undirectedGraph.AddEdge(1, 7)
	undirectedGraph.AddEdge(1, 9)
	undirectedGraph.AddEdge(3, 5)
	undirectedGraph.AddEdge(3, 6)
	undirectedGraph.AddEdge(4, 5)
	undirectedGraph.AddEdge(7, 8)
	undirectedGraph.AddEdge(7, 10)
	fmt.Println("Bfs:")
	undirectedGraph.graph.Bfs()
	fmt.Println("Dfs:")
	undirectedGraph.graph.Dfs()
}

//有向图
func TestNewDirectedGraph(t *testing.T) {
	directedGraph := &DirectedGraph{}

	directedGraph.NewGraph(11)
	directedGraph.AddEdge(0, 1)
	directedGraph.AddEdge(0, 2)
	directedGraph.AddEdge(0, 3)
	directedGraph.AddEdge(0, 4)
	directedGraph.AddEdge(1, 4)
	directedGraph.AddEdge(1, 7)
	directedGraph.AddEdge(1, 9)
	directedGraph.AddEdge(3, 5)
	directedGraph.AddEdge(3, 6)
	directedGraph.AddEdge(4, 5)
	directedGraph.AddEdge(7, 8)
	directedGraph.AddEdge(7, 10)
	directedGraph.graph.Bfs()

}

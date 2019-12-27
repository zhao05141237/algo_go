package graph

import (
	"algo/linklist/singleLinkedListNoHead"
	"algo/queue/queueArray"
	"fmt"
)

var found bool

type Graph struct {
	v   int                                //顶点个数
	adj []*singleLinkedListNoHead.ListNode //邻接表
}

func (g *Graph) Bfs() {
	for i := 0; i < g.v; i++ {
		for j := i + 1; j < g.v; j++ {
			fmt.Printf("begin:%d,end:%d\n", i, j)
			g.bfsTwoPoint(i, j)
			fmt.Printf("\n")
		}
	}
}

func (g *Graph) bfsTwoPoint(s int, t int) {
	if s == t || !checkVector(s, g) || !checkVector(t, g) {
		return
	}
	visited := make([]bool, g.v)
	prev := make([]int, g.v)
	for index := range prev {
		prev[index] = -1
	}
	queue := queueArray.NewArrayQueue(g.v)
	queue.EnQueue(s)
	visited[s] = true

	for !queue.IsEmpty() {
		w := queue.DeQueue()
		switch w.(type) {
		case int:
			for node := g.adj[w.(int)]; node != nil; node = node.Next {
				q := node.Val
				if !visited[q] {
					prev[q] = w.(int)
					if q == t {
						g.Print(prev, s, t)
						return
					}
					visited[q] = true
					queue.EnQueue(q)
				}
			}
		default:
			return
		}
	}

}

func (g *Graph) Dfs() {
	for i := 0; i < g.v; i++ {
		for j := i + 1; j < g.v; j++ {
			fmt.Printf("begin:%d,end:%d\n", i, j)
			found = false
			g.dfsTwoPoint(i, j)
			fmt.Printf("\n")
		}
	}
}

func (g *Graph) dfsTwoPoint(s int, t int) {
	if s == t || !checkVector(s, g) || !checkVector(t, g) {
		return
	}
	visited := make([]bool, g.v)
	prev := make([]int, g.v)
	for index := range prev {
		prev[index] = -1
	}
	g.recurDfs(s, t, visited, prev)
	g.Print(prev, s, t)
}

func (g *Graph) recurDfs(w int, t int, visited []bool, prev []int) {
	if found == true {
		return
	}
	visited[w] = true
	if w == t {
		found = true
		return
	}

	for node := g.adj[w]; node != nil; node = node.Next {
		q := node.Val
		if !visited[q] {
			prev[q] = w
			g.recurDfs(q, t, visited, prev)
		}
	}

}

func (g *Graph) Print(prev []int, s int, t int) {
	if prev[t] != -1 && t != s {
		g.Print(prev, s, prev[t])
	}
	if t != s {
		fmt.Printf("=>%d", t)
	} else {
		fmt.Printf("%d", t)
	}
}

type Grapher interface {
	AddEdge(s int, t int)
	NewGraph(v int)
}

type DirectedGraph struct {
	graph *Graph
}

func (d *DirectedGraph) AddEdge(s int, t int) {
	if s == t || !checkVector(s, d.graph) || !checkVector(t, d.graph) {
		return
	}
	if !d.graph.adj[s].FindValInList(t) {
		d.graph.adj[s].InsertIntoList(t)
	}
}

func (d *DirectedGraph) NewGraph(v int) {
	d.graph = newGraph(v)
}

type UndirectedGraph struct {
	graph *Graph
}

func (u *UndirectedGraph) AddEdge(s int, t int) {
	if s == t || !checkVector(s, u.graph) || !checkVector(t, u.graph) {
		return
	}
	if !u.graph.adj[s].FindValInList(t) {
		u.graph.adj[s].InsertIntoList(t)
	}
	if !u.graph.adj[t].FindValInList(s) {
		u.graph.adj[t].InsertIntoList(s)
	}
}

func (u *UndirectedGraph) NewGraph(v int) {
	u.graph = newGraph(v)
}

func checkVector(s int, graph *Graph) bool {
	if s < 0 || s >= graph.v {
		return false
	} else {
		return true
	}
}

func newGraph(v int) *Graph {
	graph := &Graph{
		v:   v,
		adj: make([]*singleLinkedListNoHead.ListNode, v),
	}

	for i := 0; i < v; i++ {
		graph.adj[i] = &singleLinkedListNoHead.ListNode{
			Val:  i,
			Next: nil,
		}
	}
	return graph
}

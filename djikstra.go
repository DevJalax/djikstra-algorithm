package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	to, cost int
}

type Graph struct {
	nodes map[int][]Edge
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[int][]Edge)}
}

func (g *Graph) AddEdge(from, to, cost int) {
	g.nodes[from] = append(g.nodes[from], Edge{to, cost})
	g.nodes[to] = append(g.nodes[to], Edge{from, cost})
}

type QueueItem struct {
	node, cost int
	index      int
}

type PriorityQueue []*QueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*QueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (g *Graph) ShortestPath(start int) ([]int, []int) {
	const INF = 1 << 60
	dist := make([]int, len(g.nodes))
	prev := make([]int, len(g.nodes))
	for i := range dist {
		dist[i] = INF
		prev[i] = -1
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &QueueItem{start, 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(*QueueItem)

		if dist[u.node] < u.cost {
			continue
		}

		for _, edge := range g.nodes[u.node] {
			v := edge.to
			newCost := u.cost + edge.cost
			if newCost < dist[v] {
				dist[v] = newCost
				prev[v] = u.node
				heap.Push(pq, &QueueItem{v, newCost})
			}
		}
	}

	path := []int{}
	for u := range prev {
		if prev[u] != -1 {
			path = append([]int{u}, path...)
		}
	}

	return dist, path
}

func main() {
	g := NewGraph()
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 5)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 20)
	g.AddEdge(2, 4, 2)

	dist, path := g.ShortestPath(0)
	fmt.Println("Shortest Path:", path)
	fmt.Println("Distances:", dist)
}

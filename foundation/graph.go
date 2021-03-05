package foundation

type Graph struct {
	v   int
	adj []LinkedList
}

func NewGraph(v int) *Graph {
	graph := &Graph{v: v}
	adj := make([]LinkedList, v)
	for i := 0; i < len(adj); i++ {
		adj[i] = LinkedList{}
	}
	return graph
}

func AddEdge(graph Graph, s int, t int) {
	graph.adj[s].next = NewNode(t)
	graph.adj[t].next = NewNode(s)
}

func BFS(graph Graph, s int, t int) {
	if s == t {
		return
	}

}

var found = false

func DFS(graph Graph, s int, t int) {
	found = false
	visited := make([]bool, graph.v)
	prev := make([]int, graph.v)
	for i := 0; i < graph.v; i++ {
		prev[i] = -1
	}
	recurDfs(graph, s, t, visited, prev)
}

func recurDfs(graph Graph, w int, t int, visited []bool, prev []int) {
	if found {
		return
	}
	visited[w] = true
	if w == t {
		found = true
		return
	}
	for i := 0; i < ListSize(graph.adj[w]); i++ {
		q := graph.adj[w].v
		if !visited[q] {
			prev[q] = w
			recurDfs(graph, q, t, visited, prev)
		}
	}
}

package graph

type Graph struct {
	nodes  []string            //点
	edges  map[string][]string //边
	status map[string]string   //遍历标识
}

func NewGraph() *Graph {
	return &Graph{
		nodes:  []string{},
		edges:  make(map[string][]string),
		status: make(map[string]string),
	}
}

func (g *Graph) AddNode(node string) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) AddEdge(src, target string) {
	g.edges[src] = append(g.edges[src], target)
}

func (g *Graph) DFS(node string, visited map[string]bool) bool {
	visited[node] = true
	g.status[node] = "visiting"

	for _, neighbor := range g.edges[node] {
		if visited[neighbor] {
			if g.status[neighbor] == "visiting" {
				// 循环检测到了
				return true
			}
			continue
		}

		if g.DFS(neighbor, visited) {
			return true
		}
	}

	g.status[node] = "visited"
	return false
}

//有向图结构的循环检测
func HasCycle(graph *Graph) bool {
	visited := make(map[string]bool)

	for _, node := range graph.nodes {
		if !visited[node] {
			if graph.DFS(node, visited) {
				return true
			}
		}
	}
	return false
}

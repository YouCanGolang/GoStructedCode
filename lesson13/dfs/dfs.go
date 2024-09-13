package main

import "fmt"

// Graph 定义图结构，使用邻接表
type Graph struct {
	// 图结构
	vertices map[string][]string
	// 已访问的节点
	visited map[string]bool
}

// NewGraph 创建一个新的图
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string][]string),
		visited:  make(map[string]bool),
	}
}

// AddEdge 添加五向边
func (g *Graph) AddEdge(v1, v2 string) {
	g.vertices[v1] = append(g.vertices[v1], v2)
	g.vertices[v2] = append(g.vertices[v2], v1)
}

// DFS 实现深度优先搜索
func (g *Graph) DFS(v string) {
	// 标记为已访问
	g.visited[v] = true
	fmt.Printf("%s", v)

	// 递归访问所有未访问的节点
	for _, neighbor := range g.vertices[v] {
		// 未访问过
		if !g.visited[neighbor] {
			g.DFS(neighbor)
		}
	}
}

// ClearDFS 清除搜索标记
func (g *Graph) ClearDFS() {
	g.visited = make(map[string]bool)
}

func main() {
	// 创建图
	graph := NewGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "E")
	graph.AddEdge("D", "C")
	graph.AddEdge("E", "A")

	fmt.Println("深度优先搜索结果：")
	// 从 A 开始深度优先搜索
	graph.DFS("A")
	graph.ClearDFS()
	fmt.Println()
	// 从 B 开始深度搜索
	graph.DFS("B")
}

package main

import "fmt"

// Graph 定义图结构
type Graph struct {
	vertices map[string][]string
}

// NewGraph 创建一个新的图
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string][]string),
	}
}

// AddEdge 添加边
func (g *Graph) AddEdge(v1, v2 string) {
	g.vertices[v1] = append(g.vertices[v1], v2)
	g.vertices[v2] = append(g.vertices[v2], v1)
}

// Print 打印图
func (g *Graph) Print() {
	for vertex, edges := range g.vertices {
		fmt.Printf("%s -> %v\n", vertex, edges)
	}
}

func main() {
	graph := NewGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "C")
	graph.Print()
}

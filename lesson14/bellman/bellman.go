package main

import (
	"fmt"
	"math"
)

// Graph 定义图结构，使用邻接表表示
type Graph struct {
	vertices map[string]map[string]int
}

// NewGraph 创建一个新的图
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]map[string]int),
	}
}

// AddVertex 添加顶点
func (g *Graph) AddVertex(v string) {
	if g.vertices[v] == nil {
		g.vertices[v] = make(map[string]int)
	}
}

// AddEdge 添加带权重的有向边
func (g *Graph) AddEdge(v1, v2 string, weight int) {
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.vertices[v1][v2] = weight
}

// BellmanFord 实现 Bellman-Ford 算法
func (g *Graph) BellmanFord(start string) (map[string]int, bool) {
	// 初始化距离表，起始点到自己的距离为0，其他点为无穷大
	distances := make(map[string]int)
	for vertex := range g.vertices {
		distances[vertex] = math.MaxInt32
	}
	distances[start] = 0

	// 获取图中所有的边
	edges := g.getAllEdges()

	// 执行 V-1 轮松弛操作
	for i := 0; i < len(g.vertices)-1; i++ {
		for _, edge := range edges {
			if distances[edge.source] != math.MaxInt32 && distances[edge.source]+edge.weight < distances[edge.destination] {
				distances[edge.destination] = distances[edge.source] + edge.weight
			}
		}
	}

	// 检查负权环
	for _, edge := range edges {
		if distances[edge.source] != math.MaxInt32 && distances[edge.source]+edge.weight < distances[edge.destination] {
			return distances, true // 存在负权环
		}
	}

	return distances, false // 没有负权环
}

// Edge 定义边结构
type Edge struct {
	source, destination string
	weight              int
}

// 获取所有的边
func (g *Graph) getAllEdges() []Edge {
	var edges []Edge
	for v1, neighbors := range g.vertices {
		for v2, weight := range neighbors {
			edges = append(edges, Edge{source: v1, destination: v2, weight: weight})
		}
	}
	return edges
}

func main() {
	// 创建图
	graph := NewGraph()
	graph.AddEdge("A", "B", 1)
	graph.AddEdge("A", "C", 4)
	graph.AddEdge("B", "C", -2)
	graph.AddEdge("B", "D", 2)
	graph.AddEdge("C", "E", 3)
	graph.AddEdge("D", "E", 5)

	// 添加孤立的顶点，确保图中所有顶点都存在
	graph.AddVertex("E")

	// 执行 Bellman-Ford 算法
	distances, hasNegativeCycle := graph.BellmanFord("A")

	if hasNegativeCycle {
		fmt.Println("图中存在负权环")
	} else {
		fmt.Println("从顶点 A 到各个顶点的最短距离：")
		for vertex, distance := range distances {
			if distance == math.MaxInt32 {
				fmt.Printf("%s: ∞\n", vertex)
			} else {
				fmt.Printf("%s: %d\n", vertex, distance)
			}
		}
	}
}

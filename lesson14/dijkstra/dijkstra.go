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

// AddVertex 添加顶点到图中
func (g *Graph) AddVertex(vertex string) {
	if g.vertices[vertex] == nil {
		g.vertices[vertex] = make(map[string]int)
	}
}

// AddEdge 添加带权重的有向边
func (g *Graph) AddEdge(v1, v2 string, weight int) {
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.vertices[v1][v2] = weight
}

// Dijkstra 实现 Dijkstra 算法
func (g *Graph) Dijkstra(start string) map[string]int {
	// 初始化距离表，起始点到自己的距离为0，其他点为无穷大
	distances := make(map[string]int)
	for vertex := range g.vertices {
		distances[vertex] = math.MaxInt32
	}
	distances[start] = 0

	// 已访问的顶点
	visited := make(map[string]bool)

	for len(visited) < len(g.vertices) {
		// 从未访问的顶点中选出距离最小的顶点
		var minVertex string
		minDistance := math.MaxInt32
		for vertex, distance := range distances {
			if !visited[vertex] && distance < minDistance {
				minDistance = distance
				minVertex = vertex
			}
		}

		// 标记为已访问
		visited[minVertex] = true

		// 更新邻接顶点的距离
		for neighbor, weight := range g.vertices[minVertex] {
			if newDist := distances[minVertex] + weight; newDist < distances[neighbor] {
				distances[neighbor] = newDist
			}
		}
	}

	return distances
}

func main() {
	// 创建图
	graph := NewGraph()
	graph.AddEdge("A", "B", 10)
	graph.AddEdge("A", "C", 3)
	graph.AddEdge("C", "B", 1)
	graph.AddEdge("C", "D", 8)
	graph.AddEdge("B", "D", 2)
	graph.AddEdge("D", "E", 6)
	graph.AddEdge("C", "E", 4)

	// 确保所有节点都被初始化
	graph.AddVertex("E")

	// 执行 Dijkstra 算法
	distances := graph.Dijkstra("A")

	fmt.Println("最短路径：", distances)
}

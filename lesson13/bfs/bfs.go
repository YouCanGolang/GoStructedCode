package main

import (
	"container/list"
	"fmt"
)

// Graph 定义图结构，使用邻接表
type Graph struct {
	// 图结构
	vertices map[string][]string
}

// NewGraph 创建一个新的图
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string][]string),
	}
}

// AddEdge 添加五向边
func (g *Graph) AddEdge(v1, v2 string) {
	g.vertices[v1] = append(g.vertices[v1], v2)
	g.vertices[v2] = append(g.vertices[v2], v1)
}

// BFS 实现广度优先搜索
func (g *Graph) BFS(start string) {
	// 创建队列
	queue := list.New()
	// 存储访问过的节点
	visited := make(map[string]bool)

	// 初始化队列，将起始节点入队
	queue.PushBack(start)
	visited[start] = true

	// 循环处理队列中的节点
	for queue.Len() > 0 {
		// 出队
		element := queue.Front()
		v := element.Value.(string)
		queue.Remove(element)
		fmt.Printf("%s ", v)

		// 访问邻接节点
		for _, neighbor := range g.vertices[v] {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
			}
		}
	}
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

	fmt.Println("广度优先搜索结果：")
	// 从 A 开始广度优先搜索
	graph.BFS("A")
	fmt.Println()
	// 从 B 开始广度优先搜索
	graph.BFS("B")
}

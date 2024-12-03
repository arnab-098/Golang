package main

import (
	"fmt"
)

type Graph struct {
	size            int
	adjacencyMatrix [][]int
	vertex          map[string]int
}

func createGraph(vertex []string) Graph {
	size := len(vertex)
	g := Graph{size: size, adjacencyMatrix: make([][]int, size)}
	for i := range g.adjacencyMatrix {
		g.adjacencyMatrix[i] = make([]int, size)
	}
	vertexMap := make(map[string]int)
	for idx, val := range vertex {
		vertexMap[val] = idx
	}
	g.vertex = vertexMap
	return g
}

func displayGraph(g Graph) {
	vertices := make([]string, g.size)
	for idx, val := range g.vertex {
		vertices[val] = idx
	}
	for i := range g.adjacencyMatrix {
		fmt.Print(vertices[i], ":\t")
		for j := range g.adjacencyMatrix[i] {
			if g.adjacencyMatrix[i][j] != 0 {
				fmt.Print(vertices[j], "\t")
			}
		}
		fmt.Println()
	}
}

func addEdge(src string, dest string, g *Graph) error {
	if a, x := g.vertex[src]; x {
		if b, y := g.vertex[dest]; y {
			g.adjacencyMatrix[a][b] = 1
			return nil
		}
	}
	return fmt.Errorf("Vertex not found")
}

func deleteEdge(src string, dest string, g *Graph) error {
	if a, x := g.vertex[src]; x {
		if b, y := g.vertex[dest]; y {
			if g.adjacencyMatrix[a][b] != 0 {
				g.adjacencyMatrix[a][b] = 0
				fmt.Printf("Deleted edge from %s to %s successfully\n", src, dest)
				return nil
			}
			return fmt.Errorf("Edge from %s to %s not exist", src, dest)
		}
	}
	return fmt.Errorf("Vertex not found")
}

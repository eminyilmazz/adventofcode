package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Name  string
	Edges map[*Node]int
}

type Graph struct {
	Nodes map[string]*Node
}

func (g *Graph) AddNode(name string) {
	if g.Nodes == nil {
		g.Nodes = make(map[string]*Node)
	}
	if _, ok := g.Nodes[name]; !ok {
		g.Nodes[name] = &Node{
			Name:  name,
			Edges: make(map[*Node]int),
		}
	}
}

func (g *Graph) GetNode(name string) *Node {
	return g.Nodes[name]
}

func (g *Graph) AddEdge(from, to *Node, weight int) {
	from.Edges[to] = weight
	to.Edges[from] = weight
}

type Path struct {
	Distance int
	Nodes    []*Node
}

func (g *Graph) FindPaths() (Path, error) {
	// Generate all permutations of nodes
	paths := g.GeneratePermutations()

	// Find the shortest path among the permutations
	shortestPath := Path{Distance: math.MaxInt32}
	for _, nodes := range paths {
		distance, err := calculatePathDistance(nodes)
		if err != nil {
			continue
		}

		if distance < shortestPath.Distance {
			shortestPath.Distance = distance
			shortestPath.Nodes = nodes
		}
	}

	if shortestPath.Distance == math.MaxInt32 {
		return Path{}, errors.New("no path found")
	}

	return shortestPath, nil
}

func calculatePathDistance(nodes []*Node) (int, error) {
	distance := 0
	for i := 0; i < len(nodes)-1; i++ {
		if d, ok := nodes[i].Edges[nodes[i+1]]; ok {
			distance += d
		} else {
			return 0, errors.New("no path exists between nodes")
		}
	}
	return distance, nil
}

func (g *Graph) GeneratePermutations() [][]*Node {
	var result [][]*Node
	var keys []string

	for key := range g.Nodes {
		keys = append(keys, key)
	}

	permute(keys, 0, &result, g.Nodes)

	return result
}

func permute(keys []string, start int, result *[][]*Node, nodes map[string]*Node) {
	if start == len(keys) {
		var permutation []*Node
		for _, key := range keys {
			permutation = append(permutation, nodes[key])
		}
		*result = append(*result, permutation)
	} else {
		for i := start; i < len(keys); i++ {
			keys[start], keys[i] = keys[i], keys[start]
			permute(keys, start+1, result, nodes)
			keys[start], keys[i] = keys[i], keys[start]
		}
	}
}

func createGraph() Graph {
	inputs := readInput()
	g := Graph{}
	for _, input := range inputs {
		s := strings.Split(input, " ")
		g.AddNode(s[0])
		g.AddNode(s[2])
		weight, _ := strconv.Atoi(s[4])
		g.AddEdge(g.GetNode(s[0]), g.GetNode(s[2]), weight)
	}
	return g
}

func getPathNames(nodes []*Node) []string {
	var names []string
	for _, node := range nodes {
		names = append(names, node.Name)
	}
	return names
}

func partOne() {
	g := createGraph()
	g.PrintGraph()
	fmt.Println()
	shortest, err := g.FindPaths()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Shortest Path:", getPathNames(shortest.Nodes))
		fmt.Println("Distance:", shortest.Distance)
	}
}

func main() {
	partOne()
}

func readInput() []string {
	input, err := os.Open("input.txt")
	var lines []string
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func (g *Graph) PrintGraph() {
	for _, node := range g.Nodes {
		fmt.Printf("Node: %s\n", node.Name)

		for to, weight := range node.Edges {
			fmt.Printf("  Edge: %s -> %s (Weight: %d)\n", node.Name, to.Name, weight)
		}

		fmt.Println()
	}
}

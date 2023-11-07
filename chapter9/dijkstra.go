package chapter9

import (
	"errors"
	"math"
	"slices"
)

type Node struct {
	Name      string
	neighbors map[string]int
}

type Graph struct {
	Nodes []Node
}

type nodeInfo struct {
	predecessor string
	weight      int
	visited     bool
}

func (g *Graph) ShortestPath(start string, end string) (ShortestPath, error) {
	_, err := g.node(start)
	if err != nil {
		return ShortestPath{}, err
	}
	_, err = g.node(end)
	if err != nil {
		return ShortestPath{}, err
	}
	nodes := map[string]*nodeInfo{
		start: {
			predecessor: "",
			weight:      0,
			visited:     false,
		},
	}

	err = g.dijkstra(start, end, nodes)
	if err != nil {
		return ShortestPath{}, err
	}
	return g.buildResult(start, end, nodes), nil
}

func (g *Graph) dijkstra(start string, end string, nodes map[string]*nodeInfo) error {
	cheapestUnvisitedNode := start
	for cheapestUnvisitedNode != end {
		info := nodes[cheapestUnvisitedNode]
		info.visited = true
		node, err := g.node(cheapestUnvisitedNode)
		if err != nil {
			return err
		}

		for neighbor, weight := range node.neighbors {
			neighborInfo := nodes[neighbor]
			if neighborInfo == nil {
				nodes[neighbor] = &nodeInfo{
					predecessor: node.Name,
					weight:      info.weight + weight,
					visited:     false,
				}
			} else {
				weightFromCurrentNode := info.weight + weight
				if neighborInfo.weight > weightFromCurrentNode {
					neighborInfo.predecessor = node.Name
					neighborInfo.weight = weightFromCurrentNode
				}
			}
		}
		cheapestUnvisitedNode = nextNode(nodes)
	}
	return nil
}

func (g *Graph) buildResult(start string, end string, nodes map[string]*nodeInfo) ShortestPath {
	path := make([]string, 0, len(nodes))

	currentNode := nodes[end]
	weight := currentNode.weight
	path = append(path, end)
	for currentNode != nodes[start] {
		path = append(path, currentNode.predecessor)
		currentNode = nodes[currentNode.predecessor]
	}
	slices.Reverse(path)
	return ShortestPath{
		Path:   path,
		Weight: weight,
	}
}

func nextNode(nodes map[string]*nodeInfo) string {
	var cheapestNeighbor string
	weight := math.MaxInt
	for name, info := range nodes {
		if info.visited {
			continue
		}
		if info.weight < weight {
			weight = info.weight
			cheapestNeighbor = name
		}
	}
	return cheapestNeighbor
}

func (g *Graph) node(name string) (Node, error) {
	for _, node := range g.Nodes {
		if name == node.Name {
			return node, nil
		}
	}
	return Node{}, errors.New("no node named " + name + " found")
}

type ShortestPath struct {
	Path   []string
	Weight int
}

func NewGraph(nodes []Node) Graph {
	return Graph{
		nodes,
	}
}

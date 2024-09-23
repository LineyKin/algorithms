package alg

import "fmt"

type Node string

type WeightedEdge map[Node]int

type WeightedGraph map[Node]WeightedEdge

type Costs map[Node]int

type Processed []Node

const inf = 1000000000

func Dijkstra(graph WeightedGraph, start, end Node) Costs {

	// таблица стоимостей
	costs := make(Costs)

	// узлы, которые мы уже проверили
	var processed Processed

	var neighbors WeightedEdge

	for node, _ := range graph {
		if node != start {
			startEdge := graph[start]
			val, ok := startEdge[node]
			if ok {
				costs[node] = val
			} else {
				costs[node] = inf
			}
		}
	}

	node := findNodeLowestCost(costs, processed)

	fmt.Println(node)

	for node != "" {
		cost := costs[node]
		neighbors = graph[node]
		for nNode, nCost := range neighbors {
			newCost := cost + nCost
			if newCost < costs[nNode] {
				costs[nNode] = newCost
			}
		}

		// вершину данной итерации добавляем в список уже обработанных
		processed = append(processed, node)
		node = findNodeLowestCost(costs, processed)
	}

	return costs
}

func findNodeLowestCost(costs Costs, processed Processed) Node {
	lowestCost := inf
	var lowestNode Node
	for node, cost := range costs {
		if cost < lowestCost && !processed.hasNode(node) {
			lowestCost = cost
			lowestNode = node
		}
	}

	return lowestNode
}

// проверка, является ли узел из числа уже обработанных
func (p Processed) hasNode(node Node) bool {
	for _, processedNode := range p {
		if processedNode == node {
			return true
		}
	}

	return false
}

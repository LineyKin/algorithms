package alg

// вершина (узел) графа
type Node string

// ребро взвешенного графа
type WeightedEdge map[Node]int

// взвешенный граф
type WeightedGraph map[Node]WeightedEdge

// таблица стоимостей
type Costs map[Node]int

// узлы, которые мы уже проверили
type Processed []Node

const Inf = 1000000000

// возвращает кратчайшие пути из вершины start в остальные вершины графа graph
func Dijkstra(graph WeightedGraph, start Node) Costs {
	costs := make(Costs)
	var processed Processed

	// рёбра, выходящие из рассматриваемой вершины в соседние вершины
	var neighbors WeightedEdge

	for node := range graph {
		if node != start {
			startEdge := graph[start]
			val, ok := startEdge[node]
			if ok {
				costs[node] = val
			} else {
				costs[node] = Inf
			}
		}
	}

	node := findNodeLowestCost(costs, processed)

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
	lowestCost := Inf
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

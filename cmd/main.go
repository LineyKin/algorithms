package main

import (
	"algorithms/alg"
	"fmt"
)

func main() {

	graph := make(alg.WeightedGraph)
	graph["a"] = alg.WeightedEdge{"b": 2, "c": 1}
	graph["b"] = alg.WeightedEdge{"f": 7}
	graph["c"] = alg.WeightedEdge{"d": 5, "e": 2}
	graph["d"] = alg.WeightedEdge{"f": 2}
	graph["e"] = alg.WeightedEdge{"f": 1}
	graph["f"] = alg.WeightedEdge{"g": 1}
	graph["g"] = alg.WeightedEdge{}

	fmt.Println(alg.Dijkstra(graph, "a", "g"))

}

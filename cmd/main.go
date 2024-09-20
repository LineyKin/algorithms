package main

import (
	"algorithms/alg"
	"fmt"
)

func main() {

	graph := make(alg.Graph)
	graph["a"] = alg.Edge{"b", "c"}
	graph["b"] = alg.Edge{"f"}
	graph["c"] = alg.Edge{"d", "e"}
	graph["d"] = alg.Edge{"f"}
	graph["e"] = alg.Edge{"f"}
	graph["f"] = alg.Edge{"g"}

	fmt.Println(alg.BreadthFirstSearch(graph, "a", "g"))

}

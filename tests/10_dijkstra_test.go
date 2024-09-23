package test

import (
	"algorithms/alg"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Dijkstra struct {
	graph alg.WeightedGraph
	start alg.Node
	want  alg.Costs
}

func TestDijkstra(t *testing.T) {
	graph := make(alg.WeightedGraph)
	graph["a"] = alg.WeightedEdge{"b": 2, "c": 1}
	graph["b"] = alg.WeightedEdge{"f": 7}
	graph["c"] = alg.WeightedEdge{"d": 5, "e": 2}
	graph["d"] = alg.WeightedEdge{"f": 2}
	graph["e"] = alg.WeightedEdge{"f": 1}
	graph["f"] = alg.WeightedEdge{"g": 1}
	graph["g"] = alg.WeightedEdge{}

	testData := []Dijkstra{
		{graph, "a", alg.Costs{"b": 2, "c": 1, "d": 6, "e": 3, "f": 4, "g": 5}},
		{graph, "g", alg.Costs{"b": alg.Inf, "c": alg.Inf, "d": alg.Inf, "e": alg.Inf, "f": alg.Inf, "a": alg.Inf}},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, alg.Dijkstra(v.graph, v.start))
	}
}

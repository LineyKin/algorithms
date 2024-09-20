package test

import (
	"algorithms/alg"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BreadthFirstSearch struct {
	graph alg.Graph
	start string
	end   string
	want  bool
}

func TestBreadthFirstSearch(t *testing.T) {
	graph1 := make(alg.Graph)
	graph1["a"] = alg.Edge{"b", "c"}
	graph1["b"] = alg.Edge{"f"}
	graph1["c"] = alg.Edge{"d", "e"}
	graph1["d"] = alg.Edge{"f"}
	graph1["e"] = alg.Edge{"f"}
	graph1["f"] = alg.Edge{"g"}

	testData := []BreadthFirstSearch{
		{graph1, "a", "g", true},
		{graph1, "a", "x", false},
		{graph1, "a", "b", true},
		{graph1, "x", "g", false},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, alg.BreadthFirstSearch(v.graph, v.start, v.end))
	}
}

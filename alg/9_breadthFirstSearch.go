package alg

// рёбра графа
type Edge []string

// сам граф
type Graph map[string]Edge

// поиск в ширину
// используется очередь, принцип FIFO
// показывает, существует ли путь из одной вершины в другую
func BreadthFirstSearch(g Graph, start, end string) bool {
	queue := make([]string, 0, len(g))
	queue = append(queue, start)
	for len(queue) > 0 {

		// прописываем поведение очереди
		current := queue[0]
		queue = queue[1:]
		node, ok := g[current]
		if !ok {
			g[current] = Edge{}
		}

		if node.hasElement(end) {
			return true
		} else {
			queue = append(queue, node...)
		}
	}

	return false
}

func (e Edge) hasElement(elem string) bool {
	for _, v := range e {
		if v == elem {
			return true
		}
	}

	return false
}

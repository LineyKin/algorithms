package alg

type TreeNode struct {
	Val      int
	Children Tree
}

// дерево n-размерности
type Tree []TreeNode

// возвращает сумму всех вершин дерева
func (t Tree) TreeSum() int {
	if len(t) == 0 {
		return 0
	}

	var sum int
	stack := []TreeNode{}
	for _, node := range t {
		stack = append(stack, node)
	}
	for len(stack) > 0 {
		lastIndex := len(stack) - 1
		node := stack[lastIndex]
		stack = stack[0:lastIndex]
		sum += node.Val
		if node.Children != nil {
			for _, n := range node.Children {
				stack = append(stack, n)
			}
		}
	}

	return sum
}

// В отличии от примера на js по которому я изучал алгоритм
// здесь, в рекурсивном обходе, не надо делать проверку на пустоту детей, для выхода из рекурсии.
//
// Отсуствие детей я задаю как Children: nil
// В свою очередь Children ожидает слайс, а слайс с len=0 И cap=0 и есть nil
// RecursiveTreeSum() запускается для Children равным nil, но т.к это пустой слайс, то
// цикл не запускается, а просто возвращается сумма. А значит и не делается рекурсивный вызов, т.к. он объявлен в цикле
func (t Tree) RecursiveTreeSum() int {
	var sum int
	for _, node := range t {
		sum += node.Val
		sum += node.Children.RecursiveTreeSum()
	}

	return sum
}

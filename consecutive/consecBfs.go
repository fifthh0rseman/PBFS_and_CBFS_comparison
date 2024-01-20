package consecutive

import (
	"BFS/utils"
	"fmt"
)

func BFSCube(dim, start, finish int) {
	visited := make([]bool, dim*dim*dim)
	q := utils.NewQueue()
	q.Push(start)
	for !q.Empty() {
		node := q.Pop()
		visited[node] = true
		children := utils.GetNeighbours(node, dim)
		for _, c := range children {
			if visited[c] == false {
				q.Push(c)
				visited[c] = true
			}
		}
	}
	fmt.Printf("Node %d: unreachable", finish)
}

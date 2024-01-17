package parallel

import (
	"BFS/utils"
	"fmt"
	"sync"
	"sync/atomic"
)

func ParallelBfs(dim, start, finish int) {
	numNodes := dim * dim * dim
	visited := make([]atomic.Bool, numNodes)
	visited[start].CompareAndSwap(false, true)
	frontier := make([]int, numNodes)
	frontier[0] = start
	pNext := atomic.Int32{}
	pNext.Store(1)
	pPrev := int32(0)
	for !pNext.CompareAndSwap(pPrev, pPrev) {
		var wg sync.WaitGroup
		pNextCopy := pNext.Load()
		for i := pPrev; i < pNextCopy; i++ {
			wg.Add(1)
			i := i
			go func(wg *sync.WaitGroup, i int32) {
				cur := frontier[i]
				fmt.Println("Cur node:", cur)
				for next := range utils.GetNeighbours(cur, dim) {
					if visited[next].CompareAndSwap(false, true) {
						if cur == finish {
							fmt.Println("Finish node reached")
							return
						}

						frontier[pNext.Add(1)] = next
						fmt.Println("func finish", i, "node: ", next)
					}
				}
				wg.Done()
			}(&wg, i)
		}
		wg.Wait()
		fmt.Println("pprev: ", pPrev)
		fmt.Println("pnext: ", pNext.Load())
		pPrev = pNextCopy
	}
}

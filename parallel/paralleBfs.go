package parallel

import (
	"BFS/utils"
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
	var block int32 = 100
	for !pNext.CompareAndSwap(pPrev, pPrev) {
		pNextCopy := pNext.Load()
		var wg sync.WaitGroup
		for i := pPrev; i < pNextCopy; i += block {
			wg.Add(1)
			go func(wg *sync.WaitGroup, i int32) {
				for j := i; j < pNextCopy && j < i+block; j++ {
					cur := frontier[j]
					for _, next := range utils.GetNeighbours(cur, dim) {
						if visited[next].CompareAndSwap(false, true) {
							frontier[pNext.Add(1)-1] = next
						}
					}
				}
				wg.Done()
			}(&wg, i)
		}
		wg.Wait()
		pPrev = pNextCopy
	}
}

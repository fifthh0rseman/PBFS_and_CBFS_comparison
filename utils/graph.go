package utils

func GetNeighbours(node, dim int) []int {
	//node number is n^2*k + ni + j
	k := node / dim / dim
	i := (node - dim*dim*k) / dim
	j := node - dim*dim*k - dim*i
	ret := make([]int, 6)
	ctr := 0
	if k < dim-1 {
		ret[ctr] = dim*dim*(k+1) + dim*i + j
		ctr++
	}
	if k > 0 {
		ret[ctr] = dim*dim*(k-1) + dim*i + j
		ctr++
	}

	if i < dim-1 {
		ret[ctr] = dim*dim*k + dim*(i+1) + j
		ctr++
	}
	if i > 0 {
		ret[ctr] = dim*dim*k + dim*(i-1) + j
		ctr++
	}

	if j < dim-1 {
		ret[ctr] = dim*dim*k + dim*i + j + 1
		ctr++
	}
	if j > 0 {
		ret[ctr] = dim*dim*k + dim*i + j - 1
		ctr++
	}
	ret = ret[0:ctr]
	return ret
}

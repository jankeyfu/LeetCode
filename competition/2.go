package competition

func numWays(n int, relation [][]int, k int) int {
	nodes := map[int][]int{}
	for _, rel := range relation {
		if _, ok := nodes[rel[0]]; !ok {
			nodes[rel[0]] = []int{rel[1]}
		} else {
			nodes[rel[0]] = append(nodes[rel[0]], rel[1])
		}
	}

	// fmt.Println(isok(n, nodes, 0, 1, 0))
	return isok(n, k, nodes, 0, 1, 0)
}

func isok(ns, ks int, node map[int][]int, nodeN, num, total int) int {
	// fmt.Println(ns, nodeN, num, total)
	if num > ks {
		return total
	}
	if _, ok := node[nodeN]; !ok {
		return total
	}
	for _, n := range node[nodeN] {
		if n == ns-1 && num == ks {
			total++
		}
		total = isok(ns, ks, node, n, num+1, total)
	}
	return total
}

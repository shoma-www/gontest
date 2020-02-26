package shortestpath

// WarshallFloyd ワーシャルフロイド法
// 動的計画法を用いて全点対の最短距離(最小コスト)を求める。
// TimeLimit 1sec程度で|V|の500~600程度が限界
// 計算量：O(|V|^3)
func WarshallFloyd(edges Edges, vertex, start, end int) []int {
	d := make([][]int, vertex)
	nodeIndex := make([][]int, vertex)
	for i := 0; i < vertex; i++ {
		d[i] = make([]int, vertex)
		nodeIndex[i] = make([]int, vertex)
		for j := 0; j < vertex; j++ {
			nodeIndex[i][j] = i
			if i == j {
				d[i][j] = 0
			} else {
				d[i][j] = 1e+6
			}
		}
	}

	for _, e := range edges {
		d[e.from][e.to] = int32min(d[e.from][e.to], e.cost)
		d[e.to][e.from] = int32min(d[e.to][e.from], e.cost)
	}

	for i := 0; i < vertex; i++ {
		for j := 0; j < vertex; j++ {
			for k := 0; k < vertex; k++ {
				if d[j][k] > d[j][i] + d[i][k] {
					d[j][k] = d[j][i] + d[i][k]
					nodeIndex[j][k] = nodeIndex[i][k]
				}
			}
		}
	}

	var route []int
	route = append(route, end)
	current := end
	for true {
		next := nodeIndex[start][current]

		route = append(route[0:1], route[0:]...)
		route[0] = next
		if next != start {
			current = next
		} else {
			break
		}
	}
	return route
}
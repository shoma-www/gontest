package shortestpath

import (
	"fmt"
	"math"
)

// Edge 辺の定義
type Edge struct {
	from int
	to   int
	cost int
}

// Edges 辺の配列
type Edges []Edge

// BellmanFord コスト付きのグラフの最短経路を探索するアルゴリズム
// ダイクストラ法のほうが高速なため、コストに負が含まれている場合に使用
// また、負の閉路が存在することがわかる
// コストの更新が起こらなくなるまで更新する
// 計算量は、O(VE)。
// 参考 http://www.thothchildren.com/chapter/5b5740b7103f2f3168716d80
func BellmanFord(edges Edges, vertex, start, end int) ([]int, error) {
	// 求めたい経路のコストを初期化
	d := make([]int, vertex)
	for i := 0; i < vertex; i++ {
		d[i] = math.MaxInt32
	}

	// 始点はゼロで初期化
	d[start] = 0
	// 各頂点でのコストの計算
	for i := 0; i < vertex; i++ {
		for _, e := range edges {
			if d[e.to] > d[e.from] + e.cost {
				d[e.to] = d[e.from] + e.cost
				if i == vertex - 1 {
					return nil, fmt.Errorf("negative loop exists")
				}
			}
		}
	}

	// 最短経路の出力
	p := make([]int, 1)
	p[0] = end
	pos := end
	cost := d[end]
	for cost > 0 {
		for _, e := range edges {
			if e.to == pos {
				if d[e.from] == cost - e.cost {
					p = append(p[0:1], p[0:]...)
					p[0] = e.from
					pos = e.from
					cost = cost - e.cost
					break
				}
			}
		}
	}

	return p, nil
}
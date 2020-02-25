package shortestpath

import (
	"math"
	"strconv"

	"github.com/shoma-www/my-contest-go/algorithm/structure"
)

// Dijkstra ダイクストラ法
// 負のコストを含まないグラフで利用
// コストが低いノードのみを処理することでベルマンフォード法より高速
// 計算量
// 通常の実装：O(|V|^2)  ※非常に密なグラフに対して優位
// 優先度付きキュー：O(|E|log|V|)
// 参考 http://www.thothchildren.com/chapter/5b6ae3282787593b86358a54
func Dijkstra(edges Edges, vertex, start, end int) []int {
	// 求めたい経路のコストを初期化
	d := make([]int, vertex)
	for i := 0; i < vertex; i++ {
		d[i] = math.MaxInt32
	}
	d[start] = 0

	graph := make([]Edges, vertex)
	for _, e := range edges {
		graph[e.from] = append(graph[e.from], e)
		toE := Edge{
			from: e.to,
			to: e.from,
			cost: e.cost,
		}
		graph[toE.from] = append(graph[toE.from], toE)
	}

	used := make([]bool, vertex)
	for i := 0; i < vertex; i++ {
		id := i
		for j := 0; j < vertex; j++ {
			if !used[j] && d[id] > d[j] {
				id = j
			}
		}
		used[id] = true

		for _, e := range graph[id] {
			d[e.to] = int32min(d[e.to], d[id] + e.cost)
		}
	}

	// 最短経路の出力
	p := make([]int, 1)
	p[0] = end
	pos := end
	cost := d[end]
	for cost > 0 {
		for _, e := range graph[pos] {
			if d[e.to] == cost - e.cost {
				p = append(p[0:1], p[0:]...)
				p[0] = e.to
				pos = e.to
				cost = d[e.to]
				break
			}
		}
	}

	return p
}

// DijkstraP ダイクストラ法(優先度付きキュー)
// 負のコストを含まないグラフで利用
// コストが低いノードのみを処理することでベルマンフォード法より高速
// 計算量
// 通常の実装：O(|V|^2)  ※非常に密なグラフに対して優位
// 優先度付きキュー：O(|E|log|V|)
// 参考 http://www.thothchildren.com/chapter/5b6ae3282787593b86358a54
func DijkstraP(edges Edges, vertex, start, end int) []int {
	// 求めたい経路のコストを初期化
	d := make([]int, vertex)
	for i := 0; i < vertex; i++ {
		d[i] = math.MaxInt32
	}
	d[start] = 0

	structure.NewPriorityQueue()

	graph := make([]Edges, vertex)
	for _, e := range edges {
		graph[e.from] = append(graph[e.from], e)
		toE := Edge{
			from: e.to,
			to: e.from,
			cost: e.cost,
		}
		graph[toE.from] = append(graph[toE.from], toE)
	}

	q := structure.NewPriorityQueue()
	q.Push(&structure.Item{
		Priority: 0,
		Value: strconv.Itoa(start),
		Index: 0,
	})
	for !q.Empty() {
		item := q.Pop()
		value, _ := strconv.Atoi(item.Value)
		if d[value] < item.Priority { continue }
		for _, e := range graph[value] {
			if d[e.to] > d[value] + e.cost {
				d[e.to] = d[value] + e.cost
				q.Push(&structure.Item{
					Priority: d[e.to],
					Value: strconv.Itoa(e.to),
				})
			}
		}
	}

	// 最短経路の出力
	p := make([]int, 1)
	p[0] = end
	pos := end
	cost := d[end]
	for cost > 0 {
		for _, e := range graph[pos] {
			if d[e.to] == cost - e.cost {
				p = append(p[0:1], p[0:]...)
				p[0] = e.to
				pos = e.to
				cost = d[e.to]
				break
			}
		}
	}

	return p
}

func int32min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

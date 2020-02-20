package algorithm

// UnionFindTree 重み付き素集合データ構造
// union by rank
// parentに各頂点の親の頂点を格納
// rankに各頂点の木の高さを格納
// weightに各頂点の重みを格納
//
// 制約：Union-Find木はまとめられても分割はできない
// 計算量：経路圧縮とランクで「O(α(n))」　※α(n)はアッカーマン関数の逆関数(< lon(n))
type UnionFindTree struct {
	parent []int
	rank   []int
	weight []int
}

// InitUnionFindTree 初期化
func InitUnionFindTree(n int) *UnionFindTree {
	p := make([]int, n)
	r := make([]int, n)
	w := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}

	return &UnionFindTree{
		parent: p,
		rank:   r,
		weight: w,
	}
}

// Root Union-Find木からルートの値を取得
func (u *UnionFindTree) Root(x int) int {
	if u.parent[x] == x {
		return x
	}
	// 経路圧縮
	r := u.Root(u.parent[x])
	// 累積和
	u.weight[x] += u.weight[u.parent[x]]
	u.parent[x] = r
	return u.parent[x]
}

// Weight 重みを返す
func (u *UnionFindTree) Weight(x int) int {
	// 経路圧縮により重みを累積和で計算
	u.Root(x)
	return u.weight[x]
}

// Same ２値のルートが同じかチェック
func (u *UnionFindTree) Same(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

// Unite xとyの属する集合を結合
func (u *UnionFindTree) Unite(x, y, w int) {
	w += u.Weight(x)
	w -= u.Weight(y)
	xR, yR := u.Root(x), u.Root(y)
	
	if xR == yR {
		return
	}

	if u.rank[xR] < u.rank[yR] {
		xR, yR = yR, xR
		w = -1 * w
	}
		
	if u.rank[xR] == u.rank[yR] {
		u.rank[xR]++
	}

	u.parent[yR] = xR
	u.weight[yR] = w
}

// Diff 重みの差分を取得
func (u *UnionFindTree) Diff(x, y int) int {
	return u.Weight(y) - u.Weight(x)
}
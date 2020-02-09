package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(in io.Reader, out io.Writer) {
	var bs = NewBufScanner(in)
	var bw = NewBufWriter(out)
	defer bw.w.Flush()

	n, m := bs.IntScan(), bs.IntScan()
	
	uf := InitUnionFindTree(n)

	for i := 0; i < m; i++ {
		l, r, d := bs.IntScan(), bs.IntScan(), bs.IntScan()
		l--
		r--
		if uf.Same(l, r) {
			diff := uf.Diff(l, r)
			if diff != d {
				bw.Printf("No\n")
				return
			}
		} else {
			uf.Unite(l, r, d)
		}
	}
	bw.Printf("Yes\n")
}

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

// BufScanner original scanner
type BufScanner struct {
	s *bufio.Scanner
}

// NewBufScanner constructer
func NewBufScanner(in io.Reader) *BufScanner {
	s := bufio.NewScanner(in)
	s.Buffer(make([]byte, 1024), 1e+9)
	s.Split(bufio.ScanWords)
	return &BufScanner{
		s: s,
	}
}

// Scan Scan Data
func (b *BufScanner) Scan() string {
	b.s.Scan()
	return b.s.Text()
}

// IntScan Scan Data
func (b *BufScanner) IntScan() int {
	v, err := strconv.Atoi(b.Scan())
	if err != nil {
		panic(err)
	}
	return v
}

// BufWriter original writer
type BufWriter struct {
	w *bufio.Writer
}

// NewBufWriter constructer
func NewBufWriter(out io.Writer) *BufWriter {
	w := bufio.NewWriter(out)
	return &BufWriter{
		w: w,
	}
}

// Printf Output file
func (b *BufWriter) Printf(format string, a ...interface{}) {
	fmt.Fprintf(b.w, format, a...)
}
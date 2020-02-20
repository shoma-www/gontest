package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

// K 市松模様の正方形の大きさ
var K int

func solve(in io.Reader, out io.Writer) {
	var bs = NewBufScanner(in)
	var bw = NewBufWriter(out)
	defer bw.w.Flush()

	line := strings.Split(bs.Scan(), " ")
	N, _ := strconv.Atoi(line[0])
	K, _ = strconv.Atoi(line[1])

	m := make([][]int, 2*K)
	sum := make([][]int, 2*K)

	for i := 0; i < 2*K; i++ {
		m[i] = make([]int, 2*K)
		sum[i] = make([]int, 2*K)
	}

	for i := 0; i < N; i++ {
		line = strings.Split(bs.Scan(), " ")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		if line[2] == "W" {
			y += K
		}

		m[x%(2*K)][y%(2*K)]++
	}

	// 2次元累積和（ゼータ返還）
	// x方向に加算
	for y := 0; y < 2*K; y++ {
		for x := 0; x < 2*K; x++ {
			if x > 0 {
				m[x][y] += m[x-1][y]
			}
		}
	}
	// y方向に加算
	for x := 0; x < 2*K; x++ {
		for y := 0; y < 2*K; y++ {
			if y > 0 {
				m[x][y] += m[x][y-1]
			}
		}
	}

	// 範囲内で最大個数を求める
	max := 0
	for x := 0; x < 2*K; x++ {
		for y:= 0; y < 2*K; y++ {
			// 全範囲をブロックごとに計算している
			tmp := get(x, y, m)
			tmp += get(x-K, y+K, m)
			tmp += get(x+K, y-K, m)
			tmp += get(x+K, y+K, m)
			tmp += get(x-K, y-K, m)
			// 2Kの周期でサイクルすることから-2*Kをしている
			// K=6(x=0~11)の場合に x=11だとx=-1を起点に範囲ができるため、
			// x=0~5の範囲の和が求められる
			// 2Kを引いて負になり計算対象にならない場合は、
			// 希望の色になっていない範囲を示す
			tmp += get(x-2*K, y, m)
			tmp += get(x, y-2*K, m)
			tmp += get(x-2*K, y-2*K, m)
			max = int(math.Max(float64(max), float64(tmp)))
		}
	}

	bw.Printf("%d\n", max)
}

// x,yを起点とする範囲内(x~x+K-1, y~y+K-1)の個数を求める
func get(x, y int, m [][]int) int {
	if !isGet(x, y) {
		return 0
	}

	tx := int(math.Min(float64(x+K-1), float64(len(m)-1)))
	ty := int(math.Min(float64(y+K-1), float64(len(m)-1)))
	num := m[tx][ty]
	if in(x-1, y-1) {
		num += m[x-1][y-1]
	}
	if in(x-1, ty) {
		num -= m[x-1][ty]
	}
	if in(tx, y-1) {
		num -= m[tx][y-1]
	}
	return num
}

// 0~2*Kの間に存在するか
func in(x, y int) bool {
	return 0 <= x && x < 2*K && 0 <= y && y < 2*K
}

// x.yを起点とする正方形の範囲の４点（x~xK-1,y~y+K-1）のうち１つでも0~2*Kの間に存在するか
func isGet(x, y int) bool {
	return in(x, y) || in(x+K-1, y+K-1) || in(x-1, y+K-1) || in(x+K-1, y-1)
}

// BufScanner original scanner
type BufScanner struct {
	s *bufio.Scanner
}

// NewBufScanner constructer
func NewBufScanner(in io.Reader) *BufScanner {
	s := bufio.NewScanner(in)
	return &BufScanner{
		s: s,
	}
}

// Scan Scan Data
func (b *BufScanner) Scan() string {
	b.s.Scan()
	return b.s.Text()
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
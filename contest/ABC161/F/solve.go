package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(in io.Reader, out io.Writer) {
	bs := NewBufScanner(in)
	bw := NewBufWriter(out)
	defer bw.w.Flush()
	n := bs.IntScan()

	if n == 2 {
		bw.Printf("1")
		return
	}

	count := 2
	var k int64
	for k = 2; k <= int64(math.Sqrt(float64(n))); k++ {
		// kがnの約数かどうかチェック
		if isOne(n, k) {
			count++
		}

		// √nより大きい約数をチェックする
		t := n / k
		// √n以下の場合、次のkに
		if t <= k {
			continue
		}
		if isOne(n, t) {
			count++
		}
	}

	bw.Printf("%v", count)
}

func isOne(n, k int64) bool {
	x := n
	for x >= k && x % k == 0 {
		x /= k
	}
	// 割ったあまりが１、もしくはx=１の場合、対象
	if x % k == 1 {
		return true
	}
	return false
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
func (b *BufScanner) IntScan() int64 {
	v, err := strconv.ParseInt(b.Scan(), 10, 64)
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
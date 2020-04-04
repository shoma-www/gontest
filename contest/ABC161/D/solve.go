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
	bs := NewBufScanner(in)
	bw := NewBufWriter(out)
	defer bw.w.Flush()
	k := bs.IntScan()

	q := make([]int64, 0)

	/*
	差が１の値は、一つ上の桁から「-1」、「+0」、「+1」した値がその桁になる
	つまり、下２桁が同じなるように計算（10x+x%10）し、その前後±1がルンルンの値の範囲になる
	これをqueueを使って再帰的に計算すると導出できる。

	queueに1~9をenqueする。
	その後、ポップしたxに下記を行う
	１．10x+x%10-1!=9 のとき 10x+x%10-1をenque
	２．10x+x%10をenque
	３．10x+x%10+1!=0 のとき 10x+x%10+1をenque
	上記をk回になるまで繰り返し行い、k回目にdequeした数値が結果
	*/
	for i:=1; i<10; i++ {
		q = append(q, int64(i))
	}

	count := 0
	for {
		x := q[0]
		q = q[1:]
		count++
		if count == k {
			bw.Printf("%v", x)
			return
		}

		midVal := 10 * x + x % 10
		if (midVal - 1) % 10 != 9 {
			q = append(q, midVal - 1)
		}
		q = append(q, midVal)
		if (midVal + 1) % 10 != 0 {
			q = append(q, midVal + 1)
		}
	}
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

// Int64Scan Scan Data
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
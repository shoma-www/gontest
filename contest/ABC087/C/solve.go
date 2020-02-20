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

	n := bs.IntScan()
	a1 := []int{}
	for len(a1) < n {
		a1 = append(a1, bs.IntScan())
	}
	a2 := []int{}
	for len(a2) < n {
		a2 = append(a2, bs.IntScan())
	}

	a1memo := memo(a1)
	a2memo := memo(a2)

	max := 0
	for i := 0; i < n; i++ {
		sum := a1memo[i]
		sum += a2memo[n-1]
		if i - 1 >= 0 {
			sum -= a2memo[i-1]
		}
		if max < sum {
			max = sum
		}
	}
	bw.Printf("%d\n", max)
}

func memo(a []int) []int {
	r := make([]int, len(a))
	r[0] = a[0]
	for i := 1; i < len(a); i++ {
		r[i] = a[i] + r[i-1]
	}
	return r
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
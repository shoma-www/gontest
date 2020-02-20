package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"math"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(in io.Reader, out io.Writer) {
	var bs = NewBufScanner(in)
	var bw = NewBufWriter(out)
	defer bw.w.Flush()

	n := bs.IntScan()
	as := []int{}

	for i := 0; i < n; i++ {
		as = append(as,  bs.IntScan())
	}

	answers := []answer{}
	max := maxIndex(as)
	min := minIndex(as)
	if abs(as[max]) >= abs(as[min]) {
		val := as[max]
		for i := 0; i < n; i++ {
			as[i] += val
			answers = append(answers, answer{x: max, y: i})
		}

		for i := 0; i < n - 1; i++ {
			as[i+1] += as[i]
			answers = append(answers, answer{x: i, y: i+1})
		}
	} else {
		val := as[min]
		for i := 0; i < n; i++ {
			as[i] += val
			answers = append(answers, answer{x: min, y: i})
		}

		for i := n-1; i > 0; i-- {
			as[i-1] += as[i]
			answers = append(answers, answer{x: i, y: i-1})
		}
	}

	bw.Printf("%d\n", len(answers))
	for _, a := range answers {
		bw.Printf("%d %d\n", a.x+1, a.y+1)
	}
}


type answer struct {
	x int
	y int
}

func maxIndex(l []int) int {
	max := math.MinInt32
	idx := -1
	for i, v := range l {
		if max <= v {
			max = v
			idx = i
		}
	}
	return idx
}

func minIndex(l []int) int {
	min := math.MaxInt32
	idx := -1
	for i, v := range l {
		if min >= v {
			min = v
			idx = i
		}
	}
	return idx
}


func abs(x int) int {
	return int(math.Abs(float64(x)))
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
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
	x, y :=bs.IntScan(), bs.IntScan()
	n := x
	maxCount := 0
	for !isPrime(n) && n < y - 1 {
		count := countArray(n, y)
		if maxCount < count {
			maxCount = count
		}
		n++
	}
	count := countArray(n, y)
	if maxCount < count {
		maxCount = count
	}

	bw.Printf("%v", maxCount)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n % 2 == 0  {
		return false
	}

	sNum := int(math.Sqrt(float64(n)))
	for i := 3; i <= sNum; i += 2 {
		if n % int(i) == 0 {
			return false
		}
	}

	return true
}

func countArray(x, y int) int {
	count := 0
	for x <= y {
		count++
		x *= 2
	}
	return count
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
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(in io.Reader, out io.Writer) {
	var bs = NewBufScanner(in)
	var bw = NewBufWriter(out)
	defer bw.w.Flush()

	bs.IntScan()
	s := bs.IntSliceScan()

	var count int
	for calc(s) {
		count++
	}

	bw.Printf("%d\n", count)
}

func calc(s []int) bool {
	for i := 0; i < len(s); i++ {
		if s[i] % 2 == 1 {
			return false
		}
		s[i] = s[i] / 2
	}
	return true
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

// IntScan Scan Data
func (b *BufScanner) IntScan() int {
	v, err := strconv.Atoi(b.Scan())
	if err != nil {
		panic(err)
	}
	return v
}

// SliceScan Scan Slice
func (b *BufScanner) SliceScan() []string {
	return strings.Split(b.Scan(), " ")
}

// IntSliceScan Scan Slice
func (b *BufScanner) IntSliceScan() []int {
	s := b.SliceScan()
	a := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		v, err := strconv.Atoi(s[i])
		if err != nil {
			panic(err)
		}
		a = append(a, v)
	}
	return a
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
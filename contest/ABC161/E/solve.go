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
	bs := NewBufScanner(in)
	bw := NewBufWriter(out)
	defer bw.w.Flush()
	/*
	前から数えた時と後ろから数えたとき一致する箇所が必須になる
	*/
	n, k, c := bs.IntScan(), bs.IntScan(), bs.IntScan()
	s := strings.Split(bs.Scan(), "")

	l := make([]string, n)
	r := make([]string, n)

	var index int
	var lCount int
	lPos := []int{}
	var rCount int

	index = 0
	for {
		if s[index] == "o" {
			l[index] = "o"
			lCount++
			lPos = append(lPos, index)
			index += c
		}
		index++
		if index >= n {
			break
		}
	}

	index = n - 1
	for {
		if s[index] == "o" {
			ok := true
			for i:=1; i<=c; i++ {
				if index + i >= n {
					break
				}
				if r[index+i] == "o" {
					ok = false
					break
				}
			}
			if ok {
				rCount++
				r[index] = "o"
			}
		}
		index--
		if index < 0 {
			break
		}
	}

	if lCount == rCount && lCount == k {
		for _, i := range lPos {
			if l[i] == "o" && l[i] == r[i] {
				bw.Printf("%v\n", i+1)
			}
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
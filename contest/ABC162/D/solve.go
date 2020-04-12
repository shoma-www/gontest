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
	n, s := bs.IntScan(), bs.Scan()
	type box struct {
		index int
		color string
	}
	r := make(map[int]int)
	rcount := 0
	b := make(map[int]int)
	bcount := 0
	g := make(map[int]int)
	gcount := 0

	for i:=1; i<=n; i++ {
		switch(s[i-1]){
		case 'R':
			r[i] = 0
			rcount++
		case 'B':
			b[i] = 0
			bcount++
		case 'G':
			g[i] = 0
			gcount++
		}
	}

	for i:=1; i<=n; i++ {
		switch(s[i-1]){
		case 'R':
			r[i] = rcount
			rcount--
		case 'B':
			b[i] = bcount
			bcount--
		case 'G':
			g[i] = gcount
			gcount--
		}
	}

	if len(r) == 0 || len(b) == 0 || len(g) == 0 {
		bw.Printf("0")
		return
	}

	for i:=1; i<=(n-2); i++ {
		
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
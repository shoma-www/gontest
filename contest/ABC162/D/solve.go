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

	rcount := 0
	bcount := 0
	gcount := 0

	for i:=1; i<=n; i++ {
		switch(s[i-1]){
		case 'R':
			rcount++
		case 'B':
			bcount++
		case 'G':
			gcount++
		}
	}

	if rcount == 0 || gcount == 0 || bcount == 0 {
		bw.Printf("0")
		return
	}

	var sum int64 = int64(rcount) * int64(gcount) * int64(bcount)
	for i:=0; i<n-2; i++ {
		for j:=i; i<n-1; j++ {
			k := 2*j-i
			if k >= n {
				break
			}
			if s[i] != s[j] && s[j] != s[k] && s[i] != s[k] {
				sum--
			}
		}
	}

	bw.Printf("%v", sum)
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
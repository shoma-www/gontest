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
	n, k := bs.IntScan(), bs.IntScan()

	var mod int64 = 1000000000 + 7
	d := make([]int64, k + 1)
	for i:=1; i<=k; i++ {
		d[i] = modpow(int64(k/i), mod, n)
	}
	for i:=k; i>=1; i-- {
		for j:=2*i; j<=k; j+=i {
			d[i] -= d[j]
			d[i] = (d[i] + mod) % mod
		}
	}

	var ans int64
	for i:=1; i<=k; i++ {
		ans += int64(i) * d[i] % mod
		ans = (ans + mod) % mod
	}
	bw.Printf("%v\n", int(ans))
}

func modpow(x, mod int64, n int) int64 {
	var res int64 = 1
	for n > 0 {
		if n & 1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		n >>= 1
	}
	return res
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
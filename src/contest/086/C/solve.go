package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	N, _ := strconv.Atoi(bs.Scan())

	nowP := NewPlace(0, 0, 0)

	result := "Yes"

	for i := 0; i < N; i++ {
		line := strings.Split(bs.Scan(), " ")
		t, _ := strconv.Atoi(line[0])
		x, _ := strconv.Atoi(line[1])
		y, _ := strconv.Atoi(line[2])
		nextP := NewPlace(t, x, y)

		count := nextP.t - nowP.t
		move := abs(nextP.x - nowP.x) + abs(nextP.y - nowP.y)

		if move > count || move % 2 != count % 2 {
			result = "No"
			break
		}

		nowP = nextP
	}

	bw.Printf("%s\n", result)
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

type Place struct {
	t int
	x int
	y int
}

func NewPlace(t, x, y int) *Place {
	return &Place{
		t: t,
		x: x,
		y: y,
	}
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
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

var bs *BufScanner
var bw *BufWriter

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func solve(in io.Reader, out io.Writer) {
	bs = NewBufScanner(in)
	bw = NewBufWriter(out)
	// defer bw.w.Flush()

	n, q := getLine()

	var ans []int
	if q == 7 {
		ans = solve7(n)
	} else {
		ans = solve100(n)
	}

	s := ""
	for _, a := range ans {
		s += string(letters[a])
	}

	bw.Printf("! %s\n", s)
}

func solve100(n int) []int {
	l := make([]int, 1, n)
	for i := 1; i < n; i++ {
		start := 0
		end := len(l)
		for {
			idx := start + (end - start) / 2
			t := query(i, l[idx])
			if t > 0 {
				if end - start <= 2  {
					l = append(l[:end + 1], l[end:]...)
					l[end] = i
					break
				}
				start = idx + 1
			} else {
				if end - start == 1  {
					l = append(l[:idx + 1], l[idx:]...)
					l[idx] = i
					break
				}
				end = idx
			}
		}
	}

	return l
}

func solve7(n int) []int {
	pairs := createPairs(n)
	full := createFull(n)
	return filter(pairs, full)
}

func filter(pairs [][]int, current [][]int) []int {
	if len(current) <= 1 {
		return current[0]
	}
	min := math.MaxInt32
	idx := -1
	for i, pair := range pairs {
		count := 0
		for _, c := range current {
			if searchIndex(c, pair[0]) - searchIndex(c, pair[1]) > 0 {
				count++
			}
		}

		v := len(current) / 2 - count
		if v < 0 {
			v = v * -1
		}
		if min > v {
			idx = i
			min = v
		}
	}

	pair := pairs[idx]
	t := query(pair[0], pair[1])

	return filter(
		pairFilter(pairs, pairs[idx], func(a []int, p []int) bool {return searchIndex(a, p[0]) < 0 || searchIndex(a, p[1]) < 0}),
		pairFilter(current, pairs[idx], func(a []int, p []int) bool {return (searchIndex(a, p[0]) - searchIndex(a, p[1])) * t > 0}),
	)
}

func getLine() (int, int) {
	s := bs.Scan()
	line := strings.Split(s, " ")
	n, _ := strconv.Atoi(line[0])
	q, _ := strconv.Atoi(line[1])
	return n, q
}

func createFull(n int) (result [][]int) {
	list := make([]int, 0, n)
	for i := 0; i < n; i++ {
		list = append(list, i)
	}
	return innerCreateFull(list)
}

func innerCreateFull(a []int) (result [][]int) {
	if len(a) == 0 {
		result = append(result, []int{})
	}
	for i, v := range a {
		for _, r := range innerCreateFull(delete(i, a)) {
			result = append(result, append([]int{v}, r...))
		}
	}
	return
}

func delete(idx int, L []int) (result []int) {
	result = append(result, L[:idx]...)
	result = append(result, L[idx+1:]...)
	return
}

func createPairs(n int) [][]int {
	pairs := make([][]int, 0, n * (n - 1) / 2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pairs = append(pairs, []int{i, j})
		}
	}
	return pairs
}

func searchIndex(a []int, t int) int {
	for i, v := range a {
		if v == t {
			return i
		}
	}
	return -1
}

func query(l, r int) int {
	bw.Printf("? %s %s\n", string(letters[l]), string(letters[r]))
	switch(bs.Scan()) {
	case ">":
		return 1
	case "<":
		return -1
	}
	panic("wow")
}

func pairFilter(a [][]int, pair []int, f func([]int, []int) bool) [][]int {
	var newSlice [][]int
	for _, p := range a {
		if f(p, pair) {
			newSlice = append(newSlice, p)
		}
	}
	return newSlice
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
	w io.Writer
}

// NewBufWriter constructer
func NewBufWriter(out io.Writer) *BufWriter {
	return &BufWriter{
		w: out,
	}
}

// Printf Output file
func (b *BufWriter) Printf(format string, a ...interface{}) {
	fmt.Fprintf(b.w, format, a...)
}
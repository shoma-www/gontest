package tree

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestUnionFindTree(t *testing.T) {
	table := []struct {
		n int
		m int
		lines []string
		expected int
	}{
		{
			n: 5,
			m: 2,
			lines: []string{
				"5 3 1 4 2",
				"1 3",
				"5 4",
			},
			expected: 2,
		},
		{
			n: 3,
			m: 2,
			lines: []string{
				"3 2 1",
				"1 2",
				"2 3",
			},
			expected: 3,
		},
	}

	for _, test := range table {
		r := solve(test.n, test.m, test.lines)
		if test.expected != r {
			t.Errorf("expected: %v actual: %v\n", test.expected, r)
		}
	}
}

func solve(n, m int, lines []string) int {
	ps := strings.Split(lines[0], " ")
	p := make([]int, n)
	for i, v := range ps {
		p[i], _ = strconv.Atoi(v)
	}
	
	ut := InitUnionFindTree(n)
	
	for i := 1; i <= m; i++ {
		pair := strings.Split(lines[i], " ")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		ut.Unite(x - 1, y - 1, 0)
	}

	fmt.Printf("%v %v %v\n", p, ut.parent, ut.rank)

	count := 0

	for i := 0; i < n; i++ {
		if ut.Same(i, p[i] - 1) {
			count++
		}
	}

	return count
}

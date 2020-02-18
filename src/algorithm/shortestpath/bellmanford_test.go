package shortestpath

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestBellmanFord(t *testing.T) {
	testCase := []struct{
		edges []string
		vertex int
		start int
		end int
		exp []int
	}{
		{
			edges: []string{
				"0 1 3",
				"0 2 4",
				"1 4 7",
				"1 3 2",
				"2 3 -1",
				"3 4 3",
				"3 5 2",
				"5 4 -1",
				"4 6 3",
			},
			vertex: 7,
			start: 0,
			end: 6,
			exp: []int{ 0, 2, 3, 5, 4, 6 },
		},
	}

	for _, c := range testCase {
		edges := make(Edges, 0, len(c.edges))
		for _, e := range c.edges {
			a := strings.Split(e, " ")
			from, _ := strconv.Atoi(a[0])
			to, _ := strconv.Atoi(a[1])
			cost, _ := strconv.Atoi(a[2])
			edge := Edge{
				from: from,
				to: to,
				cost: cost,
			}
			edges = append(edges, edge)
		}
		ans, err := BellmanFord(edges, c.vertex, c.start, c.end)
		fmt.Println(ans)

		if err != nil {
			t.Errorf("dame %v.", err)
		}
		for i := range ans {
			if c.exp[i] != ans[i] {
				t.Errorf("dame %v %v.", c.exp[i], ans[i])
			}
		}
	}
}
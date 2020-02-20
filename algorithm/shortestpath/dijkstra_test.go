package shortestpath

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestDijkstra(t *testing.T) {
	testCase := []struct{
		edges []string
		vertex int
		start int
		end int
		exp []int
	}{
		{
			edges: []string{
				"0 1 2",
				"0 2 4",
				"1 4 7",
				"1 3 3",
				"2 3 2",
				"2 5 2",
				"3 4 3",
				"4 5 1",
				"4 6 2",
				"5 6 3",
			},
			vertex: 7,
			start: 0,
			end: 6,
			exp: []int{ 0, 2, 5, 4, 6 },
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
		ans, err := Dijkstra(edges, c.vertex, c.start, c.end)
		fmt.Println(ans)

		if err != nil {
			t.Errorf("dame %v.", err)
		}
		if len(c.exp) != len(ans) {
			t.Errorf("dame len exp:%v ans:%v .", c.exp ,ans)
		}
		for i := range ans {
			if c.exp[i] != ans[i] {
				t.Errorf("dame %v %v.", c.exp[i], ans[i])
			}
		}
	}
}
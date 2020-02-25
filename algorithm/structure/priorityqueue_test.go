package structure

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	type pair struct{
		value string
		priority int
	}
	testCase := []struct{
		values []pair
		expect []pair
	}{
		{
			values: []pair{
				pair{
					value: "foo",
					priority: 5,
				},
				pair{
					value: "hoge",
					priority: 1,
				},
				pair{
					value: "var",
					priority: 7,
				},
				pair{
					value: "fuga",
					priority: 3,
				},
			},
			expect: []pair{
				pair{
					value: "var",
					priority: 7,
				},
				pair{
					value: "foo",
					priority: 5,
				},
				pair{
					value: "fuga",
					priority: 3,
				},
				pair{
					value: "hoge",
					priority: 1,
				},
			},
		},
	}

	for _, c := range testCase {
		q := NewPriorityQueue()
		for _, v := range c.values {
			q.Push(&Item{
				Value: v.value,
				Priority: v.priority,
			})
		}

		i := 0
		for !q.Empty() {
			item := q.Pop()
			if item.Priority != c.expect[i].priority || item.Value != c.expect[i].value {
				t.Errorf("not match exp:%v act:%v", c.expect[i], item)
			}
			i++
		}
	}
}
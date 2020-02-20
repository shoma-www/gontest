package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	fmt.Println("########## Start Test. ##########")
	dir, _ := filepath.Abs("./testdata")
	tb := NewTestBook(dir)

	all, success := 0, 0
	for _, testName := range tb.GetTestList() {
		fmt.Printf("%s:\n", testName)
		all++
		if InteractiveTestExecute(t, testName, tb[testName], ans, 2000, '!') {
			success++
		}
	}

	fmt.Printf("All: %d Success: %d Error: %d\n", all, success, all-success)
	fmt.Println("########## Finish Test. ##########")
}

var base = ""

func ans(buf *SyncBuffer, v *TestPath, q string) {
	l := strings.Split(q, " ")
	if len(l) != 3 {
		panic("しつもんまちがえとる！！！！！！！")
	}
	if l[0] != "?" {
		panic("最初は？じゃけええええ")
	}

	if string(l[1]) == "N" {
		fmt.Print(strings.Trim(l[2], "\n"))
	}

	if base == "" {
		fp, _ := os.Open(v.mid)
		r := bufio.NewReader(fp)
		b, _, _ := r.ReadLine()
		base = string(b)
	}

	s := "<"
	if getIndex(base, l[1]) - getIndex(base, strings.Trim(l[2], "\n")) > 0 {
		s = ">"
	}
	buf.Write([]byte(s))
}

func getIndex(a string, t string) int {
	for i := 0; i < len(a); i++ {
		if string(a[i]) == t {
			return i
		}
	}
	return -1
}

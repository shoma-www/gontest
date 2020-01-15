package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"
)

const (
	LIMIT = 2000
)

type TestPath struct {
	in  string
	exp string
}

func NewTestPath() *TestPath {
	return &TestPath{}
}

func TestSolve(t *testing.T) {
	fmt.Println("########## Start Test. ##########")
	dir, _ := filepath.Abs("./testdata")
	m := getTests(dir)
	keys := getKeys(m)

	all := 0
	success := 0
	for _, k := range keys {
		fmt.Printf("%s:\n", k)
		all++
		if execute(t, k, m[k]) {
			success++
		}
	}

	fmt.Printf("All: %d Success: %d Error: %d\n", all, success, all-success)
	fmt.Println("########## Finish Test. ##########")
}

func getTests(dir string) map[string]*TestPath {
	files, _ := ioutil.ReadDir(dir)
	m := make(map[string]*TestPath)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		p := file.Name()
		k := strings.Split(filepath.Base(p), ".")[0]

		if _, ok := m[k]; !ok {
			m[k] = NewTestPath()
		}

		switch filepath.Ext(p) {
		case ".in":
			m[k].in = filepath.Join(dir, p)
			break
		case ".out":
			m[k].exp = filepath.Join(dir, p)
			break
		}
	}

	return m
}

func execute(t *testing.T, k string, v *TestPath) (ok bool) {
	in, _ := os.Open(v.in)
	out := &bytes.Buffer{}
	defer in.Close()

	ctx, cancel := context.WithTimeout(context.Background(), LIMIT * time.Millisecond)
	defer cancel()
	c := make(chan string)
	go func() {
		s := time.Now()
		solve(in, out)
		t := time.Since(s)
		log(k, fmt.Sprintf("Elapsed Time %2.8f sec", t.Seconds()))
		c <- "finish"
	}()

	select {
	case <- c:
		break
	case <-ctx.Done():
		errorLog(t, k, "Time out !!")
		return false
	}

	line := 0
	ok = true
	mismatch := false

	actual := bufio.NewReader(out)
	file, _ := os.Open(v.exp)
	expected := bufio.NewReader(file)

	for {
		a, _, aErr := actual.ReadLine()
		e, _, eErr := expected.ReadLine()
		line++

		if aErr == io.EOF && eErr == io.EOF {
			if !mismatch {
				log(k, "Accepted!!")
			}
			break
		} else if aErr == io.EOF {
			ok = false
			errorLog(t, k, "Execution result is less than expected.")
			break
		} else if eErr == io.EOF {
			ok = false
			errorLog(t, k, "Execution result is more than expected.")
			break
		}

		if string(e) != string(a) {
			if !mismatch {
				ok = false
				mismatch = true
				errorLog(t, k, "Do not match expected value.")
			}
			fmt.Printf(getSpaces(len(k)+4)+"line:%d expected=%s actual=%s\n", line, string(e), string(a))
		}
	}

	return ok
}

func errorLog(t *testing.T, k, s string) {
	t.Errorf("%s: "+s, k)
	log(k, s)
}

func log(k, s string) {
	fmt.Print(getSpaces(len(k)) + s + "\n")
}

func getSpaces(n int) string {
	return strings.Repeat(" ", n)
}

func getKeys(m map[string]*TestPath) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"
)

// TestPath テスト用
type TestPath struct {
	in  string
	exp string
	mid string
}

// TestBook テスト一覧
type TestBook map[string]*TestPath

// NewTestBook テストのファイルパスのマップを取得する
func NewTestBook(dir string) TestBook {
	files, _ := ioutil.ReadDir(dir)
	m := make(TestBook)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		p := file.Name()
		k := strings.Split(filepath.Base(p), ".")[0]

		if _, ok := m[k]; !ok {
			m[k] = &TestPath{}
		}

		switch filepath.Ext(p) {
		case ".in":
			m[k].in = filepath.Join(dir, p)
			break
		case ".out":
			m[k].exp = filepath.Join(dir, p)
			break
		case ".mid":
			m[k].mid = filepath.Join(dir, p)
			break
		}
	}

	return m
}

// GetTestList テスト一覧を取得
func (tb TestBook) GetTestList() []string {
	keys := make([]string, 0, len(tb))
	for k := range tb {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// ログ関係

func getSpaces(n int) string {
	return strings.Repeat(" ", n)
}

func errorLog(t *testing.T, k, s string) {
	t.Errorf("%s: "+s, k)
	log(k, s)
}

func log(k, s string) {
	fmt.Print(getSpaces(len(k)) + s + "\n")
}

func fmtGetFunc(n int) {
	pc, file, line, _ := runtime.Caller(n)
	f := runtime.FuncForPC(pc)
	fmt.Printf("call:%s\nfile:%s:%d\n", f.Name(), file, line)
}

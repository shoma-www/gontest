package test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"
)

// Path テスト用
type Path struct {
	in  string
	exp string
	mid string
}

// Book テスト一覧
type Book map[string]*Path

// NewTestBook テストのファイルパスのマップを取得する
func NewTestBook(dir string) Book {
	files, _ := ioutil.ReadDir(dir)
	tb := make(Book)

	for _, file := range files {
		if file.IsDir() {
			dirName := file.Name()
			testPath := filepath.Join(dir, dirName)
			testFiles, _ := ioutil.ReadDir(testPath)

			for _, test := range testFiles {
				p := test.Name()

				if _, ok := tb[dirName]; !ok {
					tb[dirName] = &Path{}
				}

				switch p {
				case "in":
					tb[dirName].in = filepath.Join(testPath, p)
					break
				case "out":
					tb[dirName].exp = filepath.Join(testPath, p)
					break
				case "mid":
					tb[dirName].mid = filepath.Join(testPath, p)
					break
				}
			}
		}
	}
	return tb
}

// GetTestList テスト一覧を取得
func (tb Book) GetTestList() []string {
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

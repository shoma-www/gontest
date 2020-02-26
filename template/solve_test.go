package main

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/shoma-www/my-contest-go/common/test"
)

func TestSolve(t *testing.T) {
	fmt.Println("########## Start Test. ##########")
	dir, _ := filepath.Abs("./testdata")
	tb := test.NewTestBook(dir)

	all, success := 0, 0
	for _, testName := range tb.GetTestList() {
		fmt.Printf("%s:\n", testName)
		all++

		// 通常
		// if test.Execute(t, testName, tb[testName], 2000) {
		// 	success++
		// }

		// 対話式
		// if test.InteractiveTestExecute(t, testName, tb[testName], ans, 2000, '!') {
		// 	success++
		// }
	}

	fmt.Printf("All: %d Success: %d Error: %d\n", all, success, all-success)
	fmt.Println("########## Finish Test. ##########")
}

func ans(buf *SyncBuffer, v *TestPath, q string) {

}
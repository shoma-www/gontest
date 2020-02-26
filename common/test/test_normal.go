package test

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

// Execute 通常テスト
func Execute(t *testing.T, solve func(io.Reader, io.Writer), testName string, v *TestPath, limit int) (ok bool) {
	in, _ := os.Open(v.in)
	out := &bytes.Buffer{}
	defer in.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(limit * int(time.Millisecond)))
	defer cancel()
	exeCh := make(chan string)
	go func() {
		s := time.Now()
		solve(in, out)
		t := time.Since(s)
		log(testName, fmt.Sprintf("Elapsed Time %2.8f sec", t.Seconds()))
		exeCh <- "finish"
	}()

	// 実行時間制限
	select {
	case <- exeCh:// 実行完了
		break
	case <-ctx.Done():// タイムアウト
		errorLog(t, testName, "Time out !!")
		return false
	}

	// Asertやで
	line := 0
	ok, mismatch := true, false

	actual := bufio.NewReader(out)
	file, _ := os.Open(v.exp)
	expected := bufio.NewReader(file)

	for {
		a, _, aErr := actual.ReadLine()
		e, _, eErr := expected.ReadLine()
		line++

		if aErr == io.EOF && eErr == io.EOF {
			if !mismatch {
				log(testName, "Accepted!!")
			}
			break
		} else if aErr == io.EOF {
			ok = false
			errorLog(t, testName, "Execution result is less than expected.")
			break
		} else if eErr == io.EOF {
			ok = false
			errorLog(t, testName, "Execution result is more than expected.")
			break
		}

		if string(e) != string(a) {
			if !mismatch {
				ok = false
				mismatch = true
				errorLog(t, testName, "Do not match expected value.")
			}
			fmt.Printf(getSpaces(len(testName)+4)+"line:%d expected=%s actual=%s\n", line, string(e), string(a))
		}
	}

	return ok
}

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

// Answer 対話用の出力を生成する関数定義
type Answer func(buf *SyncBuffer, v *TestPath, q string)

// InteractiveTestExecute 対話型実行テスト
func InteractiveTestExecute(t *testing.T, testName string, v *TestPath, answer Answer, limit int, endWord byte) (ok bool) {
	in := NewSyncBuffer()
	out := NewSyncBuffer()
	exeCh := make(chan string)

	var act, exp []byte
	var actErr, expErr error

	// 標準入出力
	go func() {
		f, _ := os.Open(v.in)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Scan()
		s := scanner.Text()
		in.Write([]byte(s))
		for {
			b := make([]byte, 4096, 4096)
			n, _ := out.Read(b)
			if bytes.IndexByte(b[:n], endWord) >= 0 {
				act = bytes.Trim(b[:n], "\n")
				actErr = nil
				break
			}
			q := string(b[:n])
			answer(in, v, q)
		}
	}()

	// 関数実行をタイムアウトさせる
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(limit * int(time.Millisecond)))
	defer cancel()

	// 実行本体
	go func() {
		s := time.Now()
		solve(in, out)
		t := time.Since(s)
		log(testName, fmt.Sprintf("Elapsed Time %2.8f sec", t.Seconds()))
		// 実行終了を通知
		exeCh <- "finish"
	}()

	// 実行時間制限
	select {
	case <-exeCh:// 実行完了
		break
	case <-ctx.Done():// タイムアウト
		errorLog(t, testName, "Time out !!")
		return false
	}

	// Asertやで
	line := 0
	ok, mismatch := true, false

	actual := bufio.NewReader(out.buf)
	file, _ := os.Open(v.exp)
	expected := bufio.NewReader(file)

	exp, _, expErr = expected.ReadLine()

	for {
		line++

		if actErr == io.EOF && expErr == io.EOF {
			if !mismatch {
				log(testName, "Accepted!!")
			}
			break
		} else if expErr == io.EOF {
			ok = false
			errorLog(t, testName, "Execution result is more than expected.")
			break
		} else if actErr == io.EOF {
			ok = false
			errorLog(t, testName, "Execution result is more than actual.")
			break
		}

		if string(exp) != string(act) {
			if !mismatch {
				ok = false
				mismatch = true
				errorLog(t, testName, "Do not match expected value.")
			}
			fmt.Printf(getSpaces(len(testName)+4) + "line:%d expected=%s actual=%s\n", line, string(exp), string(act))
		}

		act, _, actErr = actual.ReadLine()
		exp, _, expErr = expected.ReadLine()
	}

	return ok
}


// SyncBuffer 同期バッファー
type SyncBuffer struct {
	buf *bytes.Buffer
	rc chan string
	wc chan string
}

// NewSyncBuffer コンストラクタ
func NewSyncBuffer() *SyncBuffer {
	return &SyncBuffer{
			buf: &bytes.Buffer{},
			rc: make(chan string),
			wc: make(chan string),
	}
}

// Read 同期バッファーの読み込み
func (b *SyncBuffer) Read(p []byte) (n int, err error) {
	// 書き込まれるまで待機
	<-b.wc
	n, err = b.buf.Read(p)
	// 読み込めたことを通知
	b.rc<-"r"
	return
}

// Write 同期バッファーへの書き込み
func (b *SyncBuffer) Write(p []byte) (n int, err error) {
	p = bytes.Trim(p, "\n")
	n, err = b.buf.Write(append(p, '\n'))
	// 書き込めたことを通知
	b.wc<-"w"
	// 同期させるため読み込まれるまで待機
	<-b.rc
	return
}

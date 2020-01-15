package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(in io.Reader, out io.Writer) {
	var s = bufio.NewScanner(in)
	var w = bufio.NewWriter(out)

	// ##### ここから #####
	time.Sleep(time.Second * 1)
	var text string

	for s.Scan() {
		text = s.Text()
		fmt.Fprintln(w, text)
	}
	// ##### ここまでを変更 #####

	defer w.Flush()
}

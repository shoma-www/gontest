package tree

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestRedBlackTree(t *testing.T) {
	bt := NewRedBlackTree(10, "")
	bt.Add(8, "")
	bt.Add(11, "")
	bt.Add(14, "")
	bt.Add(13, "")
	bt.Add(18, "")
	bt.Add(5, "")
	bt.Add(1, "")
	bt.Add(30, "")
	p := newRBPriter()
	p.print(bt.Root)
}

type rbprinter struct {
	buf [][]string
}

func newRBPriter() *rbprinter {
	return &rbprinter{[][]string{}}
}

func (p *rbprinter) print(bn *RedBlackNode) {
	p.createPrintData(bn, 0, 0)
	for i:=0; i<len(p.buf); i++ {
		for j:=0; j<len(p.buf[i]); j++ {
			if strings.Index(p.buf[i][j], "R") > 0 {
				fmt.Printf("\x1b[31m%v", p.buf[i][j])
			} else {
				fmt.Printf("\x1b[30m%v", p.buf[i][j])
			}
		}
		fmt.Println()
	}
}

func (p *rbprinter)createPrintData(bn *RedBlackNode, d, w int) {
	if len(p.buf) == d {
		tmp := make([]string, int(math.Pow(2, float64(d))))
		for i:=0; i<len(tmp); i++ {
			tmp[i] = "     "
		}
		p.buf = append(p.buf, tmp)
	}
	p.buf[d][w] = addChar(fmt.Sprintf("%v%v", bn.Key, bn.Color))
	if bn.Left != nil {
		p.createPrintData(bn.Left, d+1, w*2)
	}
	if bn.Right != nil {
		p.createPrintData(bn.Right, d+1, w*2+1)
	}
}

func addChar(str string) string {
	return str + strings.Repeat(" ", 5 - len(str))
}
package tree

// RedBlackTree 赤黒木の構造体
type RedBlackTree struct {
	Root *RedBlackNode
}

// NewRedBlackTree 赤黒木のコンストラクタ
func NewRedBlackTree(key int, value string) *RedBlackTree {
	return &RedBlackTree{NewRedBlackNode(BLACK, key, value)}
}

// Add 赤黒木に追加する
func (bt *RedBlackTree) Add(key int, value string) {
	if bt.Root == nil {
		bt.Root = NewRedBlackNode(BLACK, key, value)
	} else {
		bt.Root.Add(key, value)
		bt.Root.Color = BLACK
	}
}

// Color 色
type Color string

const (
	// RED 赤
	RED = Color("R")
	// BLACK 黒
	BLACK = Color("B")
)

func isRed(bm *RedBlackNode) bool {
	return bm != nil && bm.Color == RED
}

// RedBlackNode 赤黒木のノード構造体
type RedBlackNode struct {
	Color Color
	Key   int
	Value string
	Left  *RedBlackNode
	Right *RedBlackNode
}

// NewRedBlackNode 赤黒木のノードのコンストラクタ
func NewRedBlackNode(color Color, key int, value string) *RedBlackNode {
	return &RedBlackNode{color, key, value, nil, nil}
}

// Set 引数の値をコピー
func (bm *RedBlackNode) Set(v *RedBlackNode) {
	bm.Color = v.Color
	bm.Key = v.Key
	bm.Value = v.Value
	bm.Left = v.Left
	bm.Right = v.Right
}

// Add 自身の値より小さければ左ノードに追加、それ以外は右ノードに追加する
func (bm *RedBlackNode) Add(key int, value string) bool {
	if bm.Key == key {
		bm.Value = value
		return false
	}

	var isFix bool
	if bm.Key > key {
		if bm.Left == nil {
			bm.Left = NewRedBlackNode(RED, key, value)
			return true
		}
		isFix = bm.Left.Add(key, value)
	} else {
		if bm.Right == nil {
			bm.Right = NewRedBlackNode(RED, key, value)
			return true
		}
		isFix = bm.Right.Add(key, value)
	}

	if isFix {
		if isRed(bm) {
		} else if isRed(bm.Left) && isRed(bm.Left.Left) {
			bm.RotateR()
			bm.Left.Color = BLACK
		} else if isRed(bm.Left) && isRed(bm.Left.Right) {
			bm.RotateLR()
			bm.Left.Color = BLACK
		} else if isRed(bm.Right) && isRed(bm.Right.Left) {
			bm.RotateRL()
			bm.Right.Color = BLACK
		} else if isRed(bm.Right) && isRed(bm.Right.Right) {
			bm.RotateL()
			bm.Right.Color = BLACK
		} else {
			isFix = false
		}
	}
	return isFix
}

// RotateR 自分を起点に右回転させる
func (bm *RedBlackNode) RotateR() bool {
	if bm.Left == nil {
		return false
	}
	tmp := *bm
	tmp.Left = bm.Left.Right
	bm.Set(bm.Left)
	bm.Right = &tmp
	return true
}

// RotateL 自分を起点に左回転させる
func (bm *RedBlackNode) RotateL() bool {
	if bm.Right == nil {
		return false
	}
	tmp := *bm
	tmp.Right = bm.Right.Left
	bm.Set(bm.Right)
	bm.Left = &tmp
	return true
}

// RotateRL ２重回転。右ルートの平衡化
func (bm *RedBlackNode) RotateRL() bool {
	return bm.Right.RotateR() && bm.RotateL()
}

// RotateLR ２重回転。左ルートの平衡化
func (bm *RedBlackNode) RotateLR() bool {
	return bm.Left.RotateL() && bm.RotateR()
}

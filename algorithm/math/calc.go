package math

import "errors"

// Euclidean ユークリッドの互除法
// 割り算の等式：a=bq+r において，
// 「a と b の最大公約数」＝「b と r の最大公約数」
// という性質が成り立つ
// bとrに対して繰り返し上記の等式を繰り返し、
// あまりがゼロになるまで繰り返せば、bが最大公約数となる
func Euclidean(a, b int) (int, error) {
	if a < b {
		return -1, errors.New("b has to be bigger than a")
	}
	if b == 0 {
		return a, nil
	}
	return Euclidean(b, a % b)
}
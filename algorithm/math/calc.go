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

// Modpow 累乗のあまりを計算する
// 二分累乗法を用いる
// 計算量がO(log n)で求まる
// 指数を二進法展開をして、乗算の回数を圧縮する
// 3^45 ⇒ 指数が45⇒2^0+2^2+2^3+2^5
// 各べき乗で3のべき乗を計算すれば、3^45も求められる
func Modpow(x, n, mod int64) int64 {
	var r int64 = 1
	for n > 0 {
		if n & 1 == 1 {
			r = r * x % mod
		}
		x = x * x % mod
		n >>= 1
	}
	return r
}
package math

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

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

// AddBigInt 大きな数値の足し算を行う
// 和と差は、ひっ算が一番効率よく計算できる
// O(N)
func AddBigInt(x, y string) (string, error) {
	return convertBigIntFunc(addBigInt)(x, y)
}

// SubBigInt 大きな数値の引き算を行う
// O(N)
func SubBigInt(x, y string) (string, error) {
	return convertBigIntFunc(subBigInt)(x, y)
}

// MultiBigInt 大きな数値の掛け算を行う
// O(N^2)
func MultiBigInt(x, y string) (string, error) {
	return convertBigIntFunc(multiBigInt)(x, y)
}

// DivBigInt 大きな数値の割り算を行う
func DivBigInt(x, y string) (string, error) {
	return convertBigIntFunc(divBigInt)(x, y)
}

// KaratsubaMethod カラツバ法
// N桁×N桁のオーダーからN/2桁×N/2桁×3の計算量に圧縮
// ひっ算を用いた掛け算の計算と組み合わせることで高速化が可能
// O(N^1.59)
// 例: 20201215 × 12345678 を、4桁 × 4桁に落とし込んだ場合
// 2020 × 1234 × 10^8
// + (1215 × 1234 + 2020 × 5678) × 10^4
// + 1215 × 5678
func KaratsubaMethod(x, y string) (string, error) {
	return convertBigIntFunc(multiKaratsuba)(x, y)
}

type bigInt []int

// stringTobigInt 数値の文字列を受け取って、各桁ごとの配列を返却する
// in:"12345" ⇒ out:[5, 4, 3, 2, 1]
func stringTobigInt(s string) (bigInt, error) {
	base := 1
	if strings.Contains(s, "-") {
		s = s[1:]
		base = -1
	}
	n := len(s)
	d := make(bigInt, n)
	var err error
	for i, c := range s {
		d[n-i-1], err = strconv.Atoi(string(c))
		if err != nil {
			return nil, err
		}
		d[n-i-1] *= base
	}
	return d[:n], nil
}

// bigIntTostring bigIntを受け取って、数値の文字列を返却する
// in:[5, 4, 3, 2, 1] ⇒ out:"12345"
func bigIntTostring(b bigInt) string {
	n := len(b)
	s := ""
	if n == 0 {
		return s
	}
	isM := false
	s += strconv.Itoa(b[n-1])
	if len(s) > 1 {
		isM = true
	}
	for i := n-2; i >= 0; i-- {
		s += strconv.Itoa(b[i])
	}
	if isM {
		tmp := s
		s = strings.Replace(tmp, "-", "", -1)
		s = "-" + s
	}
	return s
}

func addBigInt(x, y bigInt) bigInt {
	fn := func(a, b int) int {
		return a + b
	}
	return addSubBigInt(fn)(x, y)
}

func subBigInt(x, y bigInt) bigInt {
	fn := func(a, b int) int {
		return a - b
	}
	return addSubBigInt(fn)(x, y)
}

func multiBigInt(x, y bigInt) bigInt {
	nx, ny := len(x), len(y)
	digitsAns := make(bigInt, nx + ny - 1)

	for i := 0; i < nx; i++ {
		for j := 0; j < ny; j++ {
			digitsAns[i+j] += x[i] * y[j]
		}
	}
	return carryAndFix(digitsAns)
}

func divBigInt(x, y bigInt) bigInt {
	d := len(x) - len(y)
	if d < 0 { return bigInt{0} }
	if compareBigInt(x[d:], y) >= 0 {
		d++
	}
	if d == 0 { return bigInt{0} }

	remain := x[d - 1:]
	ans := make(bigInt, d)
	for i := d - 1; i >= 0; i-- {
		ans[i] = 9
		for j := 1; j <= 9; j++ {
			tmp := multiBigInt(y, bigInt{j})
			if compareBigInt(tmp, remain) == 1{
				ans[i] = j - 1
				break
			}
		}
		tmpResult := multiBigInt(y, bigInt{ans[i]})
		remain = subBigInt(remain, tmpResult)
		if i >= 1 {
			remain = append(bigInt{x[i - 1]}, remain...)
		}
	}
	return ans
}

func compareBigInt(x, y bigInt) int {
	nx, ny := len(x), len(y)
	if nx > ny { return 1 }
	if nx < ny { return -1 }
	for i := nx - 1; i >= 0; i-- {
		if x[i] > y[i] { return 1 }
		if x[i] < y[i] { return -1 }
	}
	return 0
}

func multiKaratsuba(dx, dy bigInt) bigInt {
	if len(dx) < 3 || len(dy) < 3 {
		return multiBigInt(dx, dy)
	}
	n := len(dx) / 2
	tmp := make(bigInt, n)

	a := multiKaratsuba(dx[n:], dy[n:])
	b := multiKaratsuba(dx[:n], dy[:n])
	c := multiKaratsuba(addBigInt(dx[:n], dx[n:]), addBigInt(dy[:n], dy[n:]))
	c = subBigInt(subBigInt(c, a), b)

	a = append(append(tmp, tmp...), a...)
	c = append(tmp, c...)

	return addBigInt(addBigInt(a, b), c)
}

func addSubBigInt(fn func(a, b int) int) func(x, y bigInt) bigInt {
	return func(x, y bigInt) bigInt {
		digitsAns := append(bigInt{}, x...)
		if len(digitsAns) < len(y) {
			tmp := make(bigInt, len(y) - len(digitsAns))
			digitsAns = append(digitsAns, tmp...)
		} else if len(digitsAns) > len(y) {
			tmp := make(bigInt, len(digitsAns) - len(y))
			y = append(y, tmp...)
		}

		for i := 0; i < len(y); i++ {
			digitsAns[i] = fn(digitsAns[i], y[i])
		}

		return carryAndFix(digitsAns)
	}
}

func convertBigIntFunc(fn func(a, b bigInt) bigInt) func(x, y string)(string, error) {
	return func(x, y string) (string, error) {
		digitsX, err := stringTobigInt(x)
		if err != nil {
			return "", errors.New("x is invalid value")
		}
		digitsY, err := stringTobigInt(y)
		if err != nil {
			return "", errors.New("y is invalid value")
		}

		digitsAns := fn(digitsX, digitsY)
		return bigIntTostring(digitsAns), nil
	}
}

// carryAndFix bigIntに対して繰り上がり/繰り下がり処理を行う
func carryAndFix(digits bigInt) bigInt {
	// マイナス対応
	isM := digits[len(digits) - 1] < 0
	if isM {
		for i := 0; i < len(digits); i++ {
			digits[i] *= -1
		}
	}
	// 各桁に対して繰り上がり/繰り下がり処理
	for i := 0; i < len(digits) - 1; i++ {
		// 繰り上がり
		if digits[i] >= 10 {
			k := digits[i] / 10
			digits[i] -= k * 10
			digits[i+1] += k
		}

		// 繰り下がり
		if digits[i] < 0 {
			// -10のときは、10足せばいいので、-1している
			k := (-digits[i] - 1) / 10 + 1
			digits[i] += k * 10
			digits[i+1] -= k
		}
	}

	// 最上位の桁が10以上の場合、桁を上げる
	for digits[len(digits) - 1] >= 10 {
		k := digits[len(digits) - 1] / 10
		digits[len(digits) - 1] -= k * 10
		digits = append(digits, k)
	}

	// 最上位の桁がゼロの場合、桁を下げる
	for len(digits) > 1 && digits[len(digits) - 1] == 0 {
		digits = digits[:len(digits) - 1]
	}

	// マイナス対応
	if isM {
		for i := 0; i < len(digits); i++ {
			digits[i] *= -1
		}
	}

	return digits
}

// ToomCook Toom-Cook法
// 
// O(N^1.46)
func ToomCook(x, y string) (string, error) {
	return convertBigIntFunc(multiToomCook3)(x, y)
}

func multiToomCook3(dx, dy bigInt) bigInt {
	if len(dx) < 9 || len(dy) < 9 {
		return multiKaratsuba(dx, dy)
	}

	n := int(math.Round(float64(len(dx)) / 3))
	p := createEquations(dx, n)
	q := createEquations(dy, n)

	r2 := multiToomCook3(p("2"), q("2"))
	r1 := multiToomCook3(p("1"), q("1"))
	r0 := multiToomCook3(p("0"), q("0"))
	rm1 := multiToomCook3(p("-1"), q("-1"))
	rm2 := multiToomCook3(p("-2"), q("-2"))

	c0 := r0
	c1 := divBigInt(multiAddBigInt(rm2, multiInt("-8", rm1), multiInt("8", r1), multiInt("-1", r2)), bigInt{2, 1})
	c2 := divBigInt(multiAddBigInt(multiInt("-1", rm2), multiInt("16", rm1), multiInt("-30", r0), multiInt("16", r1), multiInt("-1", r2)), bigInt{4, 2})
	c3 := divBigInt(multiAddBigInt(multiInt("-1", rm2), multiInt("2", rm1), multiInt("-2", r1), r2), bigInt{2, 1})
	c4 := divBigInt(multiAddBigInt(rm2, multiInt("-4", rm1), multiInt("6", r0), multiInt("-4", r1), r2), bigInt{4, 2})

	return multiAddBigInt(
		multiBigInt(c4, powBigInt(bigInt{0, 1}, 4)),
		multiBigInt(c3, powBigInt(bigInt{0, 1}, 3)),
		multiBigInt(c2, powBigInt(bigInt{0, 1}, 2)),
		multiBigInt(c1, bigInt{0, 1}),
		c0)
}

func createEquations(d bigInt, n int) func(x string) bigInt {
	a := d[2*n:]
	b := d[n:2*n]
	c := d[:n]

	return func(x string) bigInt {
			input, _ := stringTobigInt(x)
			return multiAddBigInt(multiBigInt(a, powBigInt(input, 2)), multiBigInt(b, input), c)
		}
}

func multiInt(a string, x bigInt) bigInt {
	mul, _ := stringTobigInt(a)
	return multiBigInt(mul, x)
}

func multiAddBigInt(inputs ...bigInt) bigInt {
	result := bigInt{0}
	for _, input := range inputs {
		result = addBigInt(result, input)
	}
	return result
}

func powBigInt(base bigInt, n int) bigInt {
	result := bigInt{1}
	for i := 1; i < n; i++ {
		result = multiBigInt(result, base)
	}
	return result
}
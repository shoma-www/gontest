package math

import (
	"fmt"
	"testing"
)

func TestAddBigInt(t *testing.T) {
	testCases := []struct {
		x string
		y string
		ans string
	}{
		{
			x: "9234567890",
			y: "987654321",
			ans: "10222222211",
		},
		{
			x: "3",
			y: "4",
			ans: "7",
		},
		{
			x: "2015",
			y: "1121",
			ans: "3136",
		},
	}

	for _, c := range testCases {
		ans, err := AddBigInt(c.x, c.y)
		if err != nil {
			t.Errorf("Invalid Values %s", err)
		}
		fmt.Println(ans)
		if c.ans != ans {
			t.Errorf("unexpected value exp:%s, act:%s", c.ans, ans)
		}
	}
}

func TestSubBigInt(t *testing.T) {
	testCases := []struct {
		x string
		y string
		ans string
	}{
		{
			x: "9234567890",
			y: "987654321",
			ans: "8246913569",
		},
		{
			x: "3",
			y: "4",
			ans: "-1",
		},
	}

	for _, c := range testCases {
		ans, err := SubBigInt(c.x, c.y)
		if err != nil {
			t.Errorf("Invalid Values %s", err)
		}
		fmt.Println(ans)
		if c.ans != ans {
			t.Errorf("unexpected value exp:%s, act:%s", c.ans, ans)
		}
	}
}

func TestMultiBigInt(t *testing.T) {
	testCases := []struct {
		x string
		y string
		ans string
	}{
		{
			x: "0",
			y: "0",
			ans: "0",
		},
		{
			x: "100",
			y: "100",
			ans: "10000",
		},
		{
			x: "23",
			y: "4",
			ans: "92",
		},
		{
			x: "20151121",
			y: "12345678",
			ans: "248779251205038",
		},
	}

	for _, c := range testCases {
		ans, err := MultiBigInt(c.x, c.y)
		if err != nil {
			t.Errorf("Invalid Values %s", err)
		}
		fmt.Println(ans)
		if c.ans != ans {
			t.Errorf("unexpected value exp:%s, act:%s", c.ans, ans)
		}
	}
}

func TestKaratsubaMethod(t *testing.T) {
	testCases := []struct {
		x string
		y string
		ans string
	}{
		{
			x: "100",
			y: "100",
			ans: "10000",
		},
		{
			x: "23",
			y: "4",
			ans: "92",
		},
		{
			x: "20151121",
			y: "12345678",
			ans: "248779251205038",
		},
	}

	for _, c := range testCases {
		ans, err := KaratsubaMethod(c.x, c.y)
		if err != nil {
			t.Errorf("Invalid Values %s", err)
		}
		fmt.Println(ans)
		if c.ans != ans {
			t.Errorf("unexpected value exp:%s, act:%s", c.ans, ans)
		}
	}
}
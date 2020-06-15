package main

import (
	"fmt"
	"math"
	"strconv"
)

//步骤简介
//Karatsuba算法主要应用于两个大数的相乘，原理是将大数分成两段后变成较小的数位，然后做3次乘法，并附带少量的加法操作和移位操作。
//现有两个大数，x，y。
//首先将x，y分别拆开成为两部分，可得x1，x0，y1，y0。他们的关系如下：
//x = x1 * 10m + x0；
//y = y1 * 10m + y0。其中m为正整数，m < n，且x0，y0 小于 10m。
//那么 xy = (x1 * 10m + x0)(y1 * 10m + y0)
//=z2 * 102m + z1 * 10m + z0，其中：
//z2 = x1 * y1；
//z1 = x1 * y0 + x0 * y1；
//z0 = x0 * y0。
//此步骤共需4次乘法，但是由Karatsuba改进以后仅需要3次乘法。因为：
//z1 = x1 * y0+ x0 * y1
//z1 = (x1 + x0) * (y1 + y0) - x1 * y1 - x0 * y0，
//故z1 便可以由一次乘法及加减法得到。

// x = a,b   y=c,d
// z2 = a * c
// z1 = (a+b) * (c+d) - ac - bd
// z0 = b * d
// x * y = z2 * 10^2m + z1 + z0

//实例展示
//设x = 12345，y=6789，令m=3。那么有：
//12345 = 12 * 1000 + 345；
//6789 = 6 * 1000 + 789。
//下面计算：
//z2 = 12 * 6 = 72；
//z0 = 345 * 789 = 272205；
//z1 = (12 + 345) * (6 + 789) - z2 - z0 = 11538。
//然后我们按照移位公式（xy = z2 * 10^(2m) + z1 * 10^(m) + z0）可得：
//xy = 72 * 10002 + 11538 * 1000 + 272205 = 83810205。

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	p := 4
	if len(num1) < p && len(num2) < p {
		n1, _ := strconv.ParseInt(num1, 10, 64)
		n2, _ := strconv.ParseInt(num2, 10, 64)
		return strconv.FormatInt(n1*n2, 10)
	}
	nby2 := max(len(num1), len(num2)) / 2
	m1len := len(num1) - nby2
	a, b := "0", "0"
	if m1len > 0 {
		a = num1[:m1len]
		b = num1[m1len:]
	} else {
		b = num1
	}

	c, d := "0", "0"
	m2len := len(num2) - nby2
	if m2len > 0 {
		c = num2[:m2len]
		d = num2[m2len:]
	} else {
		d = num2
	}
	ac := multiply(a, c)
	bd := multiply(b, d)

	mid := addition(ac, bd)

	ad_plus_bc := subtract(multiply(addition(a, b), addition(c, d)), mid)

	if ac != "0" {
		for i := 0; i < 2*nby2; i++ {
			ac = fmt.Sprintf("%s0", ac)
		}
	}

	if ad_plus_bc != "0" {
		for i := 0; i < nby2; i++ {
			ad_plus_bc = fmt.Sprintf("%s0", ad_plus_bc)
		}
	}
	mid = addition(ac, ad_plus_bc)
	result := addition(mid, bd)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addition(a, b string) string {
	negative := ""
	if a[0] == '-' && b[0] == '-' {
		negative = "-"
	} else if a[0] == '-' {
		return subtract(b, a[1:])
	} else if b[0] == '-' {
		return subtract(a, b[1:])
	}
	return negative + add(a, b)
}
func add(a, b string) string {

	p := 10
	if len(a) < len(b) {
		return add(b, a)
	}
	if len(b) > p {
		_ma := len(a) - p
		_mb := len(b) - p
		low := add(a[_ma:], b[_mb:])
		high := add(a[:_ma], b[:_mb])
		return addsub(high, low, len(low)-p)
	} else {
		if len(a) > p {
			_m := len(a) - len(b)
			low := add(a[_m:], b)
			return addsub(a[:_m], low, len(low)-len(b))
		} else {
			x, _ := strconv.ParseInt(a, 10, 64)
			y, _ := strconv.ParseInt(b, 10, 64)
			return strconv.FormatInt(x+y, 10)
		}
	}

}

func addsub(a, b string, hascarry int) string {
	if hascarry > 0 {
		return add(a, b[0:1]) + b[1:]
	} else if hascarry < 0 {
		for i := 0; i < -hascarry; i++ {
			a = a + "0"
		}
		return a + b
	} else {
		return a + b
	}
}

func subtract(a, b string) string {
	if a[0] == '-' && b[0] == '-' {
		return subtract(b[1:], a[1:])
	} else if a[0] == '-' {
		return addition(a, "-"+b)
	} else if b[0] == '-' {
		return addition(a, b[1:])
	}
	x, y := a, b
	negative := ""
	if len(x) < len(y) || len(x) == len(y) && a < b {
		negative = "-"
		x, y = b, a
	}
	res := sub(x, y)
	index := 0
	// 检测是否全是 0
	for _, w := range res {
		if w != '0' {
			break
		}
		index += 1
	}
	if index == len(res) {
		return "0"
	}
	return negative + res[index:]
}

func sub(a, b string) string {
	p := 10
	if a == b {
		return "0"
	} else if b == "0" {
		return a
	}
	if len(b) > p {
		_ma := len(a) - p
		_mb := len(b) - p
		low := sub(a[_ma:], b[_mb:])
		high := sub(a[:_ma], b[:_mb])
		return subsub(high, low, len(a[:_ma]), p)
	} else {
		if len(a) > p {
			_ma := len(a) - len(b)
			low := sub(a[_ma:], b)
			return subsub(a[:_ma], low, -1, len(b))
		} else {
			x, _ := strconv.ParseInt(a, 10, 64)
			y, _ := strconv.ParseInt(b, 10, 64)
			res := x - y
			if res < 0 {
				return fmt.Sprintf("-%d", int64(math.Pow10(len(a)))+res)
			} else {
				return strconv.FormatInt(res, 10)
			}
		}
	}
}
func subsub(high, low string, size_h, size_l int) string {
	if low[0] == '-' {
		high = sub(high, "1")
		low = low[1:]
	}
	if len(low) < size_l {
		max := size_l - len(low)
		for i := 0; i < max; i++ {
			low = "0" + low
		}
	}
	if size_h > 0 && size_h > len(high) {
		max := size_h - len(high)
		for i := 0; i < max; i++ {
			high = "0" + high
		}
	}
	return high + low
}

func multiply_simple(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	x, y := len(num1), len(num2)
	length := x + y
	ans := make([]int, length)
	for i := x - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		for j := y - 1; j >= 0; j-- {
			sum := ans[i+j+1] + n1*int(num2[j]-'0')
			ans[i+j+1] = sum % 10
			ans[i+j] += sum / 10
		}
	}

	res := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		if i == 0 && ans[i] == 0 {
			continue
		}
		res = append(res, byte(ans[i])+'0')
	}
	return string(res)
}
func main() {
	//num1 := "5"
	//println(num1[3:])
	println(addition("9865678", "22222222222"))
	println(subtract("76553345", "22222222222"))
	//"93553535314"
	//"25247452591474"
	//println(multiply("93553535314","25247452591474"))
	//println(multiply("123","123456789"))
	println(multiply("52795315", "269"))
	//"14201939735"
	//"1421939735

	//"2361907897605003674312836
	//"2361988447605003674312836"
	//"2361988447605003674312836
	//"2361988447605003674312836
	//ad_plus_bc := subtract(multiply(addition("0", "1"), addition("2", "0")), "0")
	//println(ad_plus_bc)
	a := multiply(addition("5279", "5215"), addition("0", "269"))
	println(a)
	println(subtract("2849786", "1429735"))

}

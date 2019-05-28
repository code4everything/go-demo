package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// 计算大数字
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1965)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("big int: %s\n", ip)

	// 计算小数
	rm := big.NewRat(math.MaxInt64, 1965)
	rn := big.NewRat(999, 233)
	fmt.Println("big rat: ", rm.Sub(rm, rn))
}

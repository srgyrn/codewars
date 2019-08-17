package easydiagonal

import "math/big"

func Diagonal(n, p int) int {
	var binomial big.Int
	return int(binomial.Binomial(int64(n+1), int64(p+1)).Int64())
}

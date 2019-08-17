package easydiagonal

import "math/big"

func Diagonal(n, p int) int {
	var result, binomial1, binomial2 big.Int

	binomial1.Binomial(int64(n), int64(p))
	binomial2.Binomial(int64(n), int64(p+1))

	result = *result.Add(&binomial1, &binomial2)

	return int(result.Int64())
}

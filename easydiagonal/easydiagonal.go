package easydiagonal

import "math/big"

func Diagonal(n, p int) int {
	var result, binomial big.Int

	for i := 0; i <= n; i++ {
		binomial.Binomial(int64(i), int64(p))
		result = *result.Add(&result, &binomial)
	}

	return int(result.Int64())
}

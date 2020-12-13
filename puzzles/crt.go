package puzzles

import "math/big"

//Chinese Reminder Theorem implementation

type CRTEquation struct {
	b, n int
}

type CRTEquations []CRTEquation

func NewEquation(b, n int) CRTEquation {
	return CRTEquation{b, n}
}

func (eqs CRTEquations) Solve() int64 {
	var product int64 = 1
	for _, eq := range eqs {
		product *= int64(eq.n)
	}

	N := make([]int64, len(eqs))
	x := make([]*big.Int, len(eqs))
	for i, eq := range eqs {
		N[i] = product / int64(eq.n)
		b := big.NewInt(1)
		x[i] = b.ModInverse(big.NewInt(N[i]), big.NewInt(int64(eq.n)))
	}

	var X int64
	for i, eq := range eqs {
		X += int64(eq.b) * N[i] * x[i].Int64()
	}

	return X%product

}

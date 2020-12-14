package puzzles

import (
	"io"
	"math"
	"math/big"
	"sort"
)

type DockingData struct {
}

type bitOperations []bitOperation

type bitOperation struct {
	maskOp *maskOperation
	memOp  *memOperation
}

type maskOperation struct {
	val map[int]int
}

type memOperation struct {
	loc int64
	val int64
}

func (d DockingData) Puzzle1(reader io.Reader) (Result, error) {
	ops, err := ParseInitializationProgram(reader)
	if err != nil {
		return nil, err
	}

	var maskBits map[int]int
	var data = make(map[int64]int64)
	for _, op := range ops {
		if op.maskOp != nil {
			maskBits = op.maskOp.val
		}

		if op.memOp != nil {
			b := big.NewInt(op.memOp.val)
			for index, mbit := range maskBits {
				if mbit == -1 {
					continue
				}
				b.SetBit(b, index, uint(mbit))
			}
			data[op.memOp.loc] = b.Int64()
		}
	}

	var sum int64
	for _, v := range data {
		sum += v
	}
	return int64Result(sum), nil
}

func (d DockingData) Puzzle2(reader io.Reader) (Result, error) {
	ops, err := ParseInitializationProgram(reader)
	if err != nil {
		return nil, err
	}

	var maskBits map[int]int
	var data = make(map[int64]int64)
	for _, op := range ops {
		if op.maskOp != nil {
			maskBits = op.maskOp.val
		}

		if op.memOp != nil {
			address := op.memOp.loc
			newAddresses := decodeAddress(address, maskBits)
			for _, n := range newAddresses {
				data[n] = op.memOp.val
			}
		}
	}
	var sum int64
	for _, v := range data {
		sum += v
	}
	return int64Result(sum), nil
}

func decodeAddress(address int64, maskBits map[int]int) []int64 {
	b := big.NewInt(address)

	var floatingBits []int
	for index, mbit := range maskBits {
		if mbit == 1 {
			b.SetBit(b, index, uint(mbit))
		}
		if mbit == -1 {
			floatingBits = append(floatingBits, index)
		}
	}

	sort.Ints(floatingBits)

	var newAddresses []int64
	limit := math.Pow(float64(2), float64(len(floatingBits)))

	for i := 0; i < int(limit); i++ {
		x := big.NewInt(b.Int64())
		toReplace := big.NewInt(int64(i))
		for i, f := range floatingBits {
			x.SetBit(x, f, toReplace.Bit(i))
		}

		newAddresses = append(newAddresses, x.Int64())
	}
	return newAddresses
}

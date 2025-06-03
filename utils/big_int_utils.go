package utils

import (
	"errors"
	"math/big"
)

func BigIntToString(bigInt *big.Int) string {
	return bigInt.String()
}

func StringToBigInt(str string) (*big.Int, bool) {
	return new(big.Int).SetString(str, 10)
}

func CompareBigIntStrings(string1 *string, string2 *string) (int, error) {
	if string1 == nil && string2 == nil {
		return 0, nil
	} else if string1 == nil {
		return -1, nil
	} else if string2 == nil {
		return 1, nil
	} else {
		bigInt1, ok1 := new(big.Int).SetString(*string1, 10)
		bigInt2, ok2 := new(big.Int).SetString(*string2, 10)
		if !ok1 || !ok2 {
			return 0, errors.New("invalid big int string")
		}
		return bigInt1.Cmp(bigInt2), nil
	}
}

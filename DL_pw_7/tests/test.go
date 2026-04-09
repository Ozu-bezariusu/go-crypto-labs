package tests

import (
	ec "DL_pw_7/EC"
	"math/big"
	"reflect"
)

func ECTest() bool {
	var g ec.ECPoint = ec.BasePointGGet()
	k := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(130), nil)
	d := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(130), nil)

	h1 := ec.ScalarMult(*d, g)
	h2 := ec.ScalarMult(*k, h1)

	h3 := ec.ScalarMult(*k, g)
	h4 := ec.ScalarMult(*d, h3)

	return reflect.DeepEqual(h2, h4)
}

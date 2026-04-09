package ec

import (
	"encoding/json"
	"fmt"
	"math/big"

	secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

func getCurve() *secp256k1.KoblitzCurve {
	return secp256k1.S256()
}

func ECPointGen(x, y *big.Int) ECPoint {
	return ECPoint{x, y}
}

func BasePointGGet() ECPoint {
	curve := getCurve()

	return ECPointGen(curve.Gx, curve.Gy)
}

func IsOnCurveCheck(point ECPoint) bool {
	return getCurve().IsOnCurve(point.X, point.Y)
}

func AddECPoints(a, b ECPoint) ECPoint {
	resX, resY := getCurve().Add(a.X, a.Y, b.X, b.Y)

	return ECPointGen(resX, resY)
}

func DoubleECPoints(a ECPoint) (c ECPoint) {
	resX, resY := getCurve().Double(a.X, a.Y)

	return ECPointGen(resX, resY)
}

func ScalarMult(k big.Int, a ECPoint) (c ECPoint) {
	resX, resY := getCurve().ScalarMult(a.X, a.Y, k.Bytes())

	return ECPointGen(resX, resY)
}

func ECPointToString(point ECPoint) (s string) {
	jsonData, err := json.Marshal(point)
	if err != nil {
		fmt.Printf("JSON Marshalling failed: %s", err)
	}

	return string(jsonData)
}

func StringToECPoint(s string) (point ECPoint) {
	var p ECPoint
	err := json.Unmarshal([]byte(s), &p)
	if err != nil {
		fmt.Printf("JSON Unmarshalling failed: %s", err)
	}

	return p
}

func PrintECPoint(point ECPoint) {
	fmt.Printf("X coordinate of the point: %d, Y coordinate: %d\n", point.X, point.Y)
}

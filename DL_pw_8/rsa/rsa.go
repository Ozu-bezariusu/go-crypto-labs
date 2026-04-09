package rsa

import (
	"crypto/rand"
	"math/big"
)

type PublicKey struct {
	E *big.Int
	N *big.Int
}

type PrivateKey struct {
	D *big.Int
	N *big.Int
}

func KeyGen() (public PublicKey, private PrivateKey) {
	bitSize := 1024

	p, _ := rand.Prime(rand.Reader, bitSize)
	q, _ := rand.Prime(rand.Reader, bitSize)

	n := new(big.Int).Mul(p, q)

	pMinusOne := new(big.Int).Sub(p, big.NewInt(1))
	qMinusOne := new(big.Int).Sub(q, big.NewInt(1))
	m := new(big.Int).Mul(pMinusOne, qMinusOne)

	e := big.NewInt(3)
	for {
		gcd := new(big.Int).GCD(nil, nil, e, m)
		if gcd.Cmp(big.NewInt(1)) == 0 {
			break
		}

		e.Add(e, big.NewInt(2))
	}

	d := new(big.Int).ModInverse(e, m)

	return PublicKey{e, n}, PrivateKey{d, n}
}

func Encrypt(data *big.Int, key *PublicKey) (result *big.Int) {
	return new(big.Int).Exp(data, key.E, key.N)
}

func Decrypt(data *big.Int, key *PrivateKey) (result *big.Int) {
	return new(big.Int).Exp(data, key.D, key.N)
}

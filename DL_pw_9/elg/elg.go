package elg

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// basic check, not full
func findPrimitiveRoot(p *big.Int) (*big.Int, error) {
	two := big.NewInt(2)
	pSub1 := new(big.Int).Sub(p, big.NewInt(1))

	for {
		g, err := generateRandomInt(pSub1)
		if err != nil {
			return nil, err
		}

		if new(big.Int).Exp(g, two, p).Cmp(big.NewInt(1)) != 0 && new(big.Int).Exp(g, pSub1.Div(pSub1, two), p).Cmp(big.NewInt(1)) != 0 {
			return g, nil
		}
	}
}

func generateRandomPrime(bits int) (*big.Int, error) {
	return rand.Prime(rand.Reader, bits)
}

func generateRandomInt(max *big.Int) (*big.Int, error) {
	n, err := rand.Int(rand.Reader, new(big.Int).Sub(max, big.NewInt(1)))

	if err != nil {
		return nil, err
	}

	return n.Add(n, big.NewInt(1)), nil
}

func hashToBigInt(data []byte) *big.Int {
	hash := sha256.Sum256(data)
	return new(big.Int).SetBytes(hash[:])
}

func PrepareParams(bits int) (p *big.Int, g *big.Int) {
	p, err := generateRandomPrime(bits)
	if err != nil {
		fmt.Println("Помилка при генерації простого числа:", err)
		return
	}
	g, err = findPrimitiveRoot(p)
	if err != nil {
		fmt.Println("Помилка при пошуку примітивного кореня:", err)
		return
	}
	return p, g
}

func GenerateKeys(p *big.Int, g *big.Int) (a *big.Int, b *big.Int) {
	a, err := generateRandomInt(p)
	if err != nil {
		fmt.Println("Помилка при генерації особистого ключа:", err)
		return
	}
	b = new(big.Int).Exp(g, a, p)

	return a, b
}

func Sign(p *big.Int, g *big.Int, a *big.Int, b *big.Int, message []byte) (r *big.Int, s *big.Int) {
	var kInv *big.Int
	pSub1 := new(big.Int).Sub(p, big.NewInt(1))
	hm := hashToBigInt(message)

	for {
		k, err := generateRandomInt(pSub1)
		if err != nil {
			fmt.Println("Помилка при генерації k:", err)
			return
		}

		r = new(big.Int).Exp(g, k, p)
		s = new(big.Int).Sub(hm, new(big.Int).Mul(a, r))
		s.Mod(s, pSub1)

		kInv = new(big.Int).ModInverse(k, pSub1)
		if kInv == nil {
			continue
		}

		s.Mul(s, kInv).Mod(s, pSub1)
		if new(big.Int).GCD(nil, nil, s, pSub1).Cmp(big.NewInt(1)) == 0 {
			break
		}
	}

	return r, s
}

func CheckSignature(p *big.Int, g *big.Int, a *big.Int, b *big.Int, r *big.Int, s *big.Int, message []byte) bool {
	pSub1 := new(big.Int).Sub(p, big.NewInt(1))
	hm := hashToBigInt(message)

	y := new(big.Int).ModInverse(b, p)
	sInv := new(big.Int).ModInverse(s, pSub1)
	u1 := new(big.Int).Mul(hm, sInv)
	u1.Mod(u1, pSub1)
	u2 := new(big.Int).Mul(r, sInv)
	u2.Mod(u2, pSub1)
	gu1 := new(big.Int).Exp(g, u1, p)
	yu2 := new(big.Int).Exp(y, u2, p)
	v := new(big.Int).Mul(gu1, yu2)
	v.Mod(v, p)

	return v.Cmp(r) == 0
}

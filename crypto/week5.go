package main

import (
	"fmt"
	"math/big"
)

func computeX0(p, g, B, x0 *big.Int) *big.Int {
	res := new(big.Int).Exp(g, B, p)
	res.Exp(res, x0, p)

	return res
}

func computeX1(p, g, h, x1 *big.Int) *big.Int {
	t1 := new(big.Int).Exp(g, x1, p)
	t2 := t1.ModInverse(t1, p)
	t3 := t2.Mul(t2, h)
	t3.Mod(t3, p)

	return t3
}

func main() {
	B := new(big.Int).Exp(big.NewInt(2), big.NewInt(20), big.NewInt(0))

	p := new(big.Int)
	g := new(big.Int)
	h := new(big.Int)

	_, err := fmt.Sscan("13407807929942597099574024998205846127479365820592393377723561443721764030073546976801874298166903427690031858186486050853753882811946569946433649006084171", p)
	_, err = fmt.Sscan("11717829880366207009516117596335367088558084999998952205599979459063929499736583746670572176471460312928594829675428279466566527115212748467589894601965568", g)
	_, err = fmt.Sscan("3239475104050450443565264378728065788649097520952449527834792452971981976143292558073856937958553180532878928001494706097394108577585732452307673444020333", h)
	if err != nil {
		fmt.Println(err)
	}

	// Create a map that holds the calculations for x1's as strings
	hashTable := make(map[string]int64, B.Int64())

	// Calculate the x1's and store them in the map
	var i int64
	for i = 0; i < B.Int64(); i++ {
		hashTable[computeX1(p, g, h, big.NewInt(i)).String()] = i
	}

	// Calculate the x0's and check if they are part of the hash table
	for i = 0; i < B.Int64(); i++ {
		x0Res := computeX0(p, g, B, big.NewInt(i))

		x1, found := hashTable[x0Res.String()]
		if found {
			x := B.Mul(B, big.NewInt(i))
			x.Add(x, big.NewInt(x1))
			fmt.Println("Solution:", x)
			return
		}
	}
}

package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/cznic/mathutil"
)

func factor(A, N *big.Int) (*big.Int, *big.Int) {

	t2 := big.NewInt(0)
	t2.Exp(A, big.NewInt(2), nil)
	t2.Sub(t2, N)
	if t2.Cmp(big.NewInt(0)) <= 0 {
		return big.NewInt(0), big.NewInt(0)
	}
	x := mathutil.SqrtBig(t2)

	q := big.NewInt(0)
	p := big.NewInt(0)

	q = q.Sub(A, x)
	p = p.Add(A, x)

	return q, p
}

func checkFactors(p, q, N *big.Int) bool {
	temp := big.NewInt(0)
	if temp.Mul(q, p).Cmp(N) == 0 {
		return true
	}
	return false
}

func main() {

	// Challenge #1
	N := new(big.Int)
	fmt.Sscan("179769313486231590772930519078902473361797697894230657273430081157732675805505620686985379449212982959585501387537164015710139858647833778606925583497541085196591615128057575940752635007475935288710823649949940771895617054361149474865046711015101563940680527540071584560878577663743040086340742855278549092581", N)

	A := mathutil.SqrtBig(N)
	A = A.Add(A, big.NewInt(1))

	q, p := factor(A, N)

	if checkFactors(p, q, N) {
		fmt.Println("Challenge #1 Solution")
		fmt.Println("q:", q)
		fmt.Println("p:", p)
	}

	// Challenge #2
	N2 := new(big.Int)
	fmt.Sscan("648455842808071669662824265346772278726343720706976263060439070378797308618081116462714015276061417569195587321840254520655424906719892428844841839353281972988531310511738648965962582821502504990264452100885281673303711142296421027840289307657458645233683357077834689715838646088239640236866252211790085787877", N2)

	A2 := mathutil.SqrtBig(N2)

	for i := A2; i.Cmp(N2) <= 0; i.Add(i, big.NewInt(1)) {
		q2, p2 := factor(A2, N2)
		if checkFactors(p2, q2, N2) {
			fmt.Println("Challenge #2 Solution")
			fmt.Println("q:", q2)
			fmt.Println("p:", p2)
			break
		}
	}

	// Challenge #4
	cipher := new(big.Int)
	fmt.Sscan("22096451867410381776306561134883418017410069787892831071731839143676135600120538004282329650473509424343946219751512256465839967942889460764542040581564748988013734864120452325229320176487916666402997509188729971690526083222067771600019329260870009579993724077458967773697817571267229951148662959627934791540", cipher)
	e := big.NewInt(65537)
	phi := big.NewInt(1)
	phi = phi.Mul(p.Sub(p, big.NewInt(1)), q.Sub(q, big.NewInt(1)))

	d := big.NewInt(0)
	d.ModInverse(e, phi)

	cipher.Exp(cipher, d, N)
	hexEncoded := fmt.Sprintf("%x", cipher)

	ind := strings.Index(hexEncoded, "00")
	if ind != -1 {
		fmt.Println("Challenge #4 Solution")
		res, _ := hex.DecodeString(hexEncoded[ind+2:])
		fmt.Println(string(res))
	}

}

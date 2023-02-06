//go:build !math_big_pure_go
// +build !math_big_pure_go

package multiexp

// implemented in arith_$GOARCH.s
func addVV(z, x, y []Word) (c Word)
func subVV(z, x, y []Word) (c Word)
func addVW(z, x []Word, y Word) (c Word)
func subVW(z, x []Word, y Word) (c Word)
func shlVU(z, x []Word, s uint) (c Word)
func shrVU(z, x []Word, s uint) (c Word)
func mulAddVWW(z, x []Word, y, r Word) (c Word)
func addMulVVW(z, x []Word, y Word) (c Word)

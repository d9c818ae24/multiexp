//go:build !math_big_pure_go
// +build !math_big_pure_go

package multiexp

import "internal/cpu"

func addVV_check(z, x, y []Word) (c Word)
func addVV_vec(z, x, y []Word) (c Word)
func addVV_novec(z, x, y []Word) (c Word)
func subVV_check(z, x, y []Word) (c Word)
func subVV_vec(z, x, y []Word) (c Word)
func subVV_novec(z, x, y []Word) (c Word)

var hasVX = cpu.S390X.HasVX

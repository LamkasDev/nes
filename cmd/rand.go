package main

import (
	"math/rand"
)

func GetRand16() uint8 {
	return uint8(1 + rand.Intn(16))
}

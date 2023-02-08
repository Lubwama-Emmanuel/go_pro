package helpers

import "math/rand"

func RandomNumber(n int) int {
	ranNum := rand.Intn(n)
	return ranNum
}

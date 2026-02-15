package random

import (
	"math/rand"
	"time"
)

// NewRandomString generates random string with given size
func NewRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano())) // testing and other stuff

	chars := []rune("QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890")

	rString := make([]rune, size)

	for i := 0; i < size; i++ {
		rString[i] = chars[rnd.Intn(len(chars))]
	}

	return string(rString)

}

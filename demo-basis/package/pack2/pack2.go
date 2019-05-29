package pack2

import (
	"math/rand"
	"time"
)

var letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func GetRandomString(lens int) string {
	rand.Seed(time.Now().Unix())
	bs := make([]byte, lens)
	for i := 0; i < lens; i++ {
		bs[i] = letters[rand.Intn(len(letters))]
	}
	return string(bs)
}

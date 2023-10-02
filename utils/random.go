package utils


import (
	"math/rand"
	"strings"
	"time"
)

// This file generates random test data for testing


const alphabet = "ABCDEFGHIJKLMNOPQRSTVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

//initialises the seed of the random packeage to make sure that value are indeed random
func init(){
	rand.Seed(time.Now().UnixNano())
}

// RandomString generates a random string of length n
func RandomString() string {
	n:= 24
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
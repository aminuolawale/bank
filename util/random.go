package util

import (
	"math/rand"
	"strings"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

 const alphabets = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64)int64{
	return min + rand.Int63n(max - min +1)
}

func RandomString(n int) string {
	k := len(alphabets)
	var sb strings.Builder
	for i:=0; i < n; i++ {
		c:= alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return  RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(12, 2000)
}

func RandomCurrency() string {
	return strings.ToUpper(RandomString(3))
}
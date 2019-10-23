package goapp

import (
	"fmt"
	"math/rand"
	"time"
)

type PassType int

const (
	NUmStr                 = "0123456789"
	CharStr                = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr                = "+=-@#~,.[]()!%^*$"
	Number        PassType = 0
	Char          PassType = 1
	NumberAndChar PassType = 2
	Advance       PassType = 3
)

func BuildPassword(length int, passType PassType) string {
	rand.Seed(time.Now().UnixNano())
	var pass []byte = make([]byte, length, length)
	var sourceStr string
	switch passType {
	case Number:
		sourceStr = NUmStr
		break
	case Char:
		sourceStr = CharStr
		break
	case NumberAndChar:
		sourceStr = fmt.Sprintf("%s%s", NUmStr, CharStr)
		break
	case Advance:
		sourceStr = fmt.Sprintf("%s%s%s", NUmStr, CharStr, SpecStr)
		break
	default:
		sourceStr = NUmStr
	}
	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		pass[i] = sourceStr[index]
	}
	return string(pass)
}

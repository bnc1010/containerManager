package utils

import (
	"math/rand"
	"time"
	"unsafe"
)

var src = rand.NewSource(time.Now().UnixNano())

const (
    // 6 bits to represent a letter index
    letterIdBits = 6
    // All 1-bits as many as letterIdBits
    letterIdMask = 1<<letterIdBits - 1
    letterIdMax  = 63 / letterIdBits
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randStr(n int64) string {
    b := make([]byte, n)
    // A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
    for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdMax
        }
        if idx := int(cache & letterIdMask); idx < len(letters) {
            b[i] = letters[idx]
            i--
        }
        cache >>= letterIdBits
        remain--
    }
    return *(*string)(unsafe.Pointer(&b))
}

func RandStringWithLengthN(length int64) string {
	return randStr(length)
}
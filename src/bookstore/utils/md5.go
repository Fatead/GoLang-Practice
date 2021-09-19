package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(unEnc string) (encryption string) {
	byte16 := []byte(unEnc)
	encryption = fmt.Sprintf("%x", md5.Sum(byte16))
	return
}

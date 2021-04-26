package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

//随机生成文件名
func RandomString(n int) string {
	var letter = []byte("qwertyuiopassdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXVBNM")
	res := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range res {
		res[i] = letter[rand.Intn(len(letter))]
	}
	return string(res)
}

//文件名加密
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

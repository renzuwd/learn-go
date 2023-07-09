package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func hs256(secret, data []byte) (ret string, err error) {
	hasher := hmac.New(sha256.New, secret)
	_, err = hasher.Write(data)
	if err != nil {
		return "", err
	}
	r := hasher.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(r), nil
}

func main() {
	sig, _ := hs256([]byte("secret"), []byte("aabbcc"))
	fmt.Println(sig)
}

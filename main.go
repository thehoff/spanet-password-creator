package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	secret := "1ld0gVand"
	data := os.Args[1]
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sEnc := base64.StdEncoding.EncodeToString(h.Sum(nil))
	fmt.Println(strings.Replace(strings.Replace(strings.Replace(sEnc, "=", "3", -1), "/", "2", -1), "+", "1", -1))
}

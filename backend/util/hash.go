package util

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/name5566/leaf/log"
)

func HashPwd(pwd string) string {
	t := sha1.New()
	_, err := io.WriteString(t, pwd)
	if err != nil {
		log.Fatal("%v", err)
	}
	return fmt.Sprintf("%x", t.Sum(nil))
}

func GetToken() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("%v", err)
	}
	return fmt.Sprintf("%x", b)
}

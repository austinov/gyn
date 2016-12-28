package util

import "golang.org/x/crypto/bcrypt"

func Hash(psw string) string {
	bpsw := []byte(psw)
	hpsw, err := bcrypt.GenerateFromPassword(bpsw, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hpsw)
}

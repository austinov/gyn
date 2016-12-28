package util

import "golang.org/x/crypto/bcrypt"

func Hash(text string) string {
	btext := []byte(text)
	htext, err := bcrypt.GenerateFromPassword(btext, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(htext)
}

func CompareHashAndText(hash, text string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
}

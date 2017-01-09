package util

import "github.com/fiam/gounidecode/unidecode"

func Translit(text string) string {
	return unidecode.Unidecode(text)
}

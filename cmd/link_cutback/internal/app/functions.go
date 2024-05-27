package app

import (
	"math/rand"
	"strings"
	"time"
)

func shortLinkCreator() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" +
		"_")
	length := 10
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[r.Intn(len(chars))])
	}
	//fmt.Println(b.String())
	return b.String()
}

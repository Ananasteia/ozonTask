package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var Links map[Link]string

type Link string

func GetOriginalLink(w http.ResponseWriter, r *http.Request) {
	jsong, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.Context()
	var newLink Link
	err = json.Unmarshal(jsong, &newLink)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func Shorter() {

}
func ShortLink(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", GetOriginalLink).Methods(http.MethodPost)
	http.HandleFunc("/", ShortLink).Methods(http.MethodGet)
}

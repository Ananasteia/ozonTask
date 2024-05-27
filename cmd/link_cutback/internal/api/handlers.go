package api

import (
	"AVITOtask/cmd/link_cutback/internal/app"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func (a *api) giveShortLink(w http.ResponseWriter, r *http.Request) {

	var userLink app.Link
	jsong, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(jsong, &userLink)
	if err != nil {
		fmt.Println(err)
		return
	}
	// err = json.NewDecoder(r.Body).Decode(&userLink) // можно сразу декодировать и указать куда
	fullLink, err := a.app.HandlePost(r.Context(), userLink.LongLink)
	//err = json.NewEncoder(w).Encode(fullLink)
	data, err := json.Marshal(fullLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (a *api) openLongLink(w http.ResponseWriter, r *http.Request) {
	//
	//var userLink app.Link
	//jsong, err := io.ReadAll(r.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//err = json.Unmarshal(jsong, &userLink)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	vars := mux.Vars(r)
	userLink := vars["link"]

	fullLink, err := a.app.HandleGet(r.Context(), app.Link{ShortLink: userLink})
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, fullLink.LongLink, http.StatusMovedPermanently)
}

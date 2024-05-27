package main

import (
	"AVITOtask/cmd/link_cutback/internal/adapters/repo"
	"AVITOtask/cmd/link_cutback/internal/api"
	"AVITOtask/cmd/link_cutback/internal/app"
	"log"
	"net/http"
)

func main() {
	db, err := repo.New()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	linkApp := app.New(db)
	linkApi := api.New(linkApp)
	server := http.Server{Addr: ":8080", Handler: linkApi}
	log.Fatal(server.ListenAndServe())
}

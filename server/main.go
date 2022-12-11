package main

import (
	"log"
	"net/http"

	"github.com/hararudoka/clamo/server/nethttp"
	"github.com/hararudoka/clamo/server/service"
	"github.com/hararudoka/clamo/server/storage"
)

func main() {
	db, err := storage.Open()
	if err != nil {
		log.Fatal(err)
	}

	s := service.New(db)
	h := nethttp.New(s)

	// run server
	if err := http.ListenAndServe(":80", h); err != nil {
		log.Fatal(err)
	}
}

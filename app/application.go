package app

import (
	"github.com/ArturoAguilar1/b_items-api/clients/elasticsearch"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:8082",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
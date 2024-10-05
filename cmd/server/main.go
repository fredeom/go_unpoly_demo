/*
https://practicalgobook.net/posts/go-sqlite-no-cgo/
*/
package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi/v5"

	"github.com/fredeom/go_unpoly_demo/internal/db"
	"github.com/fredeom/go_unpoly_demo/internal/handlers"
	"github.com/fredeom/go_unpoly_demo/internal/service"
)

const dbName = "db3"

func main() {
	r := chi.NewRouter()

	store, err := db.NewStore(dbName)
	if err != nil {
		log.Fatalf("failed to create store: %s", err)
	}

	cs := service.New(store)

	hCompany := handlers.NewCompanyHandler(cs)

	handlers.SetupRoutes(r, hCompany)

	err3 := http.ListenAndServe(":3000", r)
	if err3 != nil {
		log.Fatalln(err3.Error())
	}
}

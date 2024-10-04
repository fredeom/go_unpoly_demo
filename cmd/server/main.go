/*
https://practicalgobook.net/posts/go-sqlite-no-cgo/
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/fredeom/go_unpoly_demo/internal/db"
	"github.com/fredeom/go_unpoly_demo/internal/views"
)

func main() {
	// var DB, err = db.InitDatabase("db3")
	var DB, err = db.PopulateDemoDB("db3")
	if err != nil {
		log.Fatal(err.Error())
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		views.Index().Render(r.Context(), w)
	})
	r.Get("/companies", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		companies, err := db.QueryCompanies(DB, query)
		if err != nil {
			views.Error(err.Error()).Render(r.Context(), w)
		} else {
			views.Companies(companies).Render(r.Context(), w)
		}
	})
	r.Get("/companies/{id}", func(w http.ResponseWriter, r *http.Request) {
		value, _ := strconv.Atoi(chi.URLParam(r, "id"))
		company, err := db.QueryCompany(DB, int64(value))
		if err != nil {
			views.Error(err.Error()).Render(r.Context(), w)
			return
		}
		views.Company(company).Render(r.Context(), w)
	})
	r.Get("/companies/{id}/edit", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		companyId := r.Form.Get("company[ID]")
		companyName := r.Form.Get("company[name]")
		companyAddress := r.Form.Get("company[address]")

		if companyId != "" {
			id, _ := strconv.Atoi(companyId)
			db.EditCompany(DB, int64(id), companyName, companyAddress)
			http.Redirect(w, r, fmt.Sprintf("/companies/%v", id), http.StatusTemporaryRedirect)
			return
		}

		value, _ := strconv.Atoi(chi.URLParam(r, "id"))
		company, err := db.QueryCompany(DB, int64(value))
		if err != nil {
			views.Error(err.Error()).Render(r.Context(), w)
			return
		}
		views.EditCompany(company).Render(r.Context(), w)
	})
	r.Post("/companies/{id}", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		method := r.Form.Get("_method")
		if method == "DELETE" {
			value, _ := strconv.Atoi(chi.URLParam(r, "id"))
			err := db.DeleteCompany(DB, int64(value))
			if err != nil {
				views.Error(err.Error()).Render(r.Context(), w)
				return
			}
			w.Header().Set("X-Up-Events", "[{ \"type\": \"company:destroyed\"}]")
		}
	})
	r.Get("/companies/new", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		companyName := r.Form.Get("company[name]")
		companyAddress := r.Form.Get("company[address]")

		if companyName == "" || companyAddress == "" {
			views.CompanyNew().Render(r.Context(), w)
			return
		}

		companyId, err1 := db.InsertCompany(DB, companyName, companyAddress)
		if err1 != nil {
			views.Error(err1.Error()).Render(r.Context(), w)
			return
		}
		http.Redirect(w, r, "/companies/"+fmt.Sprintf("%v", companyId), http.StatusTemporaryRedirect)
	})
	err3 := http.ListenAndServe(":3000", r)
	if err3 != nil {
		log.Fatalln(err3.Error())
	}
}

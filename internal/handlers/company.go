package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fredeom/go_unpoly_demo/internal/domain"
	"github.com/fredeom/go_unpoly_demo/internal/views"
	"github.com/go-chi/chi/v5"
)

type CompanyService interface {
	QueryCompanies(query string) ([]domain.Company, error)
	QueryCompany(id int64) (domain.Company, error)
	EditCompany(id int64, name string, address string) sql.Result
	DeleteCompany(id int64) error
	NewCompany(name string, address string) (int64, error)
	PopulateStore() error
}

type CompanyHandler struct {
	CompanyService CompanyService
}

func NewCompanyHandler(cs CompanyService) *CompanyHandler {
	return &CompanyHandler{
		CompanyService: cs,
	}
}

func (ch *CompanyHandler) HandleQueryCompanies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	companies, err := ch.CompanyService.QueryCompanies(query)
	if err != nil {
		views.Error(err.Error()).Render(r.Context(), w)
	} else {
		views.Companies(companies).Render(r.Context(), w)
	}
}

func (ch *CompanyHandler) HandleShowCompany(w http.ResponseWriter, r *http.Request) {
	value, _ := strconv.Atoi(chi.URLParam(r, "id"))
	company, err := ch.CompanyService.QueryCompany(int64(value))
	if err != nil {
		views.Error(err.Error()).Render(r.Context(), w)
		return
	}
	views.Company(company).Render(r.Context(), w)
}

func (ch *CompanyHandler) HandleEditCompany(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	companyId := r.Form.Get("company[ID]")
	companyName := r.Form.Get("company[name]")
	companyAddress := r.Form.Get("company[address]")

	if companyId != "" {
		id, _ := strconv.Atoi(companyId)
		ch.CompanyService.EditCompany(int64(id), companyName, companyAddress)
		http.Redirect(w, r, fmt.Sprintf("/companies/%v", id), http.StatusTemporaryRedirect)
		return
	}

	value, _ := strconv.Atoi(chi.URLParam(r, "id"))
	company, err := ch.CompanyService.QueryCompany(int64(value))
	if err != nil {
		views.Error(err.Error()).Render(r.Context(), w)
		return
	}
	views.EditCompany(company).Render(r.Context(), w)
}

func (ch *CompanyHandler) HandleDeleteCompany(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	method := r.Form.Get("_method")
	if method == "DELETE" {
		value, _ := strconv.Atoi(chi.URLParam(r, "id"))
		err := ch.CompanyService.DeleteCompany(int64(value))
		if err != nil {
			views.Error(err.Error()).Render(r.Context(), w)
			return
		}
		w.Header().Set("X-Up-Events", "[{ \"type\": \"company:destroyed\"}]")
	}

}

func (ch *CompanyHandler) HandleNewCompany(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	companyName := r.Form.Get("company[name]")
	companyAddress := r.Form.Get("company[address]")

	if companyName == "" || companyAddress == "" {
		views.CompanyNew().Render(r.Context(), w)
		return
	}

	companyId, err1 := ch.CompanyService.NewCompany(companyName, companyAddress)
	if err1 != nil {
		views.Error(err1.Error()).Render(r.Context(), w)
		return
	}
	http.Redirect(w, r, "/companies/"+fmt.Sprintf("%v", companyId), http.StatusTemporaryRedirect)
}

func (ch *CompanyHandler) HandlePopulateStore(w http.ResponseWriter, r *http.Request) {
	err := ch.CompanyService.PopulateStore()
	if err != nil {
		w.Write([]byte("<article>Проблема в заполнении таблиц демо данными</article>"))
	}
	w.Write([]byte("<article>Таблицы обновлены демо данными</article>"))
}

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

type ProjectService interface {
	QueryProjects(query string) ([]domain.Project, error)
	NewProject(companyId int64, name string, budget int64) (int64, error)
	DeleteProject(id int64) error
	QueryProject(id int64) (domain.Project, error)
	EditProject(id int64, companyId int64, projectName string, budget int64) sql.Result

	QueryCompanyByCompanyId(companyId int64) (domain.Company, error)
	QueryCompanyNamesByCompanyIDs() (map[int64]string, error)
	QueryCompanies(query string) ([]domain.Company, error)
}

type ProjectHandler struct {
	ProjectService ProjectService
}

func NewProjectHandler(ps ProjectService) *ProjectHandler {
	return &ProjectHandler{
		ProjectService: ps,
	}
}

func (ph *ProjectHandler) HandleQueryProjects(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	projects, err := ph.ProjectService.QueryProjects(query)
	if err != nil {
		views.Error(err.Error()).Render(r.Context(), w)
	}
	companyNameByCompanyID, err2 := ph.ProjectService.QueryCompanyNamesByCompanyIDs()
	if err2 != nil {
		views.Error(err2.Error()).Render(r.Context(), w)
	}
	views.Projects(projects, companyNameByCompanyID).Render(r.Context(), w)
}

func (ph *ProjectHandler) HandleDeleteProject(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	method := r.Form.Get("_method")
	if method == "DELETE" {
		value, _ := strconv.Atoi(chi.URLParam(r, "id"))
		err := ph.ProjectService.DeleteProject(int64(value))
		if err != nil {
			views.Error(err.Error()).Render(r.Context(), w)
			return
		}
		w.Header().Set("X-Up-Events", "[{ \"type\": \"project:destroyed\"}]")
	}
}

func (ph *ProjectHandler) HandleNewProject(w http.ResponseWriter, r *http.Request) {
	companyId := r.URL.Query().Get("company_id")

	r.ParseForm()

	parentCompanyId := r.Form.Get("company[id]")
	projectName := r.Form.Get("project[name]")
	projectBudget := r.Form.Get("project[budget]")

	if projectName == "" || projectBudget == "" || parentCompanyId == "0" {
		compID, _ := strconv.Atoi(companyId)
		companies, _ := ph.ProjectService.QueryCompanies("")
		views.ProjectNew(int64(compID), companies).Render(r.Context(), w)
		return
	}

	cID, _ := strconv.Atoi(parentCompanyId)
	pBudget, _ := strconv.Atoi(projectBudget)

	projectId, err1 := ph.ProjectService.NewProject(int64(cID), projectName, int64(pBudget))
	if err1 != nil {
		views.Error(err1.Error()).Render(r.Context(), w)
		return
	}
	http.Redirect(w, r, "/projects/"+fmt.Sprintf("%v", projectId), http.StatusTemporaryRedirect)
}

func (ph *ProjectHandler) HandleShowProject(w http.ResponseWriter, r *http.Request) {
	value, _ := strconv.Atoi(chi.URLParam(r, "id"))
	project, err := ph.ProjectService.QueryProject(int64(value))
	if err != nil {
		views.Error(err.Error()).Render(r.Context(), w)
		return
	}
	company, err2 := ph.ProjectService.QueryCompanyByCompanyId(project.CompanyID)
	if err2 != nil {
		views.Error(err2.Error()).Render(r.Context(), w)
		return
	}
	views.Project(project, company).Render(r.Context(), w)
}

func (ph *ProjectHandler) HandleEditProject(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	projectId := r.Form.Get("project[ID]")
	companyId := r.Form.Get("company[ID]")
	projectName := r.Form.Get("project[name]")
	projectBudget := r.Form.Get("project[budget]")

	if projectId != "" {
		id, _ := strconv.Atoi(projectId)
		cID, _ := strconv.Atoi(companyId)
		budget, _ := strconv.Atoi(projectBudget)
		ph.ProjectService.EditProject(int64(id), int64(cID), projectName, int64(budget))
		http.Redirect(w, r, fmt.Sprintf("/projects/%v", id), http.StatusTemporaryRedirect)
		return
	}

	value, _ := strconv.Atoi(chi.URLParam(r, "id"))
	project, err := ph.ProjectService.QueryProject(int64(value))
	if err != nil {
		views.Error(err.Error()).Render(r.Context(), w)
		return
	}
	views.EditProject(project).Render(r.Context(), w)
}

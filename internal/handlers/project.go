package handlers

import (
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
	} else {
		views.Projects(projects).Render(r.Context(), w)
	}
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

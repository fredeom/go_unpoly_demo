package handlers

import (
	"net/http"

	"github.com/fredeom/go_unpoly_demo/internal/views"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(r *chi.Mux, hCompany *CompanyHandler, hProject *ProjectHandler, hTask *TaskHandler) {
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		views.Index().Render(r.Context(), w)
	})
	r.Get("/newdemo", hCompany.HandlePopulateStore)
	r.Get("/companies", hCompany.HandleQueryCompanies)
	r.Get("/companies/{id}", hCompany.HandleShowCompany)
	r.Get("/companies/{id}/edit", hCompany.HandleEditCompany)
	r.Post("/companies/{id}", hCompany.HandleDeleteCompany)
	r.Get("/companies/new", hCompany.HandleNewCompany)

	r.Get("/projects", hProject.HandleQueryProjects)
	r.Get("/projects/{id}", hProject.HandleShowProject)
	r.Get("/projects/{id}/edit", hProject.HandleEditProject)
	r.Post("/projects/{id}", hProject.HandleDeleteProject)
	r.Get("/projects/new", hProject.HandleNewProject)

	r.Get("/tasks", hTask.HandleQueryTasks)
	r.Post("/tasks", hTask.HandleNewTask)
	r.Get("/tasks/{id}", hTask.HandleShowTask)
	r.Post("/tasks/{id}/toggle_done", hTask.HandleToggleDoneTask)
	r.Get("/tasks/{id}/edit", hTask.HandleEditTask)
	r.Post("/tasks/delete_done", hTask.HandleDeleteAllDoneTasks)
	r.Get("/tasks/new", hTask.HandleNewTask)
}

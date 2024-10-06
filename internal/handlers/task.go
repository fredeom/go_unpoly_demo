package handlers

import (
	"net/http"
	"strconv"

	"github.com/fredeom/go_unpoly_demo/internal/domain"
	"github.com/fredeom/go_unpoly_demo/internal/views"
	"github.com/go-chi/chi/v5"
)

type TaskService interface {
	QueryTasks(query string) ([]domain.Task, error)
	NewTask(name string) (int64, error)
	DeleteTask(id int64) error
}

type TaskHandler struct {
	TaskService TaskService
}

func NewTaskHandler(ts TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: ts,
	}
}

func (th *TaskHandler) HandleQueryTasks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	tasks, err := th.TaskService.QueryTasks(query)
	if err != nil {
		views.Error(err.Error()).Render(r.Context(), w)
	} else {
		views.Tasks(tasks).Render(r.Context(), w)
	}
}

func (th *TaskHandler) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	method := r.Form.Get("_method")
	if method == "DELETE" {
		value, _ := strconv.Atoi(chi.URLParam(r, "id"))
		err := th.TaskService.DeleteTask(int64(value))
		if err != nil {
			views.Error(err.Error()).Render(r.Context(), w)
			return
		}
		w.Header().Set("X-Up-Events", "[{ \"type\": \"task:destroyed\"}]")
	}
}

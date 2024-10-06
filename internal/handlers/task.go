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

type TaskService interface {
	QueryTasks(query string) ([]domain.Task, error)
	QueryTask(id int64) (domain.Task, error)
	NewTask(name string) (int64, error)
	EditTask(id int64, taskName string, done int) sql.Result
	DeleteAllDoneTasks() error
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

func (th *TaskHandler) HandleShowTask(w http.ResponseWriter, r *http.Request) {
	taskID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	task, _ := th.TaskService.QueryTask(int64(taskID))
	views.Task(task).Render(r.Context(), w)
}

func (th *TaskHandler) HandleToggleDoneTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	method := r.Form.Get("_method")
	if method == "PATCH" {
		taskID, _ := strconv.Atoi(chi.URLParam(r, "id"))
		task, _ := th.TaskService.QueryTask(int64(taskID))
		th.TaskService.EditTask(task.ID, task.Name, 1-task.Done)
		http.Redirect(w, r, fmt.Sprintf("/tasks/%v", taskID), http.StatusSeeOther)
	}
}

func (th *TaskHandler) HandleEditTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	taskID := r.Form.Get("task[ID]")
	taskName := r.Form.Get("task[name]")
	taskDone := r.Form.Get("task[done]")

	if taskID != "" {
		id, _ := strconv.Atoi(taskID)
		done, _ := strconv.Atoi(taskDone)
		th.TaskService.EditTask(int64(id), taskName, done)
		http.Redirect(w, r, fmt.Sprintf("/tasks/%v", id), http.StatusTemporaryRedirect)
		return
	}

	value, _ := strconv.Atoi(chi.URLParam(r, "id"))
	task, err := th.TaskService.QueryTask(int64(value))
	if err != nil {
		views.Error(err.Error()).Render(r.Context(), w)
		return
	}
	views.EditTask(task).Render(r.Context(), w)
}

func (th *TaskHandler) HandleDeleteAllDoneTasks(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	method := r.Form.Get("_method")
	if method == "DELETE" {
		err := th.TaskService.DeleteAllDoneTasks()
		if err != nil {
			views.Error(err.Error()).Render(r.Context(), w)
			return
		}
		w.Header().Set("X-Up-Events", "[{ \"type\": \"tasks:destroyed\"}]")
		w.Header().Set("X-Up-Accept-Layer", "true")
		w.Header().Set("X-Up-Target", "none")
	}
}

func (th *TaskHandler) HandleNewTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	taskName := r.Form.Get("task[name]")

	if taskName == "" {
		views.TaskNew().Render(r.Context(), w)
		return
	}

	_, err1 := th.TaskService.NewTask(taskName)
	if err1 != nil {
		views.Error(err1.Error()).Render(r.Context(), w)
		return
	}
	w.Header().Set("X-Up-Events", "[{ \"type\": \"task:created\"}]")
	w.Header().Set("X-Up-Accept-Layer", "true")
	w.Header().Set("X-Up-Target", "none")
}

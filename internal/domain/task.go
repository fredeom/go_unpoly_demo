package domain

type Task struct {
	ID        int64
	ProjectID int64
	Name      string
	Done      int
}

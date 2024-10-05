package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fredeom/go_unpoly_demo/internal/domain"
)

func (s *Service) NewTask(projectId int64, name string) (int64, error) {
	log.Println("Inserting task record ...")
	newTaskSQL := `INSERT INTO task(project_id, name, done) VALUES (?, ?, 0)`
	statement, err := s.Store.Db.Prepare(newTaskSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var result sql.Result
	result, err = statement.Exec(projectId, name)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return result.LastInsertId()
}

func (s *Service) QueryTasks(query string) ([]domain.Task, error) {
	row, err := s.Store.Db.Query("SELECT * FROM task WHERE name like '%" + query + "%' ORDER BY name LIMIT 10")
	if err != nil {
		return []domain.Task{}, err
	}
	defer row.Close()
	var tasks = []domain.Task{}
	for row.Next() {
		task := domain.Task{}
		row.Scan(&task.ID, &task.ProjectID, &task.Name, &task.Done)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *Service) DeleteTask(id int64) error {
	_, err := s.Store.Db.Exec("DELETE FROM task WHERE id=" + fmt.Sprintf("%v", id))
	return err
}

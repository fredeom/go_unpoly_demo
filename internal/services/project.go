package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fredeom/go_unpoly_demo/internal/domain"
)

func (s *Service) QueryProjects(query string) ([]domain.Project, error) {
	row, err := s.Store.Db.Query("SELECT * FROM project WHERE name like '%" + query + "%' ORDER BY name LIMIT 10")
	if err != nil {
		return []domain.Project{}, err
	}
	defer row.Close()
	var projects = []domain.Project{}
	for row.Next() {
		project := domain.Project{}
		row.Scan(&project.ID, &project.CompanyID, &project.Name, &project.Budget)
		projects = append(projects, project)
	}
	return projects, nil
}

func (s *Service) QueryProjectsByCompanyId(companyId int64) ([]domain.Project, error) {
	stmt, _ := s.Store.Db.Prepare("SELECT * FROM project WHERE company_id=? ORDER BY name")
	row, err := stmt.Query(companyId)
	if err != nil {
		return []domain.Project{}, err
	}
	defer row.Close()
	var projects = []domain.Project{}
	for row.Next() {
		project := domain.Project{}
		row.Scan(&project.ID, &project.CompanyID, &project.Name, &project.Budget)
		projects = append(projects, project)
	}
	return projects, nil
}

func (s *Service) NewProject(companyId int64, name string, budget int64) (int64, error) {
	log.Println("Inserting project record ...")
	newProjectSQL := `INSERT INTO project(company_id, name, budget) VALUES (?, ?, ?)`
	statement, err := s.Store.Db.Prepare(newProjectSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var result sql.Result
	result, err = statement.Exec(companyId, name, budget)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return result.LastInsertId()
}

func (s *Service) DeleteProject(id int64) error {
	_, err := s.Store.Db.Exec("DELETE FROM project WHERE id=" + fmt.Sprintf("%v", id))
	return err
}

func (s *Service) QueryProject(id int64) (domain.Project, error) {
	stmt, _ := s.Store.Db.Prepare("SELECT * FROM project WHERE id=?")
	row, err := stmt.Query(id)
	if err != nil {
		return domain.Project{}, err
	}
	defer row.Close()
	var project = domain.Project{}
	for row.Next() {
		row.Scan(&project.ID, &project.CompanyID, &project.Name, &project.Budget)
	}
	return project, nil
}

func (s *Service) EditProject(id int64, companyId int64, name string, budget int64) sql.Result {
	statement, _ := s.Store.Db.Prepare("UPDATE project SET company_id=?, name=?, budget=? WHERE id=?")
	affected, _ := statement.Exec(companyId, name, budget, id)
	return affected
}

func (s *Service) QueryCompanyNamesByCompanyIDs() (map[int64]string, error) {
	companies, _ := s.QueryCompanies("")
	nameByCompanyID := make(map[int64]string)
	for i := range companies {
		nameByCompanyID[companies[i].ID] = companies[i].Name
	}
	return nameByCompanyID, nil
}

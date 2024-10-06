package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fredeom/go_unpoly_demo/internal/domain"
)

func (s *Service) QueryCompanies(query string) ([]domain.Company, error) {
	row, err := s.Store.Db.Query("SELECT * FROM company WHERE name like '%" + query + "%' ORDER BY name LIMIT 10")
	if err != nil {
		return []domain.Company{}, err
	}
	defer row.Close()
	var companies = []domain.Company{}
	for row.Next() {
		company := domain.Company{}
		row.Scan(&company.ID, &company.Name, &company.Address)
		companies = append(companies, company)
	}
	return companies, nil
}

func (s *Service) QueryCompany(id int64) (domain.Company, error) {
	row, err := s.Store.Db.Query("SELECT * FROM company WHERE id=" + fmt.Sprintf("%v", id))
	if err != nil {
		return domain.Company{}, err
	}
	defer row.Close()
	var company = domain.Company{}
	for row.Next() {
		row.Scan(&company.ID, &company.Name, &company.Address)
	}
	return company, nil
}

func (s *Service) QueryCompanyByCompanyId(companyId int64) (domain.Company, error) {
	return s.QueryCompany(companyId)
}

func (s *Service) EditCompany(id int64, name string, address string) sql.Result {
	statement, _ := s.Store.Db.Prepare("UPDATE company SET name=?, address=? WHERE id=?")
	affected, _ := statement.Exec(name, address, id)
	return affected
}

func (s *Service) DeleteCompany(id int64) error {
	_, err := s.Store.Db.Exec("DELETE FROM company WHERE id=" + fmt.Sprintf("%v", id))
	if err != nil {
		return err
	}
	_, err2 := s.Store.Db.Exec("DELETE FROM project WHERE company_id=" + fmt.Sprintf("%v", id))
	return err2
}

func (s *Service) NewCompany(name string, address string) (int64, error) {
	log.Println("Inserting company record ...")
	newCompanySQL := `INSERT INTO company(name, address) VALUES (?, ?)`
	statement, err := s.Store.Db.Prepare(newCompanySQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var result sql.Result
	result, err = statement.Exec(name, address)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return result.LastInsertId()
}

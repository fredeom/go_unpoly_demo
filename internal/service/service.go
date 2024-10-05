package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fredeom/go_unpoly_demo/internal/db"
	"github.com/fredeom/go_unpoly_demo/internal/domain"
)

type Service struct {
	Store db.Store
}

func New(store db.Store) *Service {
	return &Service{
		Store: store,
	}
}

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

func (s *Service) EditCompany(id int64, name string, address string) sql.Result {
	statement, _ := s.Store.Db.Prepare("UPDATE company SET name=?, address=? WHERE id=?")
	affected, _ := statement.Exec(name, address, id)
	return affected
}

func (s *Service) DeleteCompany(id int64) error {
	_, err := s.Store.Db.Exec("DELETE FROM company WHERE id=" + fmt.Sprintf("%v", id))
	return err
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

func (s *Service) DeleteAllData() error {
	deleteAllCompanySQL := `DELETE FROM company;`
	_, err := s.Store.Db.Exec(deleteAllCompanySQL)
	return err
}

func (s *Service) PopulateStore() error {
	s.DeleteAllData()

	s.NewCompany("Cummerata, Ryan and Senger", "Claudie Extension 37")
	s.NewCompany("Cummerata-Ullrich", "Tyrone Fields 26")
	s.NewCompany("Hessel Group", "Kasey Shores 17")
	s.NewCompany("Stanton, Schoen and Senger", "Emily Mill 93")
	s.NewCompany("Will-Hintz", "Santos Meadow 82")

	return nil
}

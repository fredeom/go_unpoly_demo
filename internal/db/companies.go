package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Company struct {
	ID      int
	Name    string
	Address string
}

func PopulateDemoDB(dbPath string) (*sql.DB, error) {
	os.Remove(dbPath)

	var err error
	DB, err = InitDatabase(dbPath)

	InsertCompany(DB, "Cummerata, Ryan and Senger", "Claudie Extension 37")
	InsertCompany(DB, "Cummerata-Ullrich", "Tyrone Fields 26")
	InsertCompany(DB, "Hessel Group", "Kasey Shores 17")
	InsertCompany(DB, "Stanton, Schoen and Senger", "Emily Mill 93")
	InsertCompany(DB, "Will-Hintz", "Santos Meadow 82")

	DisplayCompanies(DB)

	return DB, err
}

func InsertCompany(db *sql.DB, name string, address string) (int64, error) {
	log.Println("Inserting company record ...")
	insertCompanySQL := `INSERT INTO company(name, address) VALUES (?, ?)`
	statement, err := db.Prepare(insertCompanySQL)
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

func DisplayCompanies(db *sql.DB) {
	row, err := db.Query("SELECT * FROM company ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var name string
		var address string
		row.Scan(&id, &name, &address)
		log.Println("Company: ", name, " ", address)
	}
}

func QueryCompanies(db *sql.DB, query string) ([]Company, error) {
	row, err := db.Query("SELECT * FROM company WHERE name like '%" + query + "%' ORDER BY name LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var companies = []Company{}
	for row.Next() {
		var id int
		var name string
		var address string
		row.Scan(&id, &name, &address)
		companies = append(companies, Company{id, name, address})
	}
	return companies, nil
}

func QueryCompany(db *sql.DB, id int64) (Company, error) {
	row, err := db.Query("SELECT * FROM company WHERE id=" + fmt.Sprintf("%v", id))
	if err != nil {
		return Company{}, err
	}
	defer row.Close()
	var company = Company{}
	for row.Next() {
		row.Scan(&company.ID, &company.Name, &company.Address)
	}
	return company, nil
}

func DeleteCompany(db *sql.DB, id int64) error {
	_, err := db.Exec("DELETE FROM company WHERE id=" + fmt.Sprintf("%v", id))
	return err
}

func EditCompany(db *sql.DB, id int64, name string, address string) sql.Result {
	statement, _ := db.Prepare("UPDATE company SET name=?, address=? WHERE id=?")
	affected, _ := statement.Exec(name, address, int(id))
	return affected
}

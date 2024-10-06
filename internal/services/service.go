package service

import (
	"github.com/fredeom/go_unpoly_demo/internal/db"
)

type Service struct {
	Store db.Store
}

func New(store db.Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) DeleteAllData() error {
	deleteAllCompanySQL := `DELETE FROM company;`
	_, err := s.Store.Db.Exec(deleteAllCompanySQL)
	if err != nil {
		return err
	}
	deleteAllProjectSQL := `DELETE FROM project;`
	_, err2 := s.Store.Db.Exec(deleteAllProjectSQL)
	if err2 != nil {
		return err2
	}
	deleteAllTaskSQL := `DELETE FROM task;`
	_, err3 := s.Store.Db.Exec(deleteAllTaskSQL)
	if err3 != nil {
		return err3
	}
	return nil
}

func (s *Service) PopulateStore() error {
	s.DeleteAllData()

	id1, _ := s.NewCompany("Cummerata, Ryan and Senger", "Claudie Extension 37")
	id2, _ := s.NewCompany("Cummerata-Ullrich", "Tyrone Fields 26")
	id3, _ := s.NewCompany("Hessel Group", "Kasey Shores 17")
	id4, _ := s.NewCompany("Stanton, Schoen and Senger", "Emily Mill 93")
	id5, _ := s.NewCompany("Will-Hintz", "Santos Meadow 82")

	go s.NewProject(id1, "First Project", 1000)
	go s.NewProject(id1, "Some", 1500)
	go s.NewProject(id1, "There", 500)
	go s.NewProject(id1, "Over", 1200)

	go s.NewProject(id2, "The", 1300)
	go s.NewProject(id2, "Rain", 1700)
	go s.NewProject(id2, "Bow", 1450)

	go s.NewProject(id3, "Who", 2700)
	go s.NewProject(id3, "Let", 2400)
	go s.NewProject(id3, "The", 2600)
	go s.NewProject(id3, "Dog", 1200)

	go s.NewProject(id4, "Me", 120)
	go s.NewProject(id4, "We", 220)

	go s.NewProject(id5, "So", 300)
	go s.NewProject(id5, "Be", 300)

	go s.NewTask("Do right")
	go s.NewTask("Do back")
	go s.NewTask("Do free")
	go s.NewTask("Do agree")
	go s.NewTask("Make joy")
	go s.NewTask("Wishful think")
	go s.NewTask("Big deal")
	go s.NewTask("Small deal")
	go s.NewTask("Fly high")
	go s.NewTask("No divide")
	go s.NewTask("Parallel dream")
	go s.NewTask("Make the stream")
	go s.NewTask("Make dinner")
	go s.NewTask("Buy eggs")
	go s.NewTask("Stop procrastenation")

	return nil
}

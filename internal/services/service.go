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

	s.NewProject(id1, "First Project", 1000)
	s.NewProject(id1, "Some", 1500)
	s.NewProject(id1, "There", 500)
	s.NewProject(id1, "Over", 1200)

	s.NewProject(id2, "The", 1300)
	s.NewProject(id2, "Rain", 1700)
	s.NewProject(id2, "Bow", 1450)

	s.NewProject(id3, "Who", 2700)
	s.NewProject(id3, "Let", 2400)
	s.NewProject(id3, "The", 2600)
	s.NewProject(id3, "Dog", 1200)

	s.NewProject(id4, "Me", 120)
	s.NewProject(id4, "We", 220)

	s.NewProject(id5, "So", 300)
	s.NewProject(id5, "Be", 300)

	s.NewTask("Do right")
	s.NewTask("Do back")
	s.NewTask("Do free")
	s.NewTask("Do agree")
	s.NewTask("Make joy")
	s.NewTask("Wishful think")
	s.NewTask("Big deal")
	s.NewTask("Small deal")
	s.NewTask("Fly high")
	s.NewTask("No divide")
	s.NewTask("Parallel dream")
	s.NewTask("Make the stream")
	s.NewTask("Make dinner")
	s.NewTask("Buy eggs")
	s.NewTask("Stop procrastenation")

	return nil
}

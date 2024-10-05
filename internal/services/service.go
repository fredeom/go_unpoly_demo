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

	idp1, _ := s.NewProject(id1, "First Project", 1000)
	idp2, _ := s.NewProject(id1, "Some", 1500)
	idp3, _ := s.NewProject(id1, "There", 500)
	idp4, _ := s.NewProject(id1, "Over", 1200)

	idp5, _ := s.NewProject(id2, "The", 1300)
	idp6, _ := s.NewProject(id2, "Rain", 1700)
	idp7, _ := s.NewProject(id2, "Bow", 1450)

	idp8, _ := s.NewProject(id3, "Who", 2700)
	idp9, _ := s.NewProject(id3, "Let", 2400)
	idp10, _ := s.NewProject(id3, "The", 2600)
	idp11, _ := s.NewProject(id3, "Dog", 1200)

	idp12, _ := s.NewProject(id4, "Me", 120)
	idp13, _ := s.NewProject(id4, "We", 220)

	idp14, _ := s.NewProject(id5, "So", 300)
	idp15, _ := s.NewProject(id5, "Be", 300)

	s.NewTask(idp1, "Do right")
	s.NewTask(idp2, "Do back")
	s.NewTask(idp3, "Do free")
	s.NewTask(idp4, "Do agree")
	s.NewTask(idp5, "Make joy")
	s.NewTask(idp6, "Wishful think")
	s.NewTask(idp7, "Big deal")
	s.NewTask(idp8, "Small deal")
	s.NewTask(idp9, "Fly high")
	s.NewTask(idp10, "No divide")
	s.NewTask(idp11, "Parallel dream")
	s.NewTask(idp12, "Make the stream")
	s.NewTask(idp13, "Make dinner")
	s.NewTask(idp14, "Buy eggs")
	s.NewTask(idp15, "Stop procrastenation")

	return nil
}

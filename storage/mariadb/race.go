package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetRace will grab data from storage
func (s *Storage) GetRace(race *model.Race) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateRace will grab data from storage
func (s *Storage) CreateRace(race *model.Race) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRace will grab data from storage
func (s *Storage) ListRace(page *model.Page) (races []*model.Race, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRaceTotalCount will grab data from storage
func (s *Storage) ListRaceTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRaceBySearch will grab data from storage
func (s *Storage) ListRaceBySearch(page *model.Page, race *model.Race) (races []*model.Race, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRaceBySearchTotalCount will grab data from storage
func (s *Storage) ListRaceBySearchTotalCount(race *model.Race) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditRace will grab data from storage
func (s *Storage) EditRace(race *model.Race) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteRace will grab data from storage
func (s *Storage) DeleteRace(race *model.Race) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

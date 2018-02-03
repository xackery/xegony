package file

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetVariable will grab data from storage
func (s *Storage) GetVariable(variable *model.Variable) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateVariable will grab data from storage
func (s *Storage) CreateVariable(variable *model.Variable) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListVariable will grab data from storage
func (s *Storage) ListVariable(page *model.Page) (variables []*model.Variable, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListVariableTotalCount will grab data from storage
func (s *Storage) ListVariableTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListVariableBySearch will grab data from storage
func (s *Storage) ListVariableBySearch(page *model.Page, variable *model.Variable) (variables []*model.Variable, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListVariableBySearchTotalCount will grab data from storage
func (s *Storage) ListVariableBySearchTotalCount(variable *model.Variable) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditVariable will grab data from storage
func (s *Storage) EditVariable(variable *model.Variable) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteVariable will grab data from storage
func (s *Storage) DeleteVariable(variable *model.Variable) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

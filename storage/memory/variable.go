package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	variablesDatabase = []*model.Variable{}
	variableLock      = sync.RWMutex{}
)

//GetVariable will grab data from storage
func (s *Storage) GetVariable(variable *model.Variable) (err error) {
	variableLock.RLock()
	defer variableLock.RUnlock()
	for _, tmpVariable := range variablesDatabase {
		if tmpVariable.Name == variable.Name {
			*variable = *tmpVariable
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateVariable will grab data from storage
func (s *Storage) CreateVariable(variable *model.Variable) (err error) {
	variableLock.Lock()
	defer variableLock.Unlock()
	for _, tmpVariable := range variablesDatabase {
		if tmpVariable.Name == variable.Name {
			err = fmt.Errorf("variable already exists")
			return
		}
	}
	variablesDatabase = append(variablesDatabase, variable)
	return
}

//ListVariable will grab data from storage
func (s *Storage) ListVariable(page *model.Page) (variables []*model.Variable, err error) {
	variableLock.RLock()
	defer variableLock.RUnlock()

	variables = make([]*model.Variable, len(variablesDatabase))

	variables = variablesDatabase

	switch page.OrderBy {
	case "name":
		sort.Slice(variables, func(i, j int) bool {
			return variables[i].Name < variables[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(variables))
		}
	*/
	return
}

//ListVariableTotalCount will grab data from storage
func (s *Storage) ListVariableTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(variablesDatabase))
	return
}

//ListVariableBySearch will grab data from storage
func (s *Storage) ListVariableBySearch(page *model.Page, variable *model.Variable) (variables []*model.Variable, err error) {
	variableLock.RLock()
	defer variableLock.RUnlock()

	if len(variable.Name) > 0 {
		for i := range variablesDatabase {
			if strings.Contains(variablesDatabase[i].Name, variable.Name) {
				variables = append(variables, variablesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(variables, func(i, j int) bool {
			return variables[i].Name < variables[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(variables))
	//}
	return
}

//ListVariableBySearchTotalCount will grab data from storage
func (s *Storage) ListVariableBySearchTotalCount(variable *model.Variable) (count int64, err error) {
	variableLock.RLock()
	defer variableLock.RUnlock()

	variables := []*model.Variable{}
	if len(variable.Name) > 0 {
		for i := range variablesDatabase {
			if strings.Contains(variablesDatabase[i].Name, variable.Name) {
				variables = append(variables, variablesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(variables))
	return
}

//EditVariable will grab data from storage
func (s *Storage) EditVariable(variable *model.Variable) (err error) {
	variableLock.Lock()
	defer variableLock.Unlock()
	for i := range variablesDatabase {
		if variablesDatabase[i].Name == variable.Name {
			*variablesDatabase[i] = *variable
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteVariable will grab data from storage
func (s *Storage) DeleteVariable(variable *model.Variable) (err error) {
	variableLock.Lock()
	defer variableLock.Unlock()
	indexToDelete := 0
	for i := range variablesDatabase {
		if variablesDatabase[i].Name == variable.Name {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	variablesDatabase[len(variablesDatabase)-1], variablesDatabase[indexToDelete] = variablesDatabase[indexToDelete], variablesDatabase[len(variablesDatabase)-1]
	variablesDatabase = variablesDatabase[:len(variablesDatabase)-1]
	return
}

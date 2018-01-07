package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	recipeFields = `tradeskill_recipe.id, tradeskill_recipe.name, tradeskill_recipe.tradeskill, tradeskill_recipe.skillneeded, tradeskill_recipe.trivial, tradeskill_recipe.nofail, tradeskill_recipe.replace_container, tradeskill_recipe.notes, tradeskill_recipe.must_learn, tradeskill_recipe.quest, tradeskill_recipe.enabled`
	recipeSets   = `id=:id, name=:name, tradeskill=:tradeskill, skillneeded=:skillneeded, trivial=:trivial, nofail=:nofail, replace_container=:replace_container, notes=:notes, must_learn=:must_learn, quest=:quest, enabled=:enabled`
	recipeBinds  = `:id, :name, :tradeskill, :skillneeded, :trivial, :nofail, :replace_container, :notes, :must_learn, :quest, :enabled`
)

//GetRecipe will grab data from storage
func (s *Storage) GetRecipe(recipeID int64) (recipe *model.Recipe, err error) {
	recipe = &model.Recipe{}
	err = s.db.Get(recipe, fmt.Sprintf("SELECT id, %s FROM tradeskill_recipe WHERE id = ?", recipeFields), recipeID)
	if err != nil {
		return
	}
	return
}

//CreateRecipe will grab data from storage
func (s *Storage) CreateRecipe(recipe *model.Recipe) (err error) {
	if recipe == nil {
		err = fmt.Errorf("Must provide recipe")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO tradeskill_recipe(%s)
		VALUES (%s)`, recipeFields, recipeBinds), recipe)
	if err != nil {
		return
	}
	recipeID, err := result.LastInsertId()
	if err != nil {
		return
	}
	recipe.ID = recipeID
	return
}

//ListRecipe will grab data from storage
func (s *Storage) ListRecipe(pageSize int64, pageNumber int64) (recipes []*model.Recipe, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM tradeskill_recipe 
		ORDER BY id ASC LIMIT %d OFFSET %d`, recipeFields, pageSize, pageSize*pageNumber))
	if err != nil {
		return
	}

	for rows.Next() {
		recipe := model.Recipe{}
		if err = rows.StructScan(&recipe); err != nil {
			return
		}
		recipes = append(recipes, &recipe)
	}
	return
}

//ListRecipeCount will grab data from storage
func (s *Storage) ListRecipeCount() (count int64, err error) {
	err = s.db.Get(&count, `SELECT count(id) FROM tradeskill_recipe`)
	if err != nil {
		return
	}
	return
}

//SearchRecipe will grab data from storage
func (s *Storage) SearchRecipe(search string) (recipes []*model.Recipe, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM tradeskill_recipe 
		WHERE name like ? ORDER BY id DESC`, recipeFields), "%"+search+"%")
	if err != nil {
		return
	}

	for rows.Next() {
		recipe := model.Recipe{}
		if err = rows.StructScan(&recipe); err != nil {
			return
		}
		recipes = append(recipes, &recipe)
	}
	return
}

//EditRecipe will grab data from storage
func (s *Storage) EditRecipe(recipeID int64, recipe *model.Recipe) (err error) {
	recipe.ID = recipeID
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE tradeskill_recipe SET %s WHERE id = :id`, recipeSets), recipe)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}

//DeleteRecipe will grab data from storage
func (s *Storage) DeleteRecipe(recipeID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM tradeskill_recipe WHERE id = ?`, recipeID)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}
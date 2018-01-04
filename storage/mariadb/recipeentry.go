package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	recipeEntryFields = `id, recipe_id, item_id, successcount, failcount, componentcount, salvagecount, iscontainer`
	recipeEntrySets   = `id=:id, recipe_id=:recipe_id, item_id=:item_id, successcount=:successcount, failcount=:failcount, componentcount=:componentcount, salvagecount=:salvagecount, iscontainer=:iscontainer`
	recipeEntryBinds  = `:id, :recipe_id, :item_id, :successcount, :failcount, :componentcount, :salvagecount, :iscontainer`
)

//GetRecipeEntry will grab data from storage
func (s *Storage) GetRecipeEntry(recipeID int64, itemID int64) (query string, recipeEntry *model.RecipeEntry, err error) {
	recipeEntry = &model.RecipeEntry{}
	query = fmt.Sprintf(`SELECT %s FROM tradeskill_recipe_entries 
		WHERE tradeskill_recipe_entries.recipe_id = ? AND tradeskill_recipe_entries.item_id = ?`, recipeEntryFields)
	err = s.db.Get(recipeEntry, query, recipeID, itemID)
	if err != nil {
		return
	}
	return
}

//CreateRecipeEntry will grab data from storage
func (s *Storage) CreateRecipeEntry(recipeEntry *model.RecipeEntry) (query string, err error) {
	if recipeEntry == nil {
		err = fmt.Errorf("Must provide recipeEntry")
		return
	}

	query = fmt.Sprintf(`INSERT INTO tradeskill_recipe_entries(%s)
		VALUES (%s)`, recipeEntryFields, recipeEntryBinds)
	_, err = s.db.NamedExec(query, recipeEntry)
	if err != nil {
		return
	}
	return
}

//ListRecipeEntry will grab data from storage
func (s *Storage) ListRecipeEntry(recipeID int64) (query string, recipeEntrys []*model.RecipeEntry, err error) {
	query = fmt.Sprintf(`SELECT %s FROM tradeskill_recipe_entries WHERE recipe_id = ?`, recipeEntryFields)
	rows, err := s.db.Queryx(query, recipeID)
	if err != nil {
		return
	}

	for rows.Next() {
		recipeEntry := model.RecipeEntry{}
		if err = rows.StructScan(&recipeEntry); err != nil {
			return
		}
		recipeEntrys = append(recipeEntrys, &recipeEntry)
	}
	return
}

//ListRecipeEntryByItem will grab data from storage
func (s *Storage) ListRecipeEntryByItem(itemID int64) (query string, recipeEntrys []*model.RecipeEntry, err error) {

	query = fmt.Sprintf(`SELECT %s FROM tradeskill_recipe_entries
	WHERE item_id = ?`, recipeEntryFields)

	rows, err := s.db.Queryx(query, itemID)
	if err != nil {
		return
	}

	for rows.Next() {
		recipeEntry := model.RecipeEntry{}
		if err = rows.StructScan(&recipeEntry); err != nil {
			return
		}
		recipeEntrys = append(recipeEntrys, &recipeEntry)
	}
	return
}

//EditRecipeEntry will grab data from storage
func (s *Storage) EditRecipeEntry(recipeID int64, itemID int64, recipeEntry *model.RecipeEntry) (query string, err error) {

	query = fmt.Sprintf(`UPDATE tradeskill_recipe_entries SET %s WHERE tradeskill_recipe_entries.recipe_id = ? AND tradeskill_recipe_entries.item_id = ?`, recipeEntrySets)
	recipeEntry.RecipeID = recipeID
	recipeEntry.ItemID = itemID
	result, err := s.db.NamedExec(query, recipeEntry)
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

//DeleteRecipeEntry will grab data from storage
func (s *Storage) DeleteRecipeEntry(recipeID int64, itemID int64) (query string, err error) {
	query = `DELETE FROM tradeskill_recipe_entries WHERE recipe_id = ? AND item_id = ?`
	result, err := s.db.Exec(query, recipeID, itemID)
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

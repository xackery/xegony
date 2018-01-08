package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listRecipe(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       site
		Recipes    []*model.Recipe
		RecipePage *model.Page
	}

	site := a.newSite(r)
	site.Page = "recipe"
	site.Title = "Recipe List"

	recipePage := &model.Page{
		Scope: "recipe",
	}
	recipePage.PageSize = getIntParam(r, "pageSize")
	recipePage.PageNumber = getIntParam(r, "pageNumber")

	recipes, err := a.recipeRepo.List(recipePage.PageSize, recipePage.PageNumber)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	recipePage.Total, err = a.recipeRepo.ListCount()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	for _, recipe := range recipes {
		recipe.Entrys, _, err = a.recipeEntryRepo.List(recipe.ID)
		if err != nil {
			err = errors.Wrap(err, "failed to get entry")
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
		for _, entry := range recipe.Entrys {
			entry.Item, err = a.itemRepo.Get(entry.ItemID)
			if err != nil {
				continue
				//err = errors.Wrap(err, "recipeID argument is required")
				//a.writeError(w, r, err, http.StatusBadRequest)
				//return
			}
		}
	}
	content := Content{
		Site:       site,
		Recipes:    recipes,
		RecipePage: recipePage,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipe/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("recipe", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) listRecipeByTradeskill(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site        site
		Tradeskills []*model.Skill
	}

	site := a.newSite(r)
	site.Page = "recipe"
	site.Title = "Recipe By Tradeskill"

	tradeskills, err := a.skillRepo.ListByType(1)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:        site,
		Tradeskills: tradeskills,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipe/listbytradeskill.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("itemlistbyzone", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getRecipeByTradeskill(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       site
		Recipes    []*model.Recipe
		RecipePage *model.Page
		Skill      *model.Skill
	}

	tradeskillID, err := getIntVar(r, "tradeskillID")
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	skill, err := a.skillRepo.Get(tradeskillID)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("Could not get skill %d", tradeskillID))
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "recipe"
	site.Title = fmt.Sprintf("%s Recipes", skill.Name)

	recipePage := &model.Page{
		Scope: fmt.Sprintf("recipe/bytradeskill/%d", tradeskillID),
	}
	recipePage.PageSize = getIntParam(r, "pageSize")
	recipePage.PageNumber = getIntParam(r, "pageNumber")

	recipes, err := a.recipeRepo.ListByTradeskill(tradeskillID, recipePage.PageSize, recipePage.PageNumber)
	if err != nil {
		err = errors.Wrap(err, "Failed to get recipes")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	recipePage.Total, err = a.recipeRepo.ListByTradeskillCount(tradeskillID)
	if err != nil {
		err = errors.Wrap(err, "Failed to get recipepage")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	for _, recipe := range recipes {
		recipe.Entrys, _, err = a.recipeEntryRepo.List(recipe.ID)
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("failed to get recipe entry %d", recipe.ID))
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
		for _, entry := range recipe.Entrys {
			entry.Item, err = a.itemRepo.Get(entry.ItemID)
			if err != nil {
				continue
				//err = errors.Wrap(err, fmt.Sprintf("failed to get item entry %d", entry.ItemID))
				//a.writeError(w, r, err, http.StatusBadRequest)
				//return
			}
		}
	}
	content := Content{
		Site:       site,
		Recipes:    recipes,
		RecipePage: recipePage,
		Skill:      skill,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipe/getbytradeskill.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("recipe", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getRecipe(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site         site
		Recipe       *model.Recipe
		RecipeEntrys []*model.RecipeEntry
	}

	if strings.ToLower(getVar(r, "recipeID")) == "bytradeskill" {
		a.listRecipeByTradeskill(w, r)
		return
	}

	if strings.ToLower(getVar(r, "recipeID")) == "search" {
		a.searchRecipe(w, r)
		return
	}

	recipeID, err := getIntVar(r, "recipeID")
	if err != nil {
		err = errors.Wrap(err, "recipeID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	recipe, err := a.recipeRepo.Get(recipeID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	recipe.Entrys, _, err = a.recipeEntryRepo.List(recipeID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "recipe"
	site.Title = fmt.Sprintf("Recipe: %s", recipe.Name)

	content := Content{
		Site:   site,
		Recipe: recipe,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipe/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("recipe", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) searchRecipe(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site    site
		Recipes []*model.Recipe
		Search  string
	}

	search := getParam(r, "search")

	var recipes []*model.Recipe

	if len(search) > 0 {
		recipes, err = a.recipeRepo.Search(search)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}

	site := a.newSite(r)
	site.Page = "recipe"
	site.Title = "Recipe"

	content := Content{
		Site:    site,
		Recipes: recipes,
		Search:  search,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipe/search.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("recipesearch", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listRecipe(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site    site
		Recipes []*model.Recipe
	}

	site := a.newSite(r)
	site.Page = "recipe"
	site.Title = "Recipe"

	pageSize := getIntParam(r, "pageSize")
	pageNumber := getIntParam(r, "pageNumber")

	recipes, err := a.recipeRepo.List(pageSize, pageNumber)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	site.ResultCount, err = a.recipeRepo.ListCount()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	for _, recipe := range recipes {
		recipe.Entrys, _, err = a.recipeEntryRepo.List(recipe.ID)
		if err != nil {
			err = errors.Wrap(err, "recipeID argument is required")
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
		for _, entry := range recipe.Entrys {
			entry.Item, err = a.itemRepo.Get(entry.ItemID)
			if err != nil {
				err = errors.Wrap(err, "recipeID argument is required")
				a.writeError(w, r, err, http.StatusBadRequest)
				return
			}
		}
	}
	content := Content{
		Site:    site,
		Recipes: recipes,
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

func (a *Web) getRecipe(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site         site
		Recipe       *model.Recipe
		RecipeEntrys []*model.RecipeEntry
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
	site.Title = "Recipe"

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

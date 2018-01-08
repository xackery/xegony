package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listRecipeEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if strings.ToLower(getVar(r, "recipeID")) == "bytradeskill" {
		a.listRecipeByTradeskill(w, r)
		return
	}

	recipeID, err := getIntVar(r, "recipeID")
	if err != nil {
		err = errors.Wrap(err, "recipeEntryID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	recipe, err := a.recipeRepo.Get(recipeID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	type Content struct {
		Site   site
		Recipe *model.Recipe
	}

	site := a.newSite(r)
	site.Page = "recipeentry"
	site.Title = fmt.Sprintf("Recipe: %s", recipe.Name)
	site.Section = "recipeentry"

	recipe.Entrys, _, err = a.recipeEntryRepo.List(recipe.ID)
	if err != nil {
		err = errors.Wrap(err, "recipeID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	for _, entry := range recipe.Entrys {
		entry.Item, err = a.itemRepo.Get(entry.ItemID)
		if err != nil {
			continue
			//err = errors.Wrap(err, "recipeID argument is required")
			//a.writeError(w, r, err, http.StatusBadRequest)

		}
	}

	content := Content{
		Site:   site,
		Recipe: recipe,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipeentry/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("recipeentry", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getRecipeEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site        site
		RecipeEntry *model.RecipeEntry
		Npc         *model.Npc
	}

	recipeID, err := getIntVar(r, "recipeID")
	if err != nil {
		err = errors.Wrap(err, "recipeID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	if strings.ToLower(getVar(r, "npcID")) == "details" {
		a.getRecipe(w, r)
		return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	recipeEntry, _, err := a.recipeEntryRepo.Get(recipeID, npcID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	npc, err := a.npcRepo.Get(npcID)
	if err != nil {
		err = errors.Wrap(err, "Request error for npc")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	site := a.newSite(r)
	site.Page = "recipeentry"
	site.Title = "recipeentry"
	site.Section = "recipeentry"

	content := Content{
		Site:        site,
		RecipeEntry: recipeEntry,
		Npc:         npc,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipeentry/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("recipeentry", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

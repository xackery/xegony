package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) recipeEntryRoutes() (routes []*route) {
	routes = []*route{
		//RecipeEntry
		{
			"ListRecipeEntry",
			"GET",
			"/recipe/{recipeID:[0-9]+}",
			a.listRecipeEntry,
		},
		{
			"GetRecipeEntry",
			"GET",
			"/spawn/{recipeID:[0-9]+}/{recipeEntryID:[0-9]+}",
			a.getRecipeEntry,
		},
	}
	return
}

func (a *Web) listRecipeEntry(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	recipeID, err := getIntVar(r, "recipeID")
	if err != nil {
		err = errors.Wrap(err, "recipeEntryID argument is required")
		return
	}

	recipe := &model.Recipe{
		ID: recipeID,
	}
	err = a.recipeRepo.Get(recipe, user)
	if err != nil {
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

	recipe.Entrys, err = a.recipeEntryRepo.ListByRecipe(recipe, user)
	if err != nil {
		err = errors.Wrap(err, "recipeID argument is required")
		return
	}
	for _, entry := range recipe.Entrys {
		err = a.itemRepo.Get(entry.Item, user)
		if err != nil {
			continue
			//err = errors.Wrap(err, "recipeID argument is required")
			//
		}
	}

	content = Content{
		Site:   site,
		Recipe: recipe,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipeentry/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("recipeentry", tmp)
	}

	return
}

func (a *Web) getRecipeEntry(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site        site
		RecipeEntry *model.RecipeEntry
		Npc         *model.Npc
	}

	recipeID, err := getIntVar(r, "recipeID")
	if err != nil {
		err = errors.Wrap(err, "recipeID argument is required")
		return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		return
	}
	recipeEntry := &model.RecipeEntry{
		RecipeID: recipeID,
	}

	err = a.recipeEntryRepo.Get(recipeEntry, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	npc := &model.Npc{
		ID: npcID,
	}
	err = a.npcRepo.Get(npc, user)
	if err != nil {
		err = errors.Wrap(err, "Request error for npc")
		return
	}
	site := a.newSite(r)
	site.Page = "recipeentry"
	site.Title = "recipeentry"
	site.Section = "recipeentry"

	content = Content{
		Site:        site,
		RecipeEntry: recipeEntry,
		Npc:         npc,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipeentry/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("recipeentry", tmp)
	}

	return
}

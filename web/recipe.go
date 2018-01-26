package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) recipeRoutes() (routes []*route) {
	routes = []*route{
		//Recipe
		{
			"ListRecipe",
			"GET",
			"/recipe",
			a.listRecipe,
		},
		{
			"GetRecipe",
			"GET",
			"/recipe/{recipeID}/details",
			a.getRecipe,
		},
		{
			"GetRecipeByTradeskill",
			"GET",
			"/recipe/bytradeskill/{tradeskillID}",
			a.getRecipeByTradeskill,
		},
	}
	return
}

func (a *Web) listRecipe(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

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

	recipes, err := a.recipeRepo.List(recipePage.PageSize, recipePage.PageNumber, user)
	if err != nil {
		return
	}
	recipePage.Total, err = a.recipeRepo.ListCount(user)
	if err != nil {
		return
	}

	for _, recipe := range recipes {
		recipe.Entrys, err = a.recipeEntryRepo.ListByRecipe(recipe, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get entry")
			return
		}
		for _, entry := range recipe.Entrys {
			entry.Item = &model.Item{
				ID: entry.ItemID,
			}
			err = a.itemRepo.Get(entry.Item, user)
			if err != nil {
				continue
				//err = errors.Wrap(err, "recipeID argument is required")
				//				//return
			}
		}
	}
	content = Content{
		Site:       site,
		Recipes:    recipes,
		RecipePage: recipePage,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipe/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("recipe", tmp)
	}

	return
}

func (a *Web) listRecipeByTradeskill(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site        site
		Tradeskills []*model.Skill
	}

	site := a.newSite(r)
	site.Page = "recipe"
	site.Title = "Recipe By Tradeskill"

	skill := &model.Skill{
		Type: 1,
	}
	tradeskills, err := a.skillRepo.ListByType(skill, user)
	if err != nil {
		return
	}
	content = Content{
		Site:        site,
		Tradeskills: tradeskills,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipe/listbytradeskill.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("itemlistbyzone", tmp)
	}

	return
}

func (a *Web) getRecipeByTradeskill(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site       site
		Recipes    []*model.Recipe
		RecipePage *model.Page
		Skill      *model.Skill
	}

	tradeskillID, err := getIntVar(r, "tradeskillID")
	if err != nil {
		return
	}

	skill := &model.Skill{
		ID: tradeskillID,
	}
	err = a.skillRepo.Get(skill, user)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("Could not get skill %d", skill.ID))
		return
	}

	site := a.newSite(r)
	site.Page = "recipe"
	site.Title = fmt.Sprintf("%s Recipes", skill.Name)

	recipePage := &model.Page{
		Scope: fmt.Sprintf("recipe/bytradeskill/%d", skill.ID),
	}
	recipePage.PageSize = getIntParam(r, "pageSize")
	recipePage.PageNumber = getIntParam(r, "pageNumber")

	recipes, err := a.recipeRepo.ListBySkill(skill, recipePage.PageSize, recipePage.PageNumber, user)
	if err != nil {
		err = errors.Wrap(err, "Failed to get recipes")
		return
	}
	recipePage.Total, err = a.recipeRepo.ListBySkillCount(skill, user)
	if err != nil {
		err = errors.Wrap(err, "Failed to get recipepage")
		return
	}

	for _, recipe := range recipes {
		recipe.Entrys, err = a.recipeEntryRepo.ListByRecipe(recipe, user)
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("failed to get recipe entry %d", recipe.ID))
			return
		}
		for _, entry := range recipe.Entrys {
			err = a.itemRepo.Get(entry.Item, user)
			if err != nil {
				continue
				//err = errors.Wrap(err, fmt.Sprintf("failed to get item entry %d", entry.ItemID))
				//				//return
			}
		}
	}
	content = Content{
		Site:       site,
		Recipes:    recipes,
		RecipePage: recipePage,
		Skill:      skill,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipe/getbytradeskill.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("recipe", tmp)
	}

	return
}

func (a *Web) getRecipe(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site         site
		Recipe       *model.Recipe
		RecipeEntrys []*model.RecipeEntry
	}

	recipeID, err := getIntVar(r, "recipeID")
	if err != nil {
		err = errors.Wrap(err, "recipeID argument is required")
		return
	}

	recipe := &model.Recipe{
		ID: recipeID,
	}
	err = a.recipeRepo.Get(recipe, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	recipe.Entrys, err = a.recipeEntryRepo.ListByRecipe(recipe, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "recipe"
	site.Title = fmt.Sprintf("Recipe: %s", recipe.Name)

	content = Content{
		Site:   site,
		Recipe: recipe,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipe/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("recipe", tmp)
	}

	return
}

func (a *Web) searchRecipe(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site    site
		Recipes []*model.Recipe
		Search  string
	}

	search := getParam(r, "search")

	var recipes []*model.Recipe

	if len(search) > 0 {
		recipe := &model.Recipe{
			Name: search,
		}
		recipes, err = a.recipeRepo.SearchByName(recipe, user)
		if err != nil {
			return
		}
	}

	site := a.newSite(r)
	site.Page = "recipe"
	site.Title = "Recipe"

	content = Content{
		Site:    site,
		Recipes: recipes,
		Search:  search,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "recipe/search.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("recipesearch", tmp)
	}

	return
}

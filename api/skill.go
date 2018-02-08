package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SkillRequest is a list of parameters used for skill
// swagger:parameters deleteSkill editSkill getSkill
type SkillRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
}

// SkillResponse is what endpoints respond with
// swagger:response
type SkillResponse struct {
	Skill *model.Skill `json:"skill,omitempty"`
}

// SkillCreateRequest is the body parameters for creating an skill
// swagger:parameters createSkill
type SkillCreateRequest struct {
	// Skill details to create
	// in: body
	Skill *model.Skill `json:"skill"`
}

// SkillEditRequest is the body parameters for creating an skill
// swagger:parameters editSkill
type SkillEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// Skill details to edit
	// in: body
	Skill *model.Skill `json:"skill"`
}

// SkillsRequest is a list of parameters used for skill
// swagger:parameters listSkill
type SkillsRequest struct {
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: id
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// SkillsResponse is a general response to a request
// swagger:response
type SkillsResponse struct {
	Page   *model.Page  `json:"page,omitempty"`
	Skills model.Skills `json:"skills,omitempty"`
}

// SkillsBySearchRequest is a list of parameters used for skill
// swagger:parameters listSkillBySearch
type SkillsBySearchRequest struct {
	// Name is which skill to get information about
	// example: xackery
	// in: query
	Name string `json:"name"`
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: id
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// SkillsBySearchResponse is a general response to a request
// swagger:response
type SkillsBySearchResponse struct {
	Search *model.Skill `json:"search,omitempty"`
	Page   *model.Page  `json:"page,omitempty"`
	Skills model.Skills `json:"skills,omitempty"`
}

func skillRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /skill skill listSkill
		//
		// Lists skills
		//
		// This will show all available skills by default.
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//     - application/xml
		//     - application/yaml
		//
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SkillsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSkill",
			"GET",
			"/skill",
			listSkill,
		},
		// swagger:route GET /skill/search skill listSkillBySearch
		//
		// Search skills by name
		//
		// This will show all available skills by default.
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//     - application/xml
		//     - application/yaml
		//
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SkillsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSkillBySearch",
			"GET",
			"/skill/search",
			listSkillBySearch,
		},
		// swagger:route POST /skill skill createSkill
		//
		// Create an skill
		//
		// This will create an skill
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SkillResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSkill",
			"POST",
			"/skill",
			createSkill,
		},
		// swagger:route GET /skill/{ID} skill getSkill
		//
		// Get an skill
		//
		// This will get an individual skill available skills by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SkillResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSkill",
			"GET",
			"/skill/{ID:[0-9]+}",
			getSkill,
		},
		// swagger:route PUT /skill/{ID} skill editSkill
		//
		// Edit an skill
		//
		// This will edit an skill
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SkillResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSkill",
			"PUT",
			"/skill/{ID:[0-9]+}",
			editSkill,
		},
		// swagger:route DELETE /skill/{ID} skill deleteSkill
		//
		// Delete an skill
		//
		// This will delete an skill
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"DeleteSkill",
			"DELETE",
			"/skill/{ID:[0-9]+}",
			deleteSkill,
		},
	}
	return
}

func getSkill(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SkillRequest{
		ID: getIntVar(r, "ID"),
	}

	skill := &model.Skill{
		ID: request.ID,
	}

	err = cases.GetSkill(skill, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SkillResponse{
		Skill: skill,
	}
	content = response
	return
}

func createSkill(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	skill := &model.Skill{}
	err = decodeBody(r, skill)
	if err != nil {
		return
	}
	err = cases.CreateSkill(skill, user)
	if err != nil {
		return
	}
	response := &SkillResponse{
		Skill: skill,
	}
	content = response
	return
}

func deleteSkill(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SkillRequest{
		ID: getIntVar(r, "ID"),
	}

	skill := &model.Skill{
		ID: request.ID,
	}

	err = cases.DeleteSkill(skill, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
	}
	err = &model.ErrNoContent{}
	return
}

func editSkill(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SkillEditRequest{
		ID: getIntVar(r, "ID"),
	}

	skill := &model.Skill{
		ID: request.ID,
	}

	err = decodeBody(r, skill)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditSkill(skill, user)
	if err != nil {
		return
	}
	response := &SkillResponse{
		Skill: skill,
	}
	content = response
	return
}

func listSkill(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	skills, err := cases.ListSkill(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SkillsResponse{
		Page:   page,
		Skills: skills,
	}
	content = response
	return
}

func listSkillBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	skill := &model.Skill{
		Name: getQuery(r, "name"),
	}
	skills, err := cases.ListSkillBySearch(page, skill, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SkillsBySearchResponse{
		Page:   page,
		Skills: skills,
		Search: skill,
	}
	content = response
	return
}

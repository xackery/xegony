package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SlotRequest is a list of parameters used for slot
// swagger:parameters deleteSlot editSlot getSlot
type SlotRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
}

// SlotResponse is what endpoints respond with
// swagger:response
type SlotResponse struct {
	Slot *model.Slot `json:"slot,omitempty"`
}

// SlotCreateRequest is the body parameters for creating an slot
// swagger:parameters createSlot
type SlotCreateRequest struct {
	// Slot details to create
	// in: body
	Slot *model.Slot `json:"slot"`
}

// SlotEditRequest is the body parameters for creating an slot
// swagger:parameters editSlot
type SlotEditRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
	// Slot details to edit
	// in: body
	Slot *model.Slot `json:"slot"`
}

// SlotsRequest is a list of parameters used for slot
// swagger:parameters listSlot
type SlotsRequest struct {
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: name
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// SlotsResponse is a general response to a request
// swagger:response
type SlotsResponse struct {
	Page  *model.Page `json:"page,omitempty"`
	Slots model.Slots `json:"slots,omitempty"`
}

// SlotsBySearchRequest is a list of parameters used for slot
// swagger:parameters listSlotBySearch
type SlotsBySearchRequest struct {
	// Name is which slot to get information about
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
	// example: name
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// SlotsBySearchResponse is a general response to a request
// swagger:response
type SlotsBySearchResponse struct {
	Search *model.Slot `json:"search,omitempty"`
	Page   *model.Page `json:"page,omitempty"`
	Slots  model.Slots `json:"slots,omitempty"`
}

func slotRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /slot slot listSlot
		//
		// Lists slots
		//
		// This will show all available slots by default.
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
		//       200: SlotsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSlot",
			"GET",
			"/slot",
			listSlot,
		},
		// swagger:route GET /slot/search slot listSlotBySearch
		//
		// Search slots by name
		//
		// This will show all available slots by default.
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
		//       200: SlotsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSlotBySearch",
			"GET",
			"/slot/search",
			listSlotBySearch,
		},
		// swagger:route POST /slot slot createSlot
		//
		// Create an slot
		//
		// This will create an slot
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SlotResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSlot",
			"POST",
			"/slot",
			createSlot,
		},
		// swagger:route GET /slot/{ID} slot getSlot
		//
		// Get an slot
		//
		// This will get an individual slot available slots by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SlotResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSlot",
			"GET",
			"/slot/{ID:[0-9]+}",
			getSlot,
		},
		// swagger:route PUT /slot/{ID} slot editSlot
		//
		// Edit an slot
		//
		// This will edit an slot
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SlotResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSlot",
			"PUT",
			"/slot/{ID:[0-9]+}",
			editSlot,
		},
		// swagger:route DELETE /slot/{ID} slot deleteSlot
		//
		// Delete an slot
		//
		// This will delete an slot
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
			"DeleteSlot",
			"DELETE",
			"/slot/{ID:[0-9]+}",
			deleteSlot,
		},
	}
	return
}

func getSlot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SlotRequest{
		ID: getIntVar(r, "ID"),
	}

	slot := &model.Slot{
		ID: request.ID,
	}

	err = cases.GetSlot(slot, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SlotResponse{
		Slot: slot,
	}
	content = response
	return
}

func createSlot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	slot := &model.Slot{}
	err = decodeBody(r, slot)
	if err != nil {
		return
	}
	err = cases.CreateSlot(slot, user)
	if err != nil {
		return
	}
	response := &SlotResponse{
		Slot: slot,
	}
	content = response
	return
}

func deleteSlot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SlotRequest{
		ID: getIntVar(r, "ID"),
	}

	slot := &model.Slot{
		ID: request.ID,
	}

	err = cases.DeleteSlot(slot, user)
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

func editSlot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SlotEditRequest{
		ID: getIntVar(r, "ID"),
	}

	slot := &model.Slot{
		ID: request.ID,
	}

	err = decodeBody(r, slot)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditSlot(slot, user)
	if err != nil {
		return
	}
	response := &SlotResponse{
		Slot: slot,
	}
	content = response
	return
}

func listSlot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	slots, err := cases.ListSlot(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SlotsResponse{
		Page:  page,
		Slots: slots,
	}
	content = response
	return
}

func listSlotBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	slot := &model.Slot{
		Name: getQuery(r, "name"),
	}
	slots, err := cases.ListSlotBySearch(page, slot, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SlotsBySearchResponse{
		Page:   page,
		Slots:  slots,
		Search: slot,
	}
	content = response
	return
}

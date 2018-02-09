package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// ZoneImageRequest is a list of parameters used for zoneImage
// swagger:parameters deleteZoneImage editZoneImage getZoneImage
type ZoneImageRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
}

// ZoneImageResponse is what endpoints respond with
// swagger:response
type ZoneImageResponse struct {
	ZoneImage *model.ZoneImage `json:"zoneImage,omitempty"`
}

// ZoneImageCreateRequest is the body parameters for creating an zoneImage
// swagger:parameters createZoneImage
type ZoneImageCreateRequest struct {
	// ZoneImage details to create
	// in: body
	ZoneImage *model.ZoneImage `json:"zoneImage"`
}

// ZoneImageEditRequest is the body parameters for creating an zoneImage
// swagger:parameters editZoneImage
type ZoneImageEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// ZoneImage details to edit
	// in: body
	ZoneImage *model.ZoneImage `json:"zoneImage"`
}

// ZoneImagesRequest is a list of parameters used for zoneImage
// swagger:parameters listZoneImage
type ZoneImagesRequest struct {
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

// ZoneImagesResponse is a general response to a request
// swagger:response
type ZoneImagesResponse struct {
	Page       *model.Page      `json:"page,omitempty"`
	ZoneImages model.ZoneImages `json:"zoneImages,omitempty"`
}

// ZoneImagesBySearchRequest is a list of parameters used for zoneImage
// swagger:parameters listZoneImageBySearch
type ZoneImagesBySearchRequest struct {
	// ID is which zoneImage to get information about
	// example: xackery
	// in: query
	ID string `json:"ID"`
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

// ZoneImagesBySearchResponse is a general response to a request
// swagger:response
type ZoneImagesBySearchResponse struct {
	Search     *model.ZoneImage `json:"search,omitempty"`
	Page       *model.Page      `json:"page,omitempty"`
	ZoneImages model.ZoneImages `json:"zoneImages,omitempty"`
}

// ZoneImageBotEditRequest is the body parameters for creating an zoneImage
// swagger:parameters editZoneImage
type ZoneImageBotEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// Bot details to edit
	// in: body
	Bot *model.Bot `json:"bot"`
}

// ZoneImageBotResponse is what endpoints respond with
// swagger:response
type ZoneImageBotResponse struct {
	Bot *model.Bot `json:"bot,omitempty"`
}

// ZoneImageBotsResponse is what endpoints respond with
// swagger:response
type ZoneImageBotsResponse struct {
	Bots model.Bots `json:"bots,omitempty"`
}

func zoneImageRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /zone/image/bot zoneImage listZoneImageBot
		//
		// Lists zoneImage bots
		//
		// This will show all available bots that work on zone image by default.
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
		//       200: ZoneImagesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListZoneImageBot",
			"GET",
			"/zone/image/bot",
			listZoneImage,
		},
		// swagger:route GET /zone/image/bot zoneImage listZoneImageBot
		//
		// Get zoneImage bot
		//
		// This will get the bot with the matching ID
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
		//       200: ZoneImagesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetZoneImageBot",
			"GET",
			"/zone/image/bot/{botID:[0-9]+}",
			getZoneImageBot,
		},
		// swagger:route PUT /zone/image/bot zoneImage editZoneImageBot
		//
		// Edit zoneImage bot
		//
		// This will edit the bot with the matching ID
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
		//       200: ZoneImagesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditZoneImageBot",
			"PUT",
			"/zone/image/bot/{botID:[0-9]+}",
			editZoneImageBot,
		},
		// swagger:route GET /zone/image zoneImage listZoneImage
		//
		// Lists zoneImages
		//
		// This will show all available zoneImages by default.
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
		//       200: ZoneImagesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListZoneImage",
			"GET",
			"/zone/image",
			listZoneImage,
		},
		// swagger:route GET /zone/image/search zoneImage listZoneImageBySearch
		//
		// Search zoneImages by id
		//
		// This will show all available zoneImages by default.
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
		//       200: ZoneImagesBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListZoneImageBySearch",
			"GET",
			"/zone/image/search",
			listZoneImageBySearch,
		},
		// swagger:route POST /zone/image zoneImage createZoneImage
		//
		// Create an zoneImage
		//
		// This will create an zoneImage
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ZoneImageResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateZoneImage",
			"POST",
			"/zone/image",
			createZoneImage,
		},
		// swagger:route GET /zone/image/{ID} zoneImage getZoneImage
		//
		// Get an zoneImage
		//
		// This will get an individual zoneImage available zoneImages by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: ZoneImageResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetZoneImage",
			"GET",
			"/zone/image/{ID:[0-9]+}",
			getZoneImage,
		},
		// swagger:route PUT /zone/image/{ID} zoneImage editZoneImage
		//
		// Edit an zoneImage
		//
		// This will edit an zoneImage
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ZoneImageResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditZoneImage",
			"PUT",
			"/zone/image/{ID:[0-9]+}",
			editZoneImage,
		},
		// swagger:route DELETE /zone/image/{ID} zoneImage deleteZoneImage
		//
		// Delete an zoneImage
		//
		// This will delete an zoneImage
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
			"DeleteZoneImage",
			"DELETE",
			"/zone/image/{ID:[0-9]+}",
			deleteZoneImage,
		},
	}
	return
}

func getZoneImage(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ZoneImageRequest{
		ID: getIntVar(r, "ID"),
	}

	zoneImage := &model.ZoneImage{
		ID: request.ID,
	}

	err = cases.GetZoneImage(zoneImage, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &ZoneImageResponse{
		ZoneImage: zoneImage,
	}
	content = response
	return
}

func createZoneImage(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	zoneImage := &model.ZoneImage{}
	err = decodeBody(r, zoneImage)
	if err != nil {
		return
	}
	err = cases.CreateZoneImage(zoneImage, user)
	if err != nil {
		return
	}
	response := &ZoneImageResponse{
		ZoneImage: zoneImage,
	}
	content = response
	return
}

func deleteZoneImage(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ZoneImageRequest{
		ID: getIntVar(r, "ID"),
	}

	zoneImage := &model.ZoneImage{
		ID: request.ID,
	}

	err = cases.DeleteZoneImage(zoneImage, user)
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

func editZoneImage(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ZoneImageEditRequest{
		ID: getIntVar(r, "ID"),
	}

	zoneImage := &model.ZoneImage{
		ID: request.ID,
	}

	err = decodeBody(r, zoneImage)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditZoneImage(zoneImage, user)
	if err != nil {
		return
	}
	response := &ZoneImageResponse{
		ZoneImage: zoneImage,
	}
	content = response
	return
}

func listZoneImage(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	zoneImages, err := cases.ListZoneImage(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &ZoneImagesResponse{
		Page:       page,
		ZoneImages: zoneImages,
	}
	content = response
	return
}

func listZoneImageBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	zoneImage := &model.ZoneImage{
		ID: getIntQuery(r, "id"),
	}
	zoneImages, err := cases.ListZoneImageBySearch(page, zoneImage, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &ZoneImagesBySearchResponse{
		Page:       page,
		ZoneImages: zoneImages,
		Search:     zoneImage,
	}
	content = response
	return
}

func getZoneImageBot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	id := getIntVar(r, "botID")
	bot := &model.Bot{
		ID: id,
	}

	err = cases.GetZoneImageBot(bot, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &ZoneImageBotResponse{
		Bot: bot,
	}
	content = response
	return
}

func listZoneImageBot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	page := &model.Page{
		Limit: 100,
	}
	bots, err := cases.ListZoneImageBot(page, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &ZoneImageBotsResponse{
		Bots: bots,
	}
	content = response
	return
}

func editZoneImageBot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	bot := &model.Bot{}
	err = decodeBody(r, bot)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	bot.ID = getIntVar(r, "botID")

	err = cases.EditZoneImageBot(bot, user)
	if err != nil {
		return
	}
	response := &ZoneImageBotResponse{
		Bot: bot,
	}
	content = response
	return
}

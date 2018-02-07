package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListForum lists all forums accessible by provided user
func ListForum(page *model.Page, user *model.User) (forums []*model.Forum, err error) {
	err = validateOrderByForumField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("forum")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for forum")
		return
	}

	forums, err = reader.ListForum(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list forum")
		return
	}
	for i, forum := range forums {
		err = sanitizeForum(forum, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize forum element %d", i)
			return
		}
	}

	page.Total, err = reader.ListForumTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list forum toal count")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}

	return
}

//ListForumBySearch will request any forum matching the pattern of name
func ListForumBySearch(page *model.Page, forum *model.Forum, user *model.User) (forums []*model.Forum, err error) {

	err = validateOrderByForumField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareForum(forum, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre forum")
		return
	}

	err = validateForum(forum, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate forum")
		return
	}
	reader, err := getReader("forum")
	if err != nil {
		err = errors.Wrap(err, "failed to get forum reader")
		return
	}

	forums, err = reader.ListForumBySearch(page, forum)
	if err != nil {
		err = errors.Wrap(err, "failed to list forum by search")
		return
	}

	page.Total, err = reader.ListForumBySearchTotalCount(forum)
	if err != nil {
		err = errors.Wrap(err, "failed to get page total")
		return
	}

	for _, forum := range forums {
		err = sanitizeForum(forum, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize forum")
			return
		}
	}

	err = sanitizeForum(forum, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search forum")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateForum will create an forum using provided information
func CreateForum(forum *model.Forum, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list forum by search without guide+")
		return
	}
	err = prepareForum(forum, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare forum")
		return
	}

	err = validateForum(forum, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate forum")
		return
	}
	forum.ID = 0
	//forum.TimeCreation = time.Now().Unix()
	writer, err := getWriter("forum")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for forum")
		return
	}
	err = writer.CreateForum(forum)
	if err != nil {
		err = errors.Wrap(err, "failed to create forum")
		return
	}

	memWriter, err := getWriter("forum")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for forum")
		return
	}
	err = memWriter.CreateForum(forum)
	if err != nil {
		err = errors.Wrap(err, "failed to edit forum")
		return
	}

	err = sanitizeForum(forum, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize forum")
		return
	}
	return
}

//GetForum gets an forum by provided forumID
func GetForum(forum *model.Forum, user *model.User) (err error) {
	err = prepareForum(forum, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare forum")
		return
	}

	err = validateForum(forum, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate forum")
		return
	}

	reader, err := getReader("forum")
	if err != nil {
		err = errors.Wrap(err, "failed to get forum reader")
		return
	}

	err = reader.GetForum(forum)
	if err != nil {
		err = errors.Wrap(err, "failed to get forum")
		return
	}

	err = sanitizeForum(forum, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize forum")
		return
	}

	return
}

//EditForum edits an existing forum
func EditForum(forum *model.Forum, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list forum by search without guide+")
		return
	}
	err = prepareForum(forum, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare forum")
		return
	}

	err = validateForum(forum,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate forum")
		return
	}
	writer, err := getWriter("forum")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for forum")
		return
	}
	err = writer.EditForum(forum)
	if err != nil {
		err = errors.Wrap(err, "failed to edit forum")
		return
	}

	memWriter, err := getWriter("forum")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for forum")
		return
	}
	err = memWriter.EditForum(forum)
	if err != nil {
		err = errors.Wrap(err, "failed to edit forum")
		return
	}

	err = sanitizeForum(forum, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize forum")
		return
	}
	return
}

//DeleteForum deletes an forum by provided forumID
func DeleteForum(forum *model.Forum, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete forum without admin+")
		return
	}
	err = prepareForum(forum, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare forum")
		return
	}

	err = validateForum(forum, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate forum")
		return
	}
	writer, err := getWriter("forum")
	if err != nil {
		err = errors.Wrap(err, "failed to get forum writer")
		return
	}
	err = writer.DeleteForum(forum)
	if err != nil {
		err = errors.Wrap(err, "failed to delete forum")
		return
	}

	memWriter, err := getWriter("forum")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for forum")
		return
	}
	err = memWriter.DeleteForum(forum)
	if err != nil {
		err = errors.Wrap(err, "failed to delete forum")
		return
	}
	return
}

func prepareForum(forum *model.Forum, user *model.User) (err error) {
	if forum == nil {
		err = fmt.Errorf("empty forum")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateForum(forum *model.Forum, required []string, optional []string) (err error) {
	schema, err := newSchemaForum(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(forum))
	if err != nil {
		return
	}

	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}
	return
}

func validateOrderByForumField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "name"
	}

	validNames := []string{
		"id",
		"name",
	}

	possibleNames := ""
	for _, name := range validNames {
		if page.OrderBy == name {
			return
		}
		possibleNames += name + ", "
	}
	if len(possibleNames) > 0 {
		possibleNames = possibleNames[0 : len(possibleNames)-2]
	}
	err = &model.ErrValidation{
		Message: "orderBy is invalid. Possible fields: " + possibleNames,
		Reasons: map[string]string{
			"orderBy": "field is not valid",
		},
	}
	return
}

func sanitizeForum(forum *model.Forum, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}

	return
}

func newSchemaForum(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyForum(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyForum(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	jsRef := gojsonschema.NewGoLoader(s)
	schema, err = gojsonschema.NewSchema(jsRef)
	if err != nil {
		return
	}
	return
}

func getSchemaPropertyForum(field string) (prop model.Schema, err error) {
	switch field {

	case "ID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "name":
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}

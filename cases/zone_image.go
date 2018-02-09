package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadZoneImageFromFileToMemory is ran during initialization
func LoadZoneImageFromFileToMemory() (err error) {

	fr, err := file.New("config", "zoneImage.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("zoneImage-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize zoneImage-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("zoneImage-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize zoneImage-memory")
		return
	}

	fileReader, err := getReader("zoneImage-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneImage-file reader")
		return
	}

	memWriter, err := getWriter("zoneImage-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneImage-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListZoneImageTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list zoneImage count")
		return
	}
	page.Limit = page.Total

	zoneImages, err := fileReader.ListZoneImage(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list zoneImages")
		return
	}

	for _, zoneImage := range zoneImages {
		err = memWriter.CreateZoneImage(zoneImage)
		if err != nil {
			err = errors.Wrap(err, "failed to create zoneImage")
			return
		}
	}

	fmt.Printf("%d zoneImages, ", len(zoneImages))
	return
}

//ListZoneImage lists all zoneImages accessible by provided user
func ListZoneImage(page *model.Page, user *model.User) (zoneImages []*model.ZoneImage, err error) {
	err = validateOrderByZoneImageField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("zoneImage-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for zoneImage")
		return
	}

	page.Total, err = reader.ListZoneImageTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list zoneImage toal count")
		return
	}

	zoneImages, err = reader.ListZoneImage(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list zoneImage")
		return
	}
	for i, zoneImage := range zoneImages {
		err = sanitizeZoneImage(zoneImage, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize zoneImage element %d", i)
			return
		}
	}
	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}

	return
}

//ListZoneImageBot gets a bot by provided ID
func ListZoneImageBot(page *model.Page, user *model.User) (bots []*model.Bot, err error) {
	worker, err := getWorker("zoneImage")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneImage worker")
		return
	}

	bots, err = worker.ListBot(page)
	if err != nil {
		err = errors.Wrap(err, "failed to get bot")
		return
	}
	return
}

//ListZoneImageBySearch will request any zoneImage matching the pattern of name
func ListZoneImageBySearch(page *model.Page, zoneImage *model.ZoneImage, user *model.User) (zoneImages []*model.ZoneImage, err error) {

	err = validateOrderByZoneImageField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareZoneImage(zoneImage, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre zoneImage")
		return
	}

	err = validateZoneImage(zoneImage, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate zoneImage")
		return
	}
	reader, err := getReader("zoneImage-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneImage-memory reader")
		return
	}

	zoneImages, err = reader.ListZoneImageBySearch(page, zoneImage)
	if err != nil {
		err = errors.Wrap(err, "failed to list zoneImage by search")
		return
	}

	for _, zoneImage := range zoneImages {
		err = sanitizeZoneImage(zoneImage, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize zoneImage")
			return
		}
	}

	err = sanitizeZoneImage(zoneImage, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search zoneImage")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateZoneImage will create an zoneImage using provided information
func CreateZoneImage(zoneImage *model.ZoneImage, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list zoneImage by search without guide+")
		return
	}
	err = prepareZoneImage(zoneImage, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare zoneImage")
		return
	}

	err = validateZoneImage(zoneImage, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate zoneImage")
		return
	}
	zoneImage.ID = 0
	writer, err := getWriter("zoneImage-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for zoneImage")
		return
	}
	err = writer.CreateZoneImage(zoneImage)
	if err != nil {
		err = errors.Wrap(err, "failed to create zoneImage")
		return
	}

	fileWriter, err := getWriter("zoneImage-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneImage-file writer")
		return
	}
	err = fileWriter.CreateZoneImage(zoneImage)
	if err != nil {
		err = errors.Wrap(err, "failed to create zoneImage-file")
		return
	}
	err = sanitizeZoneImage(zoneImage, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize zoneImage")
		return
	}
	return
}

//GetZoneImage gets an zoneImage by provided zoneImageID
func GetZoneImage(zoneImage *model.ZoneImage, user *model.User) (err error) {
	err = prepareZoneImage(zoneImage, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare zoneImage")
		return
	}

	err = validateZoneImage(zoneImage, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate zoneImage")
		return
	}

	reader, err := getReader("zoneImage-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneImage-memory reader")
		return
	}

	err = reader.GetZoneImage(zoneImage)
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneImage")
		return
	}

	err = sanitizeZoneImage(zoneImage, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize zoneImage")
		return
	}

	return
}

//GetZoneImageBot gets a bot by provided ID
func GetZoneImageBot(bot *model.Bot, user *model.User) (err error) {
	worker, err := getWorker("zoneImage")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneImage worker")
		return
	}

	err = worker.GetBot(bot)
	if err != nil {
		err = errors.Wrap(err, "failed to get bot")
		return
	}
	return
}

//EditZoneImage edits an existing zoneImage
func EditZoneImage(zoneImage *model.ZoneImage, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list zoneImage by search without guide+")
		return
	}
	err = prepareZoneImage(zoneImage, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare zoneImage")
		return
	}

	err = validateZoneImage(zoneImage,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"male",
			"female",
			"neutral",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate zoneImage")
		return
	}
	writer, err := getWriter("zoneImage-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for zoneImage")
		return
	}
	err = writer.EditZoneImage(zoneImage)
	if err != nil {
		err = errors.Wrap(err, "failed to edit zoneImage")
		return
	}

	fileWriter, err := getWriter("zoneImage-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneImage-file writer")
		return
	}
	err = fileWriter.EditZoneImage(zoneImage)
	if err != nil {
		err = errors.Wrap(err, "failed to edit zoneImage-file")
		return
	}

	err = sanitizeZoneImage(zoneImage, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize zoneImage")
		return
	}
	return
}

//EditZoneImageBot edits an existing zoneImage
func EditZoneImageBot(bot *model.Bot, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't edit zoneImageBot by search without guide+")
		return
	}
	worker, err := getWorker("zoneImage")

	err = worker.EditBot(bot)
	if err != nil {
		err = errors.Wrap(err, "failed to edit zoneImageBot")
		return
	}
	return
}

//DeleteZoneImage deletes an zoneImage by provided zoneImageID
func DeleteZoneImage(zoneImage *model.ZoneImage, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete zoneImage without admin+")
		return
	}
	err = prepareZoneImage(zoneImage, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare zoneImage")
		return
	}

	err = validateZoneImage(zoneImage, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate zoneImage")
		return
	}
	writer, err := getWriter("zoneImage-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneImage-memory writer")
		return
	}
	err = writer.DeleteZoneImage(zoneImage)
	if err != nil {
		err = errors.Wrap(err, "failed to delete zoneImage")
		return
	}

	fileWriter, err := getWriter("zoneImage-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneImage-file writer")
		return
	}
	err = fileWriter.DeleteZoneImage(zoneImage)
	if err != nil {
		err = errors.Wrap(err, "failed to delete zoneImage-file")
		return
	}

	return
}

func prepareZoneImage(zoneImage *model.ZoneImage, user *model.User) (err error) {
	if zoneImage == nil {
		err = fmt.Errorf("empty zoneImage")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateZoneImage(zoneImage *model.ZoneImage, required []string, optional []string) (err error) {
	schema, err := newSchemaZoneImage(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(zoneImage))
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

func validateOrderByZoneImageField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "id"
	}

	validNames := []string{
		"id",
		"name",
		"bit",
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

func sanitizeZoneImage(zoneImage *model.ZoneImage, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}

	return
}

func newSchemaZoneImage(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyZoneImage(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyZoneImage(field); err != nil {
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

func getSchemaPropertyZoneImage(field string) (prop model.Schema, err error) {
	switch field {
	case "ID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "male":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "female":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "neutral":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "icon":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}

package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadSlotFromFileToMemory is ran during initialization
func LoadSlotFromFileToMemory() (err error) {

	fr, err := file.New("config", "slot.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("slot-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize slot-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("slot-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize slot-memory")
		return
	}

	fileReader, err := getReader("slot-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get slot-file reader")
		return
	}

	memWriter, err := getWriter("slot-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get slot-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListSlotTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list slot count")
		return
	}
	page.Limit = page.Total

	slots, err := fileReader.ListSlot(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list slots")
		return
	}

	for _, slot := range slots {
		err = memWriter.CreateSlot(slot)
		if err != nil {
			err = errors.Wrap(err, "failed to create slot")
			return
		}
	}

	fmt.Printf("%d slotes, ", len(slots))
	return
}

//ListSlot lists all slots accessible by provided user
func ListSlot(page *model.Page, user *model.User) (slots []*model.Slot, err error) {
	err = validateOrderBySlotField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("slot-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for slot")
		return
	}

	page.Total, err = reader.ListSlotTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list slot toal count")
		return
	}

	slots, err = reader.ListSlot(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list slot")
		return
	}
	for i, slot := range slots {
		err = sanitizeSlot(slot, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize slot element %d", i)
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

//ListSlotBySearch will request any slot matching the pattern of name
func ListSlotBySearch(page *model.Page, slot *model.Slot, user *model.User) (slots []*model.Slot, err error) {

	err = validateOrderBySlotField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSlot(slot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre slot")
		return
	}

	err = validateSlot(slot, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate slot")
		return
	}
	reader, err := getReader("slot-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get slot-memory reader")
		return
	}

	slots, err = reader.ListSlotBySearch(page, slot)
	if err != nil {
		err = errors.Wrap(err, "failed to list slot by search")
		return
	}

	for _, slot := range slots {
		err = sanitizeSlot(slot, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize slot")
			return
		}
	}

	err = sanitizeSlot(slot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search slot")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//ListSlotByBit lists all slots that match a bitmask
func ListSlotByBit(page *model.Page, slot *model.Slot, user *model.User) (slots []*model.Slot, err error) {
	err = validateOrderBySlotField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("slot-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for slot")
		return
	}

	page.Total, err = reader.ListSlotByBitTotalCount(slot)
	if err != nil {
		err = errors.Wrap(err, "failed to list slot toal count")
		return
	}

	slots, err = reader.ListSlotByBit(page, slot)
	if err != nil {
		err = errors.Wrap(err, "failed to list slot")
		return
	}

	for i, slot := range slots {
		err = sanitizeSlot(slot, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize slot element %d", i)
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

//CreateSlot will create an slot using provided information
func CreateSlot(slot *model.Slot, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list slot by search without guide+")
		return
	}
	err = prepareSlot(slot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare slot")
		return
	}

	err = validateSlot(slot, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate slot")
		return
	}
	slot.ID = 0
	writer, err := getWriter("slot-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for slot")
		return
	}
	err = writer.CreateSlot(slot)
	if err != nil {
		err = errors.Wrap(err, "failed to create slot")
		return
	}

	fileWriter, err := getWriter("slot-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get slot-file writer")
		return
	}
	err = fileWriter.CreateSlot(slot)
	if err != nil {
		err = errors.Wrap(err, "failed to create slot-file")
		return
	}
	err = sanitizeSlot(slot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize slot")
		return
	}
	return
}

//GetSlot gets an slot by provided slotID
func GetSlot(slot *model.Slot, user *model.User) (err error) {
	err = prepareSlot(slot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare slot")
		return
	}

	err = validateSlot(slot, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate slot")
		return
	}

	reader, err := getReader("slot-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get slot-memory reader")
		return
	}

	err = reader.GetSlot(slot)
	if err != nil {
		err = errors.Wrap(err, "failed to get slot")
		return
	}

	err = sanitizeSlot(slot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize slot")
		return
	}

	return
}

//EditSlot edits an existing slot
func EditSlot(slot *model.Slot, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list slot by search without guide+")
		return
	}
	err = prepareSlot(slot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare slot")
		return
	}

	err = validateSlot(slot,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"male",
			"female",
			"neutral",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate slot")
		return
	}
	writer, err := getWriter("slot-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for slot")
		return
	}
	err = writer.EditSlot(slot)
	if err != nil {
		err = errors.Wrap(err, "failed to edit slot")
		return
	}

	fileWriter, err := getWriter("slot-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get slot-file writer")
		return
	}
	err = fileWriter.EditSlot(slot)
	if err != nil {
		err = errors.Wrap(err, "failed to edit slot-file")
		return
	}

	err = sanitizeSlot(slot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize slot")
		return
	}
	return
}

//DeleteSlot deletes an slot by provided slotID
func DeleteSlot(slot *model.Slot, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete slot without admin+")
		return
	}
	err = prepareSlot(slot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare slot")
		return
	}

	err = validateSlot(slot, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate slot")
		return
	}
	writer, err := getWriter("slot-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get slot-memory writer")
		return
	}
	err = writer.DeleteSlot(slot)
	if err != nil {
		err = errors.Wrap(err, "failed to delete slot")
		return
	}

	fileWriter, err := getWriter("slot-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get slot-file writer")
		return
	}
	err = fileWriter.DeleteSlot(slot)
	if err != nil {
		err = errors.Wrap(err, "failed to delete slot-file")
		return
	}

	return
}

func prepareSlot(slot *model.Slot, user *model.User) (err error) {
	if slot == nil {
		err = fmt.Errorf("empty slot")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSlot(slot *model.Slot, required []string, optional []string) (err error) {
	schema, err := newSchemaSlot(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(slot))
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

func validateOrderBySlotField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "name"
	}

	validNames := []string{
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

func sanitizeSlot(slot *model.Slot, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaSlot(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySlot(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySlot(field); err != nil {
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

func getSchemaPropertySlot(field string) (prop model.Schema, err error) {
	switch field {
	case "ID":
		prop.Type = "integer"
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "shortName":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}

func getSlotListByBit(bit int64) string {
	page := &model.Page{
		Limit: 500,
	}
	user := &model.User{}
	slots, err := ListSlot(page, user)
	if err != nil {
		return ""
	}

	classes := ""
	for _, slot := range slots {
		if bit&slot.Bit == bit {
			classes += slot.ShortName + " "
		}
	}

	if len(classes) > 0 {
		classes = classes[0 : len(classes)-1]
	}
	if len(classes) == 0 {
		classes = "NONE"
	}
	return classes
}

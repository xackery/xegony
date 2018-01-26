package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//ItemRepository handles ItemRepository cases and is a gateway to storage
type ItemRepository struct {
	stor                   storage.Storage
	itemCategoryRepository *ItemCategoryRepository
}

//Initialize handles logic
func (c *ItemRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	c.itemCategoryRepository = &ItemCategoryRepository{}
	err = c.itemCategoryRepository.Initialize(stor)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize item category")
		return
	}

	return
}

//Get handles logic
func (c *ItemRepository) Get(item *model.Item, user *model.User) (err error) {

	err = c.stor.GetItem(item)
	if err != nil {
		err = errors.Wrap(err, "failed to get item")
		return
	}
	item.Category, err = c.itemCategoryRepository.GetByItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get item category")
		return
	}
	err = c.prepare(item)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare item")
		return
	}
	return
}

//SearchByName handles logic
func (c *ItemRepository) SearchByName(item *model.Item, user *model.User) (items []*model.Item, err error) {
	items, err = c.stor.SearchItemByName(item)
	if err != nil {
		err = errors.Wrap(err, "failed to get items")
		return
	}
	for _, item := range items {
		item.Category, err = c.itemCategoryRepository.GetByItem(item, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get item category")
		}
		err = c.prepare(item)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare item")
			return
		}
	}
	return
}

//SearchByAccount handles logic
func (c *ItemRepository) SearchByAccount(item *model.Item, account *model.Account, user *model.User) (items []*model.Item, err error) {
	items, err = c.stor.SearchItemByAccount(item, account)
	if err != nil {
		return
	}

	for _, item := range items {
		item.Category, err = c.itemCategoryRepository.GetByItem(item, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get item category")
		}
		err = c.prepare(item)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare item")
			return
		}
	}
	return
}

//Create handles logic
func (c *ItemRepository) Create(item *model.Item, user *model.User) (err error) {
	if item == nil {
		err = fmt.Errorf("Empty item")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	item.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(item))
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
	err = c.stor.CreateItem(item)
	if err != nil {
		return
	}

	item.Category, err = c.itemCategoryRepository.GetByItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get item category")
		return
	}

	err = c.prepare(item)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare item")
		return
	}
	return
}

//Edit handles logic
func (c *ItemRepository) Edit(item *model.Item, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(item))
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

	err = c.stor.EditItem(item)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *ItemRepository) Delete(item *model.Item, user *model.User) (err error) {
	err = c.stor.DeleteItem(item)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *ItemRepository) List(pageSize int64, pageNumber int64, user *model.User) (items []*model.Item, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	items, err = c.stor.ListItem(pageSize, pageNumber)
	if err != nil {
		return
	}
	for _, item := range items {
		item.Category, err = c.itemCategoryRepository.GetByItem(item, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get item category")
			return
		}

		err = c.prepare(item)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare item")
			return
		}
	}
	return
}

//ListCount handles logic
func (c *ItemRepository) ListCount(user *model.User) (count int64, err error) {

	count, err = c.stor.ListItemCount()
	if err != nil {
		return
	}
	return
}

//ListByCharacter handles logic
func (c *ItemRepository) ListByCharacter(character *model.Character, user *model.User) (items []*model.Item, err error) {
	items, err = c.stor.ListItemByCharacter(character)
	if err != nil {
		return
	}
	for _, item := range items {
		item.Category, err = c.itemCategoryRepository.GetByItem(item, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get item category")
			return
		}

		err = c.prepare(item)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare item")
			return
		}
	}
	return
}

//ListBySpell handles logic
func (c *ItemRepository) ListBySpell(spell *model.Spell, user *model.User) (items []*model.Item, err error) {
	items, err = c.stor.ListItemBySpell(spell)
	if err != nil {
		return
	}
	for _, item := range items {
		item.Category, err = c.itemCategoryRepository.GetByItem(item, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get item category")
			return
		}

		err = c.prepare(item)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare item")
			return
		}
	}
	return
}

//GetByItemCategory handles logic
func (c *ItemRepository) GetByItemCategory(itemCategory *model.ItemCategory, user *model.User) (items []*model.Item, err error) {
	items, err = c.stor.ListItemByItemCategory(itemCategory)
	if err != nil {
		return
	}
	for _, item := range items {
		item.Category, err = c.itemCategoryRepository.GetByItem(item, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get item category")
			return
		}

		err = c.prepare(item)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare item")
			return
		}
	}
	return
}

//GetByZone handles logic
func (c *ItemRepository) GetByZone(zone *model.Zone, user *model.User) (items []*model.Item, err error) {
	items, err = c.stor.ListItemByZone(zone)
	if err != nil {
		return
	}
	for _, item := range items {
		item.Category, err = c.itemCategoryRepository.GetByItem(item, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get item category")
			return
		}

		err = c.prepare(item)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare item")
			return
		}
	}
	return
}

func (c *ItemRepository) prepare(item *model.Item) (err error) {

	iColor := fmt.Sprintf("%d", item.Color)
	if len(iColor) > 8 {
		color := ""
		//color := "style=\"color: rgba("
		pos := 0
		if len(iColor) > 9 {
			pos = 2
		}
		alpha := iColor[pos : pos+2] //alpha
		pos += 2

		color += iColor[pos:pos+2] + "," //rr
		pos += 2
		color += iColor[pos:pos+2] + "," //gg
		pos += 2
		color += iColor[pos:pos+2] + "," //bb
		color += alpha                   //add alpha to end
		//color += "\");"
		item.StyleColor = color
	}

	return
}

func (c *ItemRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
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

func (c *ItemRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "zoneID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}

package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//FishingRepository handles FishingRepository cases and is a gateway to storage
type FishingRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *FishingRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *FishingRepository) Get(fishing *model.Fishing, user *model.User) (err error) {

	err = c.stor.GetFishing(fishing)
	if err != nil {
		err = errors.Wrap(err, "failed to get fishing")
		return
	}
	err = c.prepare(fishing, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare fishing")
		return
	}
	return
}

//Create handles logic
func (c *FishingRepository) Create(fishing *model.Fishing, user *model.User) (err error) {
	if fishing == nil {
		err = fmt.Errorf("Empty fishing")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	fishing.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(fishing))
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
	err = c.stor.CreateFishing(fishing)
	if err != nil {
		return
	}
	err = c.prepare(fishing, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare fishing")
		return
	}
	return
}

//Edit handles logic
func (c *FishingRepository) Edit(fishing *model.Fishing, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(fishing))
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

	err = c.stor.EditFishing(fishing)
	if err != nil {
		return
	}
	err = c.prepare(fishing, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare fishing")
		return
	}
	return
}

//Delete handles logic
func (c *FishingRepository) Delete(fishing *model.Fishing, user *model.User) (err error) {
	err = c.stor.DeleteFishing(fishing)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *FishingRepository) List(pageSize int64, pageNumber int64, user *model.User) (fishings []*model.Fishing, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	fishings, err = c.stor.ListFishing(pageSize, pageNumber)
	if err != nil {
		return
	}
	for _, fishing := range fishings {
		err = c.prepare(fishing, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare fishing")
			return
		}
	}
	return
}

//ListCount handles logic
func (c *FishingRepository) ListCount(user *model.User) (count int64, err error) {

	count, err = c.stor.ListFishingCount()
	if err != nil {
		return
	}
	return
}

//GetByZone handles logic
func (c *FishingRepository) GetByZone(zone *model.Zone, user *model.User) (fishings []*model.Fishing, err error) {
	fishings, err = c.stor.ListFishingByZone(zone)
	if err != nil {
		return
	}
	for _, fishing := range fishings {
		err = c.prepare(fishing, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare fishing")
			return
		}
	}
	return
}

//GetByNpc handles logic
func (c *FishingRepository) GetByNpc(npc *model.Npc, user *model.User) (fishings []*model.Fishing, err error) {
	fishings, err = c.stor.ListFishingByNpc(npc)
	if err != nil {
		return
	}
	for _, fishing := range fishings {
		err = c.prepare(fishing, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare fishing")
			return
		}
	}
	return
}

//GetByItem handles logic
func (c *FishingRepository) GetByItem(item *model.Item, user *model.User) (fishings []*model.Fishing, err error) {
	fishings, err = c.stor.ListFishingByItem(item)
	if err != nil {
		return
	}
	for _, fishing := range fishings {
		err = c.prepare(fishing, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare fishing")
			return
		}
	}
	return
}

func (c *FishingRepository) prepare(fishing *model.Fishing, user *model.User) (err error) {

	return
}

func (c *FishingRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *FishingRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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

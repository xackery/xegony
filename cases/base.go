package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type BaseRepository struct {
	stor storage.Storage
}

func (g *BaseRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *BaseRepository) Get(class int64, level int64) (base *model.Base, err error) {
	base, err = g.stor.GetBase(class, level)
	return
}

func (g *BaseRepository) Create(base *model.Base) (err error) {
	if base == nil {
		err = fmt.Errorf("Empty base")
		return
	}
	schema, err := base.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(base))
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
	err = g.stor.CreateBase(base)
	if err != nil {
		return
	}
	return
}

func (g *BaseRepository) Edit(class int64, level int64, base *model.Base) (err error) {
	schema, err := base.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(base))
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

	err = g.stor.EditBase(class, level, base)
	if err != nil {
		return
	}
	return
}

func (g *BaseRepository) Delete(class int64, level int64) (err error) {
	err = g.stor.DeleteBase(class, level)
	if err != nil {
		return
	}
	return
}

func (g *BaseRepository) List() (bases []*model.Base, err error) {
	bases, err = g.stor.ListBase()
	if err != nil {
		return
	}
	return
}

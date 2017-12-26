package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type ZoneRepository struct {
	stor storage.Storage
}

func (g *ZoneRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *ZoneRepository) Get(zoneId int64) (zone *model.Zone, err error) {
	if zoneId == 0 {
		err = fmt.Errorf("Invalid Zone ID")
		return
	}
	zone, err = g.stor.GetZone(zoneId)
	return
}

func (g *ZoneRepository) Create(zone *model.Zone) (err error) {
	if zone == nil {
		err = fmt.Errorf("Empty zone")
		return
	}
	schema, err := zone.NewSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}
	if zone.ZoneIdNumber < 1 {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		vErr.Reasons["accountId"] = "Account ID must be greater than 0"
		err = vErr
		return
	}
	zone.Id = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(zone))
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
	err = g.stor.CreateZone(zone)
	if err != nil {
		return
	}
	return
}

func (g *ZoneRepository) Edit(zoneId int64, zone *model.Zone) (err error) {
	schema, err := zone.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(zone))
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

	err = g.stor.EditZone(zoneId, zone)
	if err != nil {
		return
	}
	return
}

func (g *ZoneRepository) Delete(zoneId int64) (err error) {
	err = g.stor.DeleteZone(zoneId)
	if err != nil {
		return
	}
	return
}

func (g *ZoneRepository) List() (zones []*model.Zone, err error) {
	zones, err = g.stor.ListZone()
	if err != nil {
		return
	}
	return
}

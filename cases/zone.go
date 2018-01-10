package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

var (
	hotzone = func(z1, z2 *model.Zone) bool {
		return z1.Modifier() > z2.Modifier()
	}

	levels = func(z1, z2 *model.Zone) bool {
		if z1.Levels&1 == 1 && z2.Levels&1 == 0 {
			return true
		}
		if z1.Levels&2 == 2 && z2.Levels&2 == 0 {
			return true
		}
		if z1.Levels&4 == 4 && z2.Levels&4 == 0 {
			return true
		}
		if z1.Levels&8 == 8 && z2.Levels&8 == 0 {
			return true
		}
		if z1.Levels&16 == 16 && z2.Levels&16 == 0 {
			return true
		}
		if z1.Levels&32 == 32 && z2.Levels&32 == 0 {
			return true
		}
		if z1.Levels&64 == 64 && z2.Levels&64 == 0 {
			return true
		}
		if z1.Levels&128 == 128 && z2.Levels&128 == 0 {
			return true
		}
		if z1.Levels&256 == 256 && z2.Levels&256 == 0 {
			return true
		}
		if z1.Levels&512 == 512 && z2.Levels&512 == 0 {
			return true
		}
		if z1.Levels&1024 == 1024 && z2.Levels&1024 == 0 {
			return true
		}
		if z1.Levels&2048 == 2048 && z2.Levels&2048 == 0 {
			return true
		}
		if z1.Levels&4096 == 4096 && z2.Levels&4096 == 0 {
			return true
		}
		if z1.Levels&8192 == 8192 && z2.Levels&8192 == 0 {
			return true
		}
		if z1.Levels&16384 == 16384 && z2.Levels&16384 == 0 {
			return true
		}
		if z1.Levels&32768 == 32768 && z2.Levels&32768 == 0 {
			return true
		}
		if z1.Levels&65536 == 65536 && z2.Levels&65536 == 0 {
			return true
		}
		return false
	}

	expansions = func(z1, z2 *model.Zone) bool {
		return z1.Expansion > z2.Expansion
	}
)

//ZoneRepository handles ZoneRepository cases and is a gateway to storage
type ZoneRepository struct {
	stor              storage.Storage
	zoneCache         map[int64]*model.Zone
	isZoneCacheLoaded bool
}

//Initialize handler
func (c *ZoneRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}

	c.stor = stor
	c.isZoneCacheLoaded = false
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

func (c *ZoneRepository) rebuildCache() (err error) {
	if c.isZoneCacheLoaded {
		return
	}
	c.isZoneCacheLoaded = true
	c.zoneCache = make(map[int64]*model.Zone)
	zones, err := c.list()
	if err != nil {
		return
	}

	zoneLevelRepo := &ZoneLevelRepository{}
	if err = zoneLevelRepo.Initialize(c.stor); err != nil {
		return
	}

	zoneLevels, err := zoneLevelRepo.List()
	if err != nil {
		return
	}

	for _, zone := range zones {
		for _, zoneLevel := range zoneLevels {
			if zoneLevel.ZoneID == zone.ZoneIDNumber {
				zone.Levels = zoneLevel.Levels
				break

			}
		}
		c.zoneCache[zone.ZoneIDNumber] = zone
	}
	fmt.Println("Rebuilt Zone Cache")
	return
}

//Get handler
func (c *ZoneRepository) Get(zoneID int64) (zone *model.Zone, err error) {
	if zoneID == 0 {
		err = fmt.Errorf("Invalid Zone ID")
		return
	}
	zone = c.zoneCache[zoneID]
	//zone, err = c.stor.GetZone(zoneID)
	return
}

//GetByShortname gets a zone by it's short name
func (c *ZoneRepository) GetByShortname(zoneShortname string) (zone *model.Zone, err error) {
	for _, zoneC := range c.zoneCache {
		if zoneC.ShortName.String == zoneShortname {
			zone = zoneC
			return
		}
	}
	return
}

//Create handler
func (c *ZoneRepository) Create(zone *model.Zone) (err error) {
	if zone == nil {
		err = fmt.Errorf("Empty zone")
		return
	}
	schema, err := c.newSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}
	if zone.ZoneIDNumber < 1 {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		vErr.Reasons["accountID"] = "Account ID must be greater than 0"
		err = vErr
		return
	}
	zone.ID = 0 //strip ID
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
	err = c.stor.CreateZone(zone)
	if err != nil {
		return
	}
	c.isZoneCacheLoaded = false
	c.rebuildCache()
	return
}

//Edit handler
func (c *ZoneRepository) Edit(zoneID int64, zone *model.Zone) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
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

	if err = c.stor.EditZone(zoneID, zone); err != nil {
		return
	}
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

//Delete handler
func (c *ZoneRepository) Delete(zoneID int64) (err error) {
	err = c.stor.DeleteZone(zoneID)
	if err != nil {
		return
	}
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

func (c *ZoneRepository) list() (zones []*model.Zone, err error) {
	if zones, err = c.stor.ListZone(); err != nil {
		return
	}
	return
}

//List handler
func (c *ZoneRepository) List() (zones []*model.Zone, err error) {
	for _, zone := range c.zoneCache {
		zones = append(zones, zone)
	}
	return
}

//ListByHotzone handler
func (c *ZoneRepository) ListByHotzone() (zones []model.Zone, err error) {

	for _, zonePtr := range c.zoneCache {
		zone := *zonePtr
		zones = append(zones, zone)
	}
	model.ZoneBy(hotzone).Sort(zones)
	return
}

func (c *ZoneRepository) prepare(zone *model.Zone) (err error) {

	return
}

//newSchema handler
func (c *ZoneRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

//getSchemaProperty handler
func (c *ZoneRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "shortName":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}

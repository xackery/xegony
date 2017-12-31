package model

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

//Task is an everquest entry, grouping the Activities entries
type Task struct {
	Id           int64  `json:"id" db:"id"`                     //`id` int(11) unsigned NOT NULL DEFAULT '0',
	Duration     int64  `json:"duration" db:"duration"`         //`duration` int(11) unsigned NOT NULL DEFAULT '0',
	Title        string `json:"title" db:"title"`               //`title` varchar(100) NOT NULL DEFAULT '',
	Description  string `json:"description" db:"description"`   //`description` text NOT NULL,
	Reward       string `json:"reward" db:"reward"`             //`reward` varchar(64) NOT NULL DEFAULT '',
	Rewardid     int64  `json:"rewardid" db:"rewardid"`         //`rewardid` int(11) unsigned NOT NULL DEFAULT '0',
	Cashreward   int64  `json:"cashreward" db:"cashreward"`     //`cashreward` int(11) unsigned NOT NULL DEFAULT '0',
	Xpreward     int64  `json:"xpreward" db:"xpreward"`         //`xpreward` int(10) NOT NULL DEFAULT '0',
	Rewardmethod int64  `json:"rewardmethod" db:"rewardmethod"` //`rewardmethod` tinyint(3) unsigned NOT NULL DEFAULT '2',
	Startzone    int64  `json:"startzone" db:"startzone"`       //`startzone` int(11) NOT NULL DEFAULT '0',
	Minlevel     int64  `json:"minlevel" db:"minlevel"`         //`minlevel` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Maxlevel     int64  `json:"maxlevel" db:"maxlevel"`         //`maxlevel` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Repeatable   int64  `json:"repeatable" db:"repeatable"`     //`repeatable` tinyint(1) unsigned NOT NULL DEFAULT '1',
}

func (c *Task) CashRewardName() string {
	return CashName(c.Cashreward)
}

func (c *Task) RewardName() string {
	switch c.Rewardmethod {
	case 0:
		return c.Reward
	case 1:
		return fmt.Sprintf("GoalList %d", c.Rewardid)
	case 2:
		return "Perl"
	}
	return "Unknown"
}

func (c *Task) RewardMethodName() string {
	switch c.Rewardmethod {
	case 0:
		return "Item"
	case 1:
		return "Items"
	case 2:
		return "Perl"
	}
	return "Unknown"
}

func (c *Task) RepeatableName() string {
	switch c.Repeatable {
	case 0:
		return "No"
	}
	return "Yes"
}

func (c *Task) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]Schema)
	var field string
	var prop Schema
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

func (c *Task) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}

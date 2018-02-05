package file

import (
	"github.com/xackery/xegony/model"
)

func loadSpellTravelTypeDefault() model.SpellTravelTypes {
	return model.SpellTravelTypes{
		{
			ID:   0,
			Name: "None",
		},
		{
			ID:   3,
			Name: "Throw Model",
		},
		{
			ID:   16,
			Name: "Unknown",
		},
	}
}

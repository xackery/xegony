package file

import (
	"github.com/xackery/xegony/model"
)

func loadSpellAnimationTypeDefault() model.SpellAnimationTypes {
	return model.SpellAnimationTypes{
		{
			ID:        1,
			ShortName: "F",
			Name:      `Fire and Ice`,
		},
		{
			ID:        2,
			ShortName: "P",
			Name:      `Poison/Disease`,
		},
		{
			ID:        3,
			ShortName: "E",
			Name:      `Electricity`,
		},
		{
			ID:        4,
			ShortName: "S",
			Name:      `Sparkles`,
		},
		{
			ID:        5,
			ShortName: "L",
			Name:      `Light`,
		},
		{
			ID:        6,
			ShortName: "D",
			Name:      `Darkness`,
		},
		{
			ID:        7,
			ShortName: "N",
			Name:      `Nature`,
		},
		{
			ID:        8,
			ShortName: "U",
			Name:      `Blue`,
		},
		{
			ID:        9,
			ShortName: "T",
			Name:      `Teleportation`,
		},
		{
			ID:        10,
			ShortName: "M",
			Name:      `Music`,
		},
		{
			ID:        11,
			ShortName: "B",
			Name:      `Dragon Breath`,
		},
		{
			ID:        12,
			ShortName: "X",
			Name:      `Sound Effects`,
		},
		{
			ID:        13,
			ShortName: "I",
			Name:      `Shielding`,
		},
		{
			ID:        14,
			ShortName: "G",
			Name:      `Ground/Feet`,
		},
		{
			ID:        15,
			ShortName: "H",
			Name:      `Hands`,
		},
		{
			ID:        16,
			ShortName: "C",
			Name:      `Caster Only`,
		},
		{
			ID:        17,
			ShortName: "W",
			Name:      `Fireworks`,
		},
		{
			ID:        18,
			ShortName: "R",
			Name:      `Permanent/Glitched`,
		},
		{
			ID:        19,
			ShortName: "O",
			Name:      `Other/Special`,
		},
		{
			ID:        20,
			ShortName: "-",
			Name:      `Unused`,
		},
	}
}

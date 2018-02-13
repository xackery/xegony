package file

import (
	"github.com/xackery/xegony/model"
)

func loadSlotDefault() model.Slots {
	return model.Slots{
		{
			ID:        0,
			Bit:       0,
			Name:      "Unknown",
			ShortName: "UNK",
			Icon:      "xa-shield",
		},
		{
			ID:        1,
			Bit:       1 << 0,
			Name:      "Charm",
			ShortName: "CHARM",
			Icon:      "xa-shield",
		},
		{
			ID:        2,
			Bit:       1 << 2,
			Name:      "Head",
			ShortName: "HEAD",
			Icon:      "xa-shield",
		},
		{
			ID:        3,
			Bit:       1 << 3,
			Name:      "Face",
			ShortName: "FACE",
			Icon:      "xa-shield",
		},
		{
			ID:        4,
			Bit:       1<<1 | 1<<4,
			Name:      "Ears",
			ShortName: "EARS",
			Icon:      "xa-shield",
		},
		{
			ID:        5,
			Bit:       1 << 5,
			Name:      "Neck",
			ShortName: "NECK",
			Icon:      "xa-shield",
		},
		{
			ID:        6,
			Bit:       1 << 6,
			Name:      "Shoulder",
			ShortName: "SHOULDER",
			Icon:      "xa-shield",
		},
		{
			ID:        7,
			Bit:       1 << 7,
			Name:      "Arms",
			ShortName: "ARMS",
			Icon:      "xa-shield",
		},
		{
			ID:        8,
			Bit:       1 << 8,
			Name:      "Back",
			ShortName: "BACK",
			Icon:      "xa-shield",
		},
		{
			ID:        9,
			Bit:       1<<9 | 1<<10,
			Name:      "Bracers",
			ShortName: "BRACERS ",
			Icon:      "xa-shield",
		},
		{
			ID:        10,
			Bit:       1 << 11,
			Name:      "Range",
			ShortName: "RANGE",
			Icon:      "xa-shield",
		},
		{
			ID:        11,
			Bit:       1 << 12,
			Name:      "Hands",
			ShortName: "HANDS",
			Icon:      "xa-shield",
		},
		{
			ID:        12,
			Bit:       1 << 13,
			Name:      "Primary",
			ShortName: "PRIMARY",
			Icon:      "xa-shield",
		},
		{
			ID:        13,
			Bit:       1 << 14,
			Name:      "Second",
			ShortName: "SECONDARY",
			Icon:      "xa-shield",
		},
		{
			ID:        14,
			Bit:       1<<15 | 1<<16,
			Name:      "Rings",
			ShortName: "RINGS",
			Icon:      "xa-shield",
		},
		{
			ID:        15,
			Bit:       1 << 17,
			Name:      "Chest",
			ShortName: "CHEST",
			Icon:      "xa-shield",
		},
		{
			ID:        16,
			Bit:       1 << 18,
			Name:      "Legs",
			ShortName: "LEGS",
			Icon:      "xa-shield",
		},
		{
			ID:        17,
			Bit:       1 << 19,
			Name:      "Feet",
			ShortName: "FEET",
			Icon:      "xa-shield",
		},
		{
			ID:        18,
			Bit:       1 << 20,
			Name:      "Waist",
			ShortName: "WAIST",
			Icon:      "xa-shield",
		},
		{
			ID:        19,
			Bit:       1 << 21,
			Name:      "Ammo",
			ShortName: "AMMO",
			Icon:      "xa-shield",
		},
		{
			ID:        20,
			Bit:       1 << 22,
			Name:      "Power Source",
			ShortName: "POWERSOURCE",
			Icon:      "xa-shield",
		},
	}
}

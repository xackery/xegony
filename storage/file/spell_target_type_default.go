package file

import (
	"github.com/xackery/xegony/model"
)

func loadSpellTargetTypeDefault() model.SpellTargetTypes {
	return model.SpellTargetTypes{
		{
			ID:   0,
			Name: "Special",
		},
		{
			ID:   1,
			Name: "Line of Sight",
		},
		{
			ID:   2,
			Name: "Old PBAE (Unused)",
		},
		{
			ID:   3,
			Name: "Own Group",
		},
		{
			ID:   4,
			Name: "PBAE",
		},
		{
			ID:   5,
			Name: "Single",
		},
		{
			ID:   6,
			Name: "Self Only",
		},
		{
			ID:   8,
			Name: "Targeted AE",
		},
		{
			ID:   9,
			Name: "Animals Only",
		},
		{
			ID:   10,
			Name: "Undead Only",
		},
		{
			ID:   11,
			Name: "Summoned",
		},
		{
			ID:   13,
			Name: "Life Tap",
		},
		{
			ID:   14,
			Name: "Own Pet",
		},
		{
			ID:   15,
			Name: "Corpse",
		},
		{
			ID:   16,
			Name: "Plants",
		},
		{
			ID:   17,
			Name: "Special Velious Giants",
		},
		{
			ID:   18,
			Name: "Special Velious Dragons",
		},
		{
			ID:   20,
			Name: "Targeted AE Life Tap",
		},
		{
			ID:   24,
			Name: "AE Undead Only",
		},
		{
			ID:   25,
			Name: "AE Summoned Only",
		},
		{
			ID:   32,
			Name: "AE HateList/Casters Only",
		},
		{
			ID:   33,
			Name: "NPC's Hate List",
		},
		{
			ID:   34,
			Name: "Lost Dungeon Object",
		},
		{
			ID:   35,
			Name: "Muramite",
		},
		{
			ID:   36,
			Name: "AE PCs Only",
		},
		{
			ID:   37,
			Name: "AE NPCs Only",
		},
		{
			ID:   38,
			Name: "Any Summoned Pet",
		},
		{
			ID:   39,
			Name: "Group (No Pets)",
		},
		{
			ID:   40,
			Name: "Raid AE",
		},
		{
			ID:   41,
			Name: "Group (Targetable)",
		},
		{
			ID:   42,
			Name: "Directional Cone",
		},
		{
			ID:   43,
			Name: "Group With Pets",
		},
		{
			ID:   44,
			Name: "Beam",
		},
		{
			ID:   45,
			Name: "Ring",
		},
		{
			ID:   46,
			Name: "Target's Target",
		},
		{
			ID:   47,
			Name: "Targeted Pet's Master",
		},
		{
			ID:   50,
			Name: "Targeted AE (No Pets)",
		},
	}
}

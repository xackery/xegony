package file

import (
	"github.com/xackery/xegony/model"
)

func loadSpellDurationFormulaDefault() model.SpellDurationFormulas {
	return model.SpellDurationFormulas{
		{
			ID:   0,
			Name: "Instant",
		},
		{
			ID:   1,
			Name: "Level / 2 Ticks, Up to Max",
		},
		{
			ID:   2,
			Name: "Level * 3/5 Ticks, Up to Max",
		},
		{
			ID:   3,
			Name: "Level * 3 Minutes, Up to Max Ticks",
		},
		{
			ID:   4,
			Name: "Max Ticks, or 5min if Max <= 0",
		},
		{
			ID:   5,
			Name: "Max Ticks (enforced to 1-3)",
		},
		{
			ID:   6,
			Name: "Level / 2 Ticks, Up to Max (Unused in Live)",
		},
		{
			ID:   7,
			Name: "If Max == 0, Level Ticks, otherwise Max Ticks",
		},
		{
			ID:   8,
			Name: "Level + 10 Ticks, Up to Max",
		},
		{
			ID:   9,
			Name: "Level * 2 + 10 Ticks, Up to Max",
		},
		{
			ID:   10,
			Name: "Level * 3 + 10 Ticks, Up to Max",
		},
		{
			ID:   11,
			Name: "(Level + 3) * 3 Minutes, Up to Max Ticks",
		},
		{
			ID:   12,
			Name: "Max Ticks",
		},
		{
			ID:   13,
			Name: "Max Ticks",
		},
		{
			ID:   14,
			Name: "Max Ticks",
		},
		{
			ID:   15,
			Name: "Max Ticks",
		},
		{
			ID:   50,
			Name: "Permanent until Effect Is Otherwise Canceled",
		},
		{
			ID:   51,
			Name: "Permanent as Long as Target within Range of Aura",
		},
		{
			ID:   3600,
			Name: "If Max == 0, 6 hours, otherwise Max Ticks",
		},
	}
}

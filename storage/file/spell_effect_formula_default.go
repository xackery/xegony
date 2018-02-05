package file

import (
	"github.com/xackery/xegony/model"
)

func loadSpellEffectFormulaDefault() model.SpellEffectFormulas {
	return model.SpellEffectFormulas{
		{
			ID:   0,
			Name: "Base",
		},
		{
			ID:   1,
			Name: "(1-99) Base + Level * Formula",
		},
		{
			ID:   60,
			Name: "Base / 100",
		},
		{
			ID:   70,
			Name: "Base / 100",
		},
		{
			ID:   100,
			Name: "Base",
		},
		{
			ID:   101,
			Name: "(Level + Base) / 2",
		},
		{
			ID:   102,
			Name: "(Level + Base)",
		},
		{
			ID:   103,
			Name: "(Level + Base) * 2",
		},
		{
			ID:   104,
			Name: "(Level + Base) * 3",
		},
		{
			ID:   105,
			Name: "(Level + Base) * 4",
		},
		{
			ID:   106,
			Name: "(Level + Base) * 106",
		},
		{
			ID:   107,
			Name: "(Level / 2) + Base",
		},
		{
			ID:   108,
			Name: "(Level / 3) + Base",
		},
		{
			ID:   109,
			Name: "(Level / 4) + Base",
		},
		{
			ID:   110,
			Name: "(Level / 6) + Base",
		},
		{
			ID:   111,
			Name: "[(Level - 16) * 6] + Base",
		},
		{
			ID:   112,
			Name: "[(Level - 24) * 8] + Base",
		},
		{
			ID:   113,
			Name: "[(Level - 34) * 10] + Base",
		},
		{
			ID:   114,
			Name: "[(Level - 44) * 15] + Base",
		},
		{
			ID:   115,
			Name: "[(Level - 15) * 7] + Base",
		},
		{
			ID:   116,
			Name: "[(Level - 24) * 10] + Base",
		},
		{
			ID:   117,
			Name: "[(Level - 34) * 13] + Base",
		},
		{
			ID:   118,
			Name: "[(Level - 44) * 20] + Base",
		},
		{
			ID:   119,
			Name: "(Level / 8) + Base",
		},
		{
			ID:   121,
			Name: "(Level / 3) + Base",
		},
		{
			ID:   122,
			Name: "Base, counting down 12 per tick",
		},
		{
			ID:   123,
			Name: "Random from Base - Max",
		},
		{
			ID:   124,
			Name: "[(Level - 50) * 1] + Base",
		},
		{
			ID:   125,
			Name: "[(Level - 50) * 2] + Base",
		},
		{
			ID:   126,
			Name: "[(Level - 50) * 3] + Base",
		},
		{
			ID:   127,
			Name: "[(Level - 50) * 4] + Base",
		},
		{
			ID:   128,
			Name: "[(Level - 50) * 5] + Base",
		},
		{
			ID:   129,
			Name: "[(Level - 50) * 10] + Base",
		},
		{
			ID:   130,
			Name: "[(Level - 50) * 15] + Base",
		},
		{
			ID:   131,
			Name: "[(Level - 50) * 20] + Base",
		},
		{
			ID:   132,
			Name: "[(Level - 50) * 25] + Base",
		},
		{
			ID:   137,
			Name: "Base",
		},
		{
			ID:   138,
			Name: "Random from 0 - Base",
		},
		{
			ID:   139,
			Name: "[(Level - 30) / 2) + Base",
		},
		{
			ID:   140,
			Name: "(Level - 30) + Base",
		},
		{
			ID:   141,
			Name: "[(Level - 30) * 3/2) + Base",
		},
		{
			ID:   142,
			Name: "(Level - 30) + Base",
		},
		{
			ID:   143,
			Name: "(Level * 3/4) + Base",
		},
		{
			ID:   201,
			Name: "Max",
		},
		{
			ID:   203,
			Name: "Max",
		},
		{
			ID:   1001,
			Name: "(1001-1999) Base - 12 * (Formula - 1000) Per Tick",
		},
		{
			ID:   2001,
			Name: "(2001-2650) Base * Level * (Formula - 2000) Per Tick",
		},
	}
}

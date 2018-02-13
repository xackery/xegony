package file

import (
	"github.com/xackery/xegony/model"
)

func loadSizeDefault() model.Sizes {
	return model.Sizes{
		{
			ID:   0,
			Name: "TINY",
			Icon: "xa-shield",
		},
		{
			ID:   1,
			Name: "SMALL",
			Icon: "xa-shield",
		},
		{
			ID:   2,
			Name: "MEDIUM",
			Icon: "xa-shield",
		},
		{
			ID:   3,
			Name: "LARGE",
			Icon: "xa-shield",
		},
		{
			ID:   4,
			Name: "GIANT",
			Icon: "xa-shield",
		},
		{
			ID:   5,
			Name: "GIGANTIC",
			Icon: "xa-shield",
		},
	}
}

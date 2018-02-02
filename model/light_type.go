package model

import ()

// LightType is used to determine the type of light emission an item or spell should cause.
// swagger:model
type LightType struct {
	ID          int64  `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

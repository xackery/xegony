package model

import ()

// Configs is an array of Config
type Configs []*Config

// Config holds settings and configuration options
// swagger:model
type Config struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

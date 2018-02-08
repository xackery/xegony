package model

//Classs is an array of class
// swagger:model
type Classs []*Class

//Class represents classes in EQ
// swagger:model
type Class struct {
	ID        int64  `json:"ID" yaml:"ID"`
	Bit       int64  `json:"bit" yaml:"bit"`
	ShortName string `json:"shortName" yaml:"shortName"`
	Name      string `json:"name" yaml:"name"`
	Icon      string `json:"icon" yaml:"icon"`
}

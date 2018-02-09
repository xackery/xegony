package model

//ZoneImages is an array of zone map
type ZoneImages []*ZoneImage

//ZoneImage is details about the zone map processed data
type ZoneImage struct {
	ID          int64   `json:"ID,omitempty" yaml:"ID"`
	AspectRatio float64 `json:"aspectRatio,omitempty" yaml:"aspectratio"`
	XOffset     float64 `json:"xOffset,omitempty" yaml:"xOffset"`
	YOffset     float64 `json:"yOffset,omitempty" yaml:"yOffset"`
}

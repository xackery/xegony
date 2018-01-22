package model

import ()

//Schema represents http://json-schema.org/latest/json-schema-validation.html#anchor111
// swagger:response
type Schema struct {
	//Schema      string   `json:"$schema,omitempty"`
	Type              string            `json:"string,omitempty"` //integer, string, object
	Properties        map[string]Schema `json:"properties,omitempty"`
	PatternProperties map[string]Schema `json:"patternproperties,omitempty"` //"[0-9]": {}
	Items             map[string]Schema `json:"items,omitempty"`
	Required          []string          `json:"required,omitempty"`
	Description       string            `json:"description,omitempty"`
	Title             string            `json:"title,omitempty"`    //Description of property
	Minimum           int64             `json:"minimum,omitempty"`  //integer types, minimum number size
	Maximum           int64             `json:"maximum,omitempty"`  //integer types, maximum number size
	Format            string            `json:"format,omitempty"`   //string types, email, ipv4, ipv6, uri, hostname, date-time,
	Optional          bool              `json:"optional,omitempty"` //Is this field optional?
	Pattern           string            `json:"pattern,omitempty"`  //string types, regex pattern matching e.g. "^[A-Z]{2}-[0-9]{5}$"
	EnumInt           []int64           `json:"enum,omitempty"`
	MinProperties     int64             `json:"minproperties,omitempty"` //Minimum number of required properties
	MaxProperties     int64             `json:"maxproperties,omitempty"` //Maximum number of required properties
	MinLength         int64             `json:"minLength,omitempty"`     //string types, minimum length of string
	MaxLength         int64             `json:"maxLength,omitempty"`     //string types, maximum length of string
	//Enum              []string          `json:"enum,omitempty"`     //string types, Contains an explicit list of options
}

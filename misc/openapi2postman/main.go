package main

import (
	"fmt"

	"github.com/grokify/swaggman"
	"github.com/grokify/swaggman/postman2"
)

func main() {
	// Instantiate a converter with default configuration
	conv := swaggman.NewConverter(swaggman.Configuration{
		PostmanURLHostname: "{{url}}",
		PostmanHeaders: []postman2.Header{
			{
				Key:   "Authorization",
				Value: "Bearer {{authorization}}",
			},
		},
	})

	// Convert a Swagger spec
	err := conv.Convert("xegony.json", "xegony.postman.json")
	if err != nil {
		fmt.Println("Failed to convert", err.Error())
	}
}

//go:generate fileb0x box.json

package main

import (
	"fmt"
	"net/http"

	"github.com/xackery/xegony/static"
)

func main() {
	http.ListenAndServe(":8080", static.Handler)
	fmt.Println("Started")
	select {}
}

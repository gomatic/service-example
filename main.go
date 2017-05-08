//go:generate go-bindata -o swagger.json.go model/model.swagger.json

package main

import (
	"github.com/gomatic/servicer/server"
)

//
func main() {
	server.Main(run, "", "", MustAsset("model/model.swagger.json"))
}

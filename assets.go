// +build generate

package main

import (
	"net/http"

	"github.com/shurcooL/vfsgen"
)

var fs http.FileSystem = http.Dir("./templates")

func main() {

	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:     "./internal/asset_fs.go",
		PackageName:  "internal",
		VariableName: "Templates",
	})
	if err != nil {
		panic(err)
	}
}

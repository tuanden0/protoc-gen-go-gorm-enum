package main

import (
	protocgengogormenum "github.com/tuanden0/protoc-gen-go-gorm-enum"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			protocgengogormenum.GenerateEnumFile(gen, f)
		}
		return nil
	})
}

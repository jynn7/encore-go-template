package main

import (
	"fmt"
	"io"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
	"encore.app/database/models"
)

var model_list = []any{
	&models.User{},
}

func main() {
	stmts, err := gormschema.New("postgres").Load(model_list...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}

package main

import (
	"fmt"
	"github.com/nvkv/halp/pkg/config/v1"
	"github.com/nvkv/halp/pkg/datasources/googlesheets/v1"
	"github.com/nvkv/halp/pkg/types/data/v1"
	"github.com/nvkv/halp/pkg/types/datasource/v1"
)

func main() {
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		panic(err)
	}

	halpSheet := googlesheets.Spreadsheet{
		Credentials: cfg.Datasource.GoogleSheets.CredentialsFilePath,
		Tokenfile:   cfg.Datasource.GoogleSheets.TokenFilePath,
		SheetID:     cfg.Datasource.GoogleSheets.SheetID,
	}

	meals, err := halpSheet.AllMeals()
	if err != nil {
		panic(err)
	}

	for _, meal := range meals {
		fmt.Printf("%#v\n", meal)
	}

	searchResults, err := halpSheet.Select(datasource.Query{
		datasource.MealTypeField: data.Lunch,
		datasource.IsLentenField: true,
	})
	if err != nil {
		panic(err)
	}

	for _, meal := range searchResults {
		fmt.Printf("FOUND: %#v\n", meal)
	}
}

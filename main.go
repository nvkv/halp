package main

import (
	"fmt"
	"github.com/nvkv/halp/pkg/datasources/dummy/v1"
	"github.com/nvkv/halp/pkg/types/data/v1"
	"github.com/nvkv/pkg/types/datasource/v1"
)

func main() {
	ds := dummy.DummyDatasource{}
	meals := ds.Select(datasource.Query{
		MealType: data.Breakfast,
		IsLenten: false,
		IsLavish: false,
	})
	for _, meal := range meals {
		fmt.Println(meal)
	}
}

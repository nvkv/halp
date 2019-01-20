package main

import (
	"fmt"
	"github.com/nvkv/halp/pkg/datasource/v1"
	"github.com/nvkv/halp/pkg/datasource/v1/dummy"
	"github.com/nvkv/halp/pkg/types/v1"
)

func main() {
	ds := dummy.DummyDatasource{}
	meals := ds.Select(datasource.Query{
		MealType: types.Breakfast,
		IsLenten: false,
		IsLavish: false,
	})
	for _, meal := range meals {
		fmt.Println(meal)
	}
}

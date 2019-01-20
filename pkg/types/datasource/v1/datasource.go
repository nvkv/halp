package datasource

import (
	"github.com/nvkv/halp/pkg/types/data/v1"
)

type Query struct {
	MealType data.MealType
	IsLenten bool
	IsLavish bool
}

type Datasource interface {
	AllMeals() ([]data.Meal, error)
	Select(query Query) ([]data.Meal, error)
}

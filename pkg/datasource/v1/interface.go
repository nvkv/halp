package datasource

import (
	"github.com/nvkv/halp/pkg/types/v1"
)

type Query struct {
	MealType types.MealType
	IsLenten bool
	IsLavish bool
}

type Datasource interface {
	AllMeals() []types.Meal
	Select(query Query) []types.Meal
}

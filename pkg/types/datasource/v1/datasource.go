package datasource

import (
	"github.com/nvkv/halp/pkg/types/data/v1"
)

type QueryField int

const (
	MealTypeField = iota
	IsFastenField
	IsLavishField
)

type Query map[QueryField]interface{}

type Datasource interface {
	AllMeals() ([]data.Meal, error)
	Select(query Query) ([]data.Meal, error)
}

package dummy

import (
	"github.com/nvkv/halp/pkg/datasource/v1"
	"github.com/nvkv/halp/pkg/types/v1"
)

type DummyDatasource struct {
	meals []types.Meal
}

func dummySet() []types.Meal {
	return []types.Meal{
		types.Meal{
			Name:     "Irish breakfast",
			IsLenten: false,
			IsLavish: false,
			Type:     types.Breakfast,
		},
		types.Meal{
			Name:     "Sushi",
			IsLenten: true,
			IsLavish: true,
			Type:     types.Lunch,
		},
		types.Meal{
			Name:     "Shwarma",
			IsLenten: false,
			IsLavish: false,
			Type:     types.Dinner,
		},
	}
}

func (ds DummyDatasource) AllMeals() []types.Meal {
	if len(ds.meals) == 0 {
		ds.meals = dummySet()
	}

	return ds.meals
}

func (ds DummyDatasource) Select(query datasource.Query) []types.Meal {
	result := []types.Meal{}
	for _, meal := range ds.AllMeals() {
		if meal.Type == query.MealType && meal.IsLenten == query.IsLenten && meal.IsLavish == query.IsLavish {
			result = append(result, meal)
		}
	}
	return result
}

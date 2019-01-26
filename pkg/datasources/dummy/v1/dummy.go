package dummy

import (
	data "github.com/nvkv/halp/pkg/types/data/v1"
	"github.com/nvkv/halp/pkg/types/datasource/v1"
)

type DummyDatasource struct {
	meals []data.Meal
}

func dummySet() []data.Meal {
	return []data.Meal{
		data.Meal{
			Name:     "Irish breakfast",
			IsLenten: false,
			IsLavish: false,
			Type:     data.Breakfast,
		},
		data.Meal{
			Name:     "Sushi",
			IsLenten: true,
			IsLavish: true,
			Type:     data.Lunch,
		},
		data.Meal{
			Name:     "Shwarma",
			IsLenten: false,
			IsLavish: false,
			Type:     data.Dinner,
		},
	}
}

func (ds DummyDatasource) AllMeals() ([]data.Meal, error) {
	if len(ds.meals) == 0 {
		ds.meals = dummySet()
	}

	return ds.meals, nil
}

func (ds DummyDatasource) Select(query datasource.Query) ([]data.Meal, error) {
	result := []data.Meal{}
	meals, _ := ds.AllMeals()
	for _, meal := range meals {
		if meal.Type == query.MealType && meal.IsLenten == query.IsAFastDay && meal.IsLavish == query.IsLavish {
			result = append(result, meal)
		}
	}
	return result, nil
}

package googlesheets

import (
	"github.com/nvkv/halp/pkg/types/data/v1"
	"github.com/nvkv/halp/pkg/types/datasource/v1"
)

type Spreadsheet struct {
	Credentials string
	Tokenfile   string
	SheetID     string
	meals       []data.Meal
}

func (s Spreadsheet) AllMeals() ([]data.Meal, error) {
	if len(s.meals) > 0 {
		return s.meals, nil
	}
	meals, err := fetchAll(
		s.Credentials,
		s.Tokenfile,
		s.SheetID,
	)
	if err != nil {
		return nil, err
	}

	s.meals = meals
	return s.meals, nil
}

func (s Spreadsheet) Select(query datasource.Query) ([]data.Meal, error) {
	result := []data.Meal{}
	meals, err := s.AllMeals()
	if err != nil {
		return nil, err
	}
	for _, meal := range meals {
		if meal.Type == query.MealType && meal.IsLenten == query.IsLenten && meal.IsLavish == query.IsLavish {
			result = append(result, meal)
		}
	}
	return result, nil
}

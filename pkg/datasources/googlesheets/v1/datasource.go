package googlesheets

import (
	"github.com/nvkv/halp/pkg/types/data/v1"
	"github.com/nvkv/halp/pkg/types/datasource/v1"
)

type Spreadsheet struct {
	Credentials string
	Tokenfile   string
	SheetID     string
	Range       string
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
		s.Range,
	)
	if err != nil {
		return nil, err
	}

	s.meals = meals
	return s.meals, nil
}

func checkFieldValue(meal data.Meal, field datasource.QueryField, value interface{}) bool {
	switch field {
	case datasource.MealTypeField:
		return meal.Type == data.MealType(value.(int))
	case datasource.IsAFastDayField:
		return meal.IsLenten == value.(bool)
	case datasource.IsLavishField:
		return meal.IsLavish == value.(bool)
	default:
		return false
	}
}

func (s Spreadsheet) Select(query datasource.Query) ([]data.Meal, error) {
	result := []data.Meal{}
	meals, err := s.AllMeals()
	if err != nil {
		return nil, err
	}

	for _, meal := range meals {
		var matching = true
		for k, v := range query {
			if checkFieldValue(meal, k, v) == false {
				matching = false
				break
			}
		}
		if matching {
			result = append(result, meal)
		}
	}
	return result, nil
}

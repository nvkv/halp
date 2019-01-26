package schedule

import (
	"math/rand"
	"time"

	"github.com/nvkv/halp/pkg/types/data/v1"
	"github.com/nvkv/halp/pkg/types/datasource/v1"
)

// Will return empty meal if input is empty
func pickRandomMeal(meals []data.Meal) data.Meal {
	mealCount := len(meals)
	if mealCount == 0 {
		return data.Meal{}
	}
	return meals[rand.Intn(mealCount)]
}

func ScheduleDay(date time.Time, ds datasource.Datasource) (data.Day, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	isHoliday := data.IsHoliday(date)
	isLenten := data.IsLenten(date)

	breakfasts, err := ds.Select(datasource.Query{
		datasource.MealTypeField: data.Breakfast,
		datasource.IsLentenField: isLenten,
		datasource.IsLavishField: isHoliday,
	})
	if err != nil {
		return data.Day{}, err
	}
	breakfast := pickRandomMeal(breakfasts)

	lunches, err := ds.Select(datasource.Query{
		datasource.MealTypeField: data.Lunch,
		datasource.IsLentenField: isLenten,
		datasource.IsLavishField: isHoliday,
	})
	if err != nil {
		return data.Day{}, err
	}
	lunch := pickRandomMeal(lunches)

	dinners, err := ds.Select(datasource.Query{
		datasource.MealTypeField: data.Dinner,
		datasource.IsLentenField: isLenten,
		datasource.IsLavishField: isHoliday,
	})
	if err != nil {
		return data.Day{}, err
	}
	dinner := pickRandomMeal(dinners)

	day := data.Day{
		Date:      date,
		Breakfast: breakfast,
		Lunch:     lunch,
		Dinner:    dinner,
	}

	if err := day.Validate(); err != nil {
		return data.Day{}, err
	}
	return day, nil
}

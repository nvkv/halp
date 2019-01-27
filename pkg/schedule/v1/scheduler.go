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

func nextNDays(from time.Time, n int) []time.Time {
	var dates = []time.Time{}
	current := from
	for i := 0; i < n; i++ {
		current = current.AddDate(0, 0, 1)
		dates = append(dates, current)
	}
	return dates
}

func ScheduleWeek(date time.Time, ds datasource.Datasource) ([]data.Day, error) {
	dates := nextNDays(date, 7)
	schedule := []data.Day{}
	for _, d := range dates {
		day, err := ScheduleDay(d, ds)
		if err != nil {
			return nil, err
		}
		schedule = append(schedule, day)
	}
	return schedule, nil
}

func ScheduleDay(date time.Time, ds datasource.Datasource) (data.Day, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	isHoliday := data.IsHoliday(date)
	isAFastDay := data.IsAFastDay(date)

	breakfasts, err := ds.Select(datasource.Query{
		datasource.MealTypeField:   data.Breakfast,
		datasource.IsAFastDayField: isAFastDay,
		datasource.IsLavishField:   isHoliday,
	})
	if err != nil {
		return data.Day{}, err
	}
	breakfast := pickRandomMeal(breakfasts)

	lunches, err := ds.Select(datasource.Query{
		datasource.MealTypeField:   data.Lunch,
		datasource.IsAFastDayField: isAFastDay,
		datasource.IsLavishField:   isHoliday,
	})
	if err != nil {
		return data.Day{}, err
	}
	lunch := pickRandomMeal(lunches)

	dinners, err := ds.Select(datasource.Query{
		datasource.MealTypeField:   data.Dinner,
		datasource.IsAFastDayField: isAFastDay,
		datasource.IsLavishField:   isHoliday,
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

package types

import (
	"fmt"
	"time"
)

type Day struct {
	Date       time.Time
	Breakfast  Meal
	Lunch      Meal
	Dinner     Meal
	Snack      Meal
	ExtraMeals map[string]Meal
}

func (d Day) IsHoliday() bool {
	switch d.Date.Weekday() {
	case time.Saturday:
		fallthrough
	case time.Sunday:
		return true
	default:
		return false
	}
}

func (d Day) IsLenten() bool {
	// TODO: This is default Orthodox lenten weekdays
	// Something more complicated should be implemented later
	switch d.Date.Weekday() {
	case time.Wednesday:
		fallthrough
	case time.Friday:
		return true
	default:
		return false
	}
}

func (d Day) Validate() error {
	// Collect all meals in a slice
	meals := []Meal{
		d.Breakfast,
		d.Lunch,
		d.Dinner,
		d.Snack,
	}
	for _, v := range d.ExtraMeals {
		meals = append(meals, v)
	}

	// Lent check
	if d.IsLenten() {
		for _, meal := range meals {
			if meal.IsLenten == false {
				return fmt.Errorf("%v supposed to be lenten, but meal '%v' was planned, which is not", d.Date, meal.Name)
			}
		}
	}
	return nil
}

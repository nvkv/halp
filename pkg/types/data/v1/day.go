package data

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

func IsHoliday(date time.Time) bool {
	switch date.Weekday() {
	case time.Saturday:
		fallthrough
	case time.Sunday:
		return true
	default:
		return false
	}
}

func (d Day) IsHoliday() bool {
	return IsHoliday(d.Date)
}

func IsFasten(date time.Time) bool {
	// TODO: This is default Orthodox lenten weekdays
	// Something more complicated should be implemented later
	switch date.Weekday() {
	case time.Wednesday:
		fallthrough
	case time.Friday:
		return true
	default:
		return false
	}
}

func (d Day) IsFasten() bool {
	return IsFasten(d.Date)
}

func (d Day) AllMeals() []Meal {
	meals := []Meal{
		d.Breakfast,
		d.Lunch,
		d.Dinner,
		d.Snack,
	}
	for _, v := range d.ExtraMeals {
		meals = append(meals, v)
	}
	return meals
}

func (d Day) Validate() error {
	// Collect all meals in a slice
	meals := d.AllMeals()
	// Fast check
	if d.IsFasten() {
		for _, meal := range meals {
			if meal.IsFasten == false {
				return fmt.Errorf("%v supposed to be lenten, but meal '%v' was planned, which is not", d.Date, meal.Name)
			}
		}
	}
	return nil
}

func (d Day) String() string {
	fast := "No fast today!"
	if d.IsFasten() {
		fast = "Fast day!"
	}
	str := fmt.Sprintf(`
%v, %v

%s
`,
		d.Date.Format("2006-01-02"),
		d.Date.Weekday(),
		fast,
	)

	for _, meal := range d.AllMeals() {
		if meal != (Meal{}) {
			str += fmt.Sprintf("%s\n", meal)
		}
	}

	return str
}

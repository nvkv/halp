package data

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"time"

	"github.com/nvkv/halp/pkg/testhelpers/v1"
)

func (d Day) Generate(rand *rand.Rand, size int) reflect.Value {
	day := Day{
		Date:      testhelpers.RandomDate(rand),
		Breakfast: randomMeal(rand),
		Lunch:     randomMeal(rand),
		Dinner:    randomMeal(rand),
		Snack:     randomMeal(rand),
	}
	return reflect.ValueOf(day)
}

func TestAllMeals(t *testing.T) {
	checkAllMeals := func(d Day) bool {
		expected := []Meal{
			d.Breakfast,
			d.Lunch,
			d.Dinner,
			d.Snack,
		}

		for _, m := range d.ExtraMeals {
			expected = append(expected, m)
		}
		return reflect.DeepEqual(expected, d.AllMeals())
	}

	if err := quick.Check(checkAllMeals, testhelpers.DefaultConfig); err != nil {
		t.Error(err)
	}
}

// These are basic tests, will test only
// Wednesday and Friday's Orthodox lent
func TestFastDetection(t *testing.T) {

	checkFast := func(seed int64) bool {
		randSrc := rand.NewSource(seed)
		rand := rand.New(randSrc)
		date := testhelpers.RandomDate(rand)

		var shouldBeLenten = false
		if date.Weekday() == time.Wednesday || date.Weekday() == time.Friday {
			shouldBeLenten = true
		}
		return IsLenten(date) == shouldBeLenten
	}

	if err := quick.Check(checkFast, testhelpers.DefaultConfig); err != nil {
		t.Error(err)
	}
}

func TestDayFastDetection(t *testing.T) {
	checkDayFast := func(d Day) bool {
		var shouldBeLenten = false
		if d.Date.Weekday() == time.Wednesday || d.Date.Weekday() == time.Friday {
			shouldBeLenten = true
		}
		return d.IsLenten() == shouldBeLenten
	}

	if err := quick.Check(checkDayFast, testhelpers.DefaultConfig); err != nil {
		t.Error(err)
	}
}

func TestHolidayDetection(t *testing.T) {
	checkHoliday := func(d Day) bool {
		var shouldBeHoliday = false
		if d.Date.Weekday() == time.Sunday || d.Date.Weekday() == time.Saturday {
			shouldBeHoliday = true
		}
		return shouldBeHoliday == d.IsHoliday()
	}

	if err := quick.Check(checkHoliday, testhelpers.DefaultConfig); err != nil {
		t.Error(err)
	}
}

func TestValidation(t *testing.T) {
	validate := func(d Day) bool {
		validation := d.Validate()
		if d.IsLenten() {
			var trulyLenten = true
			for _, meal := range d.AllMeals() {
				if meal.IsLenten != true {
					trulyLenten = false
					break
				}
			}
			if trulyLenten {
				return validation == nil
			} else {
				if validation == nil {
					fmt.Printf("%v is %v\n", d.Date, d.Date.Weekday().String())
					for _, meal := range d.AllMeals() {
						fmt.Printf("Meal lent status: %v\n", meal.IsLenten)
					}
				}
				return validation != nil
			}
		}
		return true
	}

	if err := quick.Check(validate, testhelpers.DefaultConfig); err != nil {
		t.Error(err)
	}
}

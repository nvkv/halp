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

// This is a basic test, it will test only
// Wednesday and Friday's Orthodox lent
func TestLentDetection(t *testing.T) {
	lent := func(d Day) bool {
		var shouldBeLenten = false
		if d.Date.Weekday() == time.Wednesday || d.Date.Weekday() == time.Friday {
			shouldBeLenten = true
		}
		return d.IsLenten() == shouldBeLenten
	}

	if err := quick.Check(lent, testhelpers.DefaultConfig); err != nil {
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

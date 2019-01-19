package types

import (
	//	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"time"
)

func randomDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func (d Day) Generate(rand *rand.Rand, size int) reflect.Value {
	day := Day{
		Date:      randomDate(),
		Breakfast: randomMeal(rand),
		Lunch:     randomMeal(rand),
		Dinner:    randomMeal(rand),
		Snack:     randomMeal(rand),
	}
	return reflect.ValueOf(day)
}

func TestValidation(t *testing.T) {
	gen := func(d Day) bool {
		validation := d.Validate()
		if d.IsLenten() {
			var trulyLenten = false
			for _, meal := range d.AllMeals() {
				if meal.IsLenten != true {
					trulyLenten = false
					break
				}
			}
			if trulyLenten {
				return validation == nil
			} else {
				return validation != nil
			}
		}
		return true
	}

	if err := quick.Check(gen, nil); err != nil {
		t.Error(err)
	}
}

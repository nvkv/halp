package data

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	"github.com/nvkv/halp/pkg/testhelpers/v1"
)

func randomMeal(rand *rand.Rand) Meal {
	idv, _ := quick.Value(reflect.TypeOf(""), rand)
	namev, _ := quick.Value(reflect.TypeOf(""), rand)

	id := idv.String()
	name := namev.String()

	meal := Meal{
		Id:       id,
		Type:     MealType(rand.Int31n(4)),
		Name:     name,
		IsLenten: rand.Int31n(2) > 0,
		IsLavish: rand.Int31n(2) > 0,
	}
	return meal
}

func (t MealType) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(MealType(rand.Intn(4)))
}

func (m Meal) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(randomMeal(rand))
}

func TestMealString(t *testing.T) {
	checkString := func(m Meal) bool {
		return m.String() == fmt.Sprintf("%s: %s", m.Type, m.Name)
	}

	if err := quick.Check(checkString, testhelpers.DefaultConfig); err != nil {
		t.Error(err)
	}
}

func TestMealTypeString(t *testing.T) {
	checkString := func(t MealType) bool {
		str := t.String()
		switch t {
		case Breakfast:
			return str == "Breakfast"
		case Lunch:
			return str == "Lunch"
		case Dinner:
			return str == "Dinner"
		case Snack:
			return str == "Snack"
		default:
			return false
		}
	}

	if err := quick.Check(checkString, testhelpers.DefaultConfig); err != nil {
		t.Error(err)
	}
}

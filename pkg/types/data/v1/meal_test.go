package data

import (
	"math/rand"
	"reflect"
	"testing/quick"
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

func (m Meal) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(randomMeal(rand))
}

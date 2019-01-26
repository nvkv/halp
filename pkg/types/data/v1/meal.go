package data

import "fmt"

type MealType int

const (
	Breakfast = iota
	Lunch
	Dinner
	Snack
)

func (t MealType) String() string {
	switch t {
	case Breakfast:
		return "Breakfast"
	case Lunch:
		return "Lunch"
	case Dinner:
		return "Dinner"
	case Snack:
		return "Snack"
	default:
		return "Unknown type of meal"
	}
}

type Meal struct {
	Id       string
	Type     MealType
	Name     string
	IsFasten bool
	IsLavish bool
}

func (m Meal) String() string {
	return fmt.Sprintf("%s: %s", m.Type, m.Name)
}

package data

type MealType int

const (
	Breakfast = iota
	Lunch
	Dinner
	Snack
)

type Meal struct {
	Id       string
	Type     MealType
	Name     string
	IsLenten bool
	IsLavish bool
}

package telegram

import (
	"fmt"
	"strings"

	"github.com/nvkv/halp/pkg/types/data/v1"
)

func formatDaySchedule(schedule data.Day) string {
	meals := []string{}
	for _, meal := range schedule.AllMeals() {
		meals = append(meals, meal.String())
	}

	return fmt.Sprintf(`А что если так?

%v`,
		strings.Join(meals, "\n"),
	)
}

package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	if data != "" {
		r := []rune(data)
		if unicode.IsSpace(r[0]) {
			return 0, 0, fmt.Errorf("invalid steps data %q", data)
		}
	}

	clean := strings.TrimSpace(data)

	parts := strings.Split(clean, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid format: expected '<steps>,<training duration>'")
	}

	stepsStr := parts[0]
	if strings.TrimSpace(stepsStr) != stepsStr {
		return 0, 0, fmt.Errorf("invalid steps data %q", stepsStr)
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps data %q: %w", stepsStr, err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("invalid steps data %q", stepsStr)
	}

	durStr := strings.TrimSpace(parts[1])
	if durStr == "" {
		return 0, 0, fmt.Errorf("duration part is empty")
	}
	dur, err := time.ParseDuration(durStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration %q: %w", durStr, err)
	}
	if dur <= 0 {
		return 0, 0, fmt.Errorf("invalid duration %q", durStr)
	}

	return steps, dur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println("Failed to parse package:", err)
		return ""
	}

	if steps <= 0 || duration <= 0 {
		log.Println("Invalid data:", steps, duration)
		return ""
	}

	distance := float64(steps) * stepLength
	distanceKm := distance / float64(mInKm)

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println("Walking calories calc failed:", err)
		return ""
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps,
		distanceKm,
		calories,
	)

	return result
}

package daysteps

import (
	"fmt"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"

	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	dataSplit := strings.Split(data, ",")
	if len(dataSplit) != 2 {
		return 0, 0, fmt.Errorf("длинна строки < 2")
	}

	steps, err := strconv.Atoi(dataSplit[0])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования в целое число")
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	durationWalk, err := time.ParseDuration(dataSplit[1])
	if err != nil {
		return 0, 0, fmt.Errorf("строка не соответствует указанному формату")
	}
	return steps, durationWalk, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	countSteps, durationWalk, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if countSteps <= 0 {
		return ""
	}
	distanceKm := float64(countSteps) * stepLength / mInKm
	calories, err := spentcalories.WalkingSpentCalories(countSteps, weight, height, durationWalk)
	if err != nil {
		fmt.Println(err)
	}
	result := fmt.Sprintf("Количество шагов %d\n Дистанция составила %d\n Вы сожгли %d\n", countSteps, distanceKm, calories)
	return result
}

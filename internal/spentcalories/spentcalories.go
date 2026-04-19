package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	dataSplit := strings.Split(data, ",")
	if len(dataSplit) != 3 {
		return 0, "", 0, fmt.Errorf("длинна строки < 3")
	}

	steps, err := strconv.Atoi(dataSplit[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования в целое число")
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	durationWalk, err := time.ParseDuration(dataSplit[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("строка не соответствует указанному формату")
	}
	if durationWalk <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность тренировки должна быть положительной")
	}
	return steps, dataSplit[1], durationWalk, nil
}

func distance(steps int, height float64) float64 {

	stepLength := stepLengthCoefficient * height
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	hours := duration.Hours()
	return distance(steps, height) / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {

	steps, typeTraning, durationWalk, err := parseTraining(data)
	if err != nil {
		log.Println("ошибка при парсинге данных:", err)
		return "", err
	}
	hours := durationWalk.Hours()
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, durationWalk)

	var calories float64
	switch typeTraning {
	case "Ходьба":
		calories, _ = WalkingSpentCalories(steps, weight, height, durationWalk)
	case "Бег":
		calories, _ = RunningSpentCalories(steps, weight, height, durationWalk)
	default:
		log.Println("неизвестный тип тренировки")
		return "", fmt.Errorf("неизвестный тип тренировки: %s", typeTraning)
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		typeTraning, hours, dist, speed, calories)
	return result, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("не верное значение")
	}
	return (meanSpeed(steps, height, duration) * weight * duration.Minutes()) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("не верное значение")
	}
	return (meanSpeed(steps, height, duration) * weight * duration.Minutes()) / minInH * walkingCaloriesCoefficient, nil
}

package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories рассчитывает калл. при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("all parameters must be positive")
	}
	midSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * midSpeed * durationInMinutes) / minInH
	calories = calories * walkingCaloriesCoefficient
	return calories, nil
}

// RunningSpentCalories рассчитывает калл. при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("all parameters must be positive")
	}
	midSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * midSpeed * durationInMinutes) / minInH
	return calories, nil
}

// MeanSpeed рассчитывает среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

// Distance рассчитывает дистанцию в км. по кол-ву шагов и росту
func Distance(steps int, height float64) float64 {
	return ((height * stepLengthCoefficient) * float64(steps)) / mInKm
}

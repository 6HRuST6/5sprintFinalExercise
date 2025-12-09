package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories рассчитывает количество калорий,
// потраченных при ходьбе за указанное количество шагов.
// Возвращает количество калорий и ошибку, если входные данные некорректны.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	if err := validateInput(steps, weight, height, duration); err != nil {
		return 0, err
	}
	mSpeed := MeanSpeed(steps, height, duration)

	minut := duration.Minutes()

	return ((weight * mSpeed * minut) / minInH) * walkingCaloriesCoefficient, nil

}

// WalkingSpentCalories рассчитывает количество калорий,
// потраченных при беге за указанное количество шагов.
// Возвращает количество калорий и ошибку, если входные данные некорректны.

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if err := validateInput(steps, weight, height, duration); err != nil {
		return 0, err
	}

	mSpeed := MeanSpeed(steps, height, duration)

	minut := duration.Minutes()

	return (weight * mSpeed * minut) / minInH, nil

}

// MeanSpeed рассчитывает среднюю скорость движения в км/ч
// по количеству шагов, росту и длительности движения.
//
// Если steps <= 0 или duration <= 0, функция возвращает 0.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}

	dist := Distance(steps, height)
	hour := duration.Hours()

	return dist / hour

}

// Distance вычисляет пройденное расстояние в километрах
// по количеству шагов и росту человека.
func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	lenghtStep := height * stepLengthCoefficient

	dist := (lenghtStep * float64(steps)) / mInKm

	return dist
}

// validateInput проверяет корректность входных данных для функций,
// рассчитывающих потраченные калории.
func validateInput(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return fmt.Errorf("вес должен быть положительным")
	}
	if height <= 0 {
		return fmt.Errorf("рост должен быть положительным")
	}
	if duration <= 0 {
		return fmt.Errorf("длительность должна быть положительной")
	}
	return nil
}

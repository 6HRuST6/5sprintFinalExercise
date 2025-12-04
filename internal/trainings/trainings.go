package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "шаги,тип,длительность" и
// заполняет поля структуры Training.
// Возвращает ошибку, если строка имеет неверный формат или
// данные не удалось преобразовать.
func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	sl := strings.Split(datastring, ",")
	if len(sl) != 3 {
		return fmt.Errorf("Количество параметров не соответствует ,получено %d", len(sl))
	}
	step := sl[0]
	trType := strings.TrimSpace(sl[1])
	durat := strings.TrimSpace(sl[2])

	st, err := strconv.Atoi(step)
	if err != nil {
		return fmt.Errorf("не удалось преобразовать шаги %q: %w", step, err)
	}
	if st <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	t.Steps = st

	t.TrainingType = trType

	dur, err := time.ParseDuration(durat)
	if err != nil {
		return fmt.Errorf("не удалось преобразовать время %q :%w", durat, err)
	}
	if dur <= 0 {
		return fmt.Errorf("значение длительности должно быть больше 0,получено %q", durat)
	}

	t.Duration = dur

	return nil
}

// ActionInfo возвращает строку с информацией о тренировке:
// тип, длительность, дистанцию, скорость и потраченные калории.
// В случае ошибки расчёта калорий возвращает ошибку.
func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dist := spentenergy.Distance(t.Steps, t.Height)
	mSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var spCallories float64
	var err error
	switch t.TrainingType {

	case "Ходьба":

		spCallories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	case "Бег":

		spCallories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", fmt.Errorf("ошибка вычисления потраченных каллорий: %w", err)
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, mSpeed, spCallories), nil
}

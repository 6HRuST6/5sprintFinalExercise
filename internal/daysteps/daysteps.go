package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "шаги, длительность" (например: "10000, 30m")
// и заполняет поля структуры DaySteps.
// В случае ошибки парсинга или некорректных значений возвращает ошибку.
func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	sl := strings.Split(datastring, ",")
	if len(sl) != 2 {
		return fmt.Errorf("Количество параметров не соответствует ,получено %d", len(sl))
	}
	step := sl[0]
	durat := strings.TrimSpace(sl[1])
	st, err := strconv.Atoi(step)
	if err != nil {
		return fmt.Errorf("не удалось преобразовать шаги %q: %w", step, err)
	}
	if st <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	dur, err := time.ParseDuration(durat)
	if err != nil {
		return fmt.Errorf("не удалось преобразовать время %q :%w", durat, err)
	}
	if dur <= 0 {
		return fmt.Errorf("значение длительности должно быть больше 0,получено %q", durat)
	}

	ds.Steps = st
	ds.Duration = dur

	return nil

}

// ActionInfo рассчитывает дистанцию и количество потраченных калорий
// за день
// В случае ошибки расчёта калорий возвращает её.
func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	spCallories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка вычисления потраченных калорий: %w", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, spCallories), nil
}

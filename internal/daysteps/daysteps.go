package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку с данными о прогулке
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("Не правильный формат строки, нужно 2 части!")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("Не правильный формат количества шагов")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return errors.New("Не правильный формат продолжительности")
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// ActionInfo формирует строку с информацией о прогулке
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли %.2f ккал.", ds.Steps, distance, calories)
	return info, nil
}

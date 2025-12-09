package daysteps

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
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

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}

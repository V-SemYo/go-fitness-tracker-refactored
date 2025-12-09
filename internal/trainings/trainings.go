package trainings

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку с данными о тренировке и сохраняет в структуру Training
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("Не правильный формат строки, нужно 3 части!")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("Не правильный формат количества шагов")
	}
	trainingType := parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return errors.New("Не правильный формат продолжительности")
	}
	t.Steps = steps
	t.TrainingType = trainingType
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}

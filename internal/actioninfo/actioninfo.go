package actioninfo

import "log"

// DataParser итерфейс для структур Training и DaySteps
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info обрабатывает набор данных с помощью интерфейса DataParser
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println("Ошибка парсинга: ", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("Ошибка создания информации: ", err)
			continue
		}
		log.Println(info)
	}
}

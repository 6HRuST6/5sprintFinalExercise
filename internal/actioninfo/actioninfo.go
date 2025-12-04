package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {

		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка парсинга строки %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования отчёта для строки %q: %v", data, err)
			continue
		}

		fmt.Println(info)
	}
}

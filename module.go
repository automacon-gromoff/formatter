package formatter

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type patient struct {
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Email string `json:"email"`
}

func Do(patients string, result string) error {
	// Считываем данные
	src, err := decode(patients)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	// Сортируем
	sort.Slice(src, func(i, j int) bool {
		return src[i].Age < src[j].Age
	})

	// Записываем данные
	err = encode(src, result)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func decode(patients string) ([]patient, error) {
	src := make([]patient, 0, 3)

	// открываем файл
	f, err := os.Open(patients)
	if err != nil {
		return src, fmt.Errorf("ошибка открытия файла %s: %s", patients, err)
	}

	// отложенное закрытие
	defer f.Close()

	// Получаем данные json
	dec := json.NewDecoder(f)
	for dec.More() {
		c := patient{}
		err = dec.Decode(&c)
		if err != nil {
			return src, fmt.Errorf("ошибка чтения файла %s: %s", patients, err)
		}
		src = append(src, c)
	}

	return src, nil
}

func encode(res []patient, result string) error {
	// Создаем файл
	f, err := os.Create(result)
	if err != nil {
		return fmt.Errorf("ошибка создания файла %s: %s", result, err)
	}

	// отложенное закрытие
	defer f.Close()

	// Записываем данные json
	err = json.NewEncoder(f).Encode(res)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл %s: %s", result, err)
	}

	return nil
}

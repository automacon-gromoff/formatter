package formatter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type patient struct {
	Name  string
	Age   uint
	Email string
}

type patients struct {
	List []patient `xml:"Patient"`
}

func Do(patients string, result string) error {
	// Считываем данные
	src, err := decode(patients)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

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

	// Записываем данные xml
	p := patients{}
	p.List = append(p.List, res...)

	f.WriteString(xml.Header)

	enc := xml.NewEncoder(f)
	enc.Indent("", " ")
	err = enc.Encode(p)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл %s: %s", result, err)
	}

	return nil
}

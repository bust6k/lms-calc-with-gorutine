package API

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"project_yandex_lms/structures"
	"strings"
)

func GetNewTask(url string) (structures.Task, error) {
	resp, err := http.Get(url)

	if err != nil {
		return structures.Task{}, fmt.Errorf("ошибка при запросе для получения новых задач: %v", err)
	}
	bytesresp, errread := io.ReadAll(resp.Body)
	if errread != nil {
		return structures.Task{}, fmt.Errorf("ошибка при чтении тела ответа:%v", errread)
	}
	defer resp.Body.Close()
	if strings.Contains(string(bytesresp), "слайс задач пуст") {
		return structures.Task{}, fmt.Errorf("ошибка, новых задач пока что нет")

	}

	var NewTask structures.Task
	err = json.Unmarshal(bytesresp, &NewTask)
	if err != nil {
		return structures.Task{}, fmt.Errorf("ошибка при десериализации ответа: %v", err)
	}

	return NewTask, nil
}

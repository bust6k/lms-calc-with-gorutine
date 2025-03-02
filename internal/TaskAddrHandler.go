package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"project_yandex_lms/structures"
	"project_yandex_lms/variables"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		bytesresp, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "ошибка при попытке прочитать запрос", 422)
		}

		err = json.Unmarshal(bytesresp, &variables.CurrentTask)
		if err != nil {
			http.Error(w, "ошибка при десеарлизации ответа", 500)
		}

	} else if r.Method == http.MethodGet {
		bytesCurrentTask, err := json.Marshal(variables.CurrentTask)
		if err != nil {
			http.Error(w, "ошибка при сереализации ответа", 500)

		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytesCurrentTask)
		variables.CurrentTask = structures.Task{0, 0, 0, "", variables.CurrentTask.Operation_time}

	}
}

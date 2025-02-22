package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"project_yandex_lms/structures"
)

func InteralHandler(w http.ResponseWriter, r *http.Request) {

	var slicceOfTasks []structures.Task
	buffer, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ошибка при чтении данных запроса", 422)
		//}
		defer r.Body.Close()

		err = json.Unmarshal(buffer, &slicceOfTasks)
		if err != nil {
			http.Error(w, "ошибка при десеарилизации запроса", 422)
		}

		firstElement := slicceOfTasks[0]
		slicceOfTasks = slicceOfTasks[1:]
		if len(slicceOfTasks) == 0 {

			w.Write([]byte("слайс задач пуст"))
			return

		}
		bytesjson, err := json.Marshal(firstElement)
		if err != nil {
			http.Error(w, "ошибка при сериализации ответа", 500)
		}
		bytestopost := bytes.NewReader(bytesjson)
		http.Post("http://localhost:8080/internal/task", "application/json", bytestopost)
	}

}

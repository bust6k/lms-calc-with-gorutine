package important

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"project_yandex_lms/structures"
	"project_yandex_lms/variables"
)

func InteralHandler(w http.ResponseWriter, r *http.Request) {
	var slicceOfTasks []structures.Task
	buffer, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ошибка при чтении данных запроса", 422)

		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(buffer, &slicceOfTasks)
	if err != nil {
		http.Error(w, "ошибка при десеарилизации запроса", 422)

		return
	}

	if len(slicceOfTasks) == 0 {
		w.Write([]byte("слайс задач пуст"))
		return
	}

	variables.TheTasks = slicceOfTasks

	for len(variables.TheTasks) > 0 {
		firstElement := variables.TheTasks[0]
		variables.CurrentTask = firstElement

		variables.TheTasks = variables.TheTasks[1:]

		var bytesjson bytes.Buffer
		err = json.NewEncoder(&bytesjson).Encode(firstElement)
		if err != nil {
			fmt.Println(err)
		}

		if err != nil {
			http.Error(w, "ошибка при сериализации ответа", 422)

			return
		}

		bytestopost := bytes.NewReader(bytesjson.Bytes())

		_, err = http.Post("http://localhost:8080/internal/task", "application/json", bytestopost)
		if err != nil {
			log.Println("Ошибка при отправке задачи на сервер:", err)
		}
	}

	w.Write([]byte("Все задачи обработаны"))
}

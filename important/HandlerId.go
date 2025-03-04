package important

import (
	"encoding/json"
	"net/http"
	"project_yandex_lms/variables"
	"strconv"
)

func HandlerId(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Path[len("/api/v1/expressions/"):]

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		http.Error(w, "ошибка при преобразовании id в целочисленный тип", 422)
	}
	for _, expression := range variables.Expressions {
		if expression.Id == valueInt {
			bytesjson, err := json.Marshal(expression)
			if err != nil {
				http.Error(w, "ошибка при сереализации в json", 422)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(bytesjson)
		}
	}
}

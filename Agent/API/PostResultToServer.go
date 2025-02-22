package API

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project_yandex_lms/internal"
	"project_yandex_lms/variables"
)

func PostTaskToServer(result float64) error {
	var expression internal.Expression
	expression.Id = variables.Count_Root_Id
	expression.Status = "ready"
	expression.Result = result

	bytesexpression, err := json.Marshal(expression)
	if err != nil {
		return fmt.Errorf("ошибка при сериализации данных:%v", err)
	}

	body := bytes.NewReader(bytesexpression)

	http.Post("http://localhost:8080/api/v1/expressions", "application/json", body)
	return nil
}

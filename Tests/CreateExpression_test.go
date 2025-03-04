package Tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"project_yandex_lms/important"
	"project_yandex_lms/structures"
	"project_yandex_lms/variables"
	"strings"
	"testing"
)

func TestCreateRootExpressionHandlerBadCase(t *testing.T) {
	bytesresp, err := json.Marshal(structures.RootExpression{Expression: "2+2 ^2"})
	if err != nil {
		t.Error("ошибка при сереализации ")
	}
	req := httptest.NewRequest("POST", "http://localhost:8080/api/v1/calculate", strings.NewReader(string(bytesresp)))
	w := httptest.NewRecorder()
	important.CreateRootExpressionHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Error("ожидаемый статус код не соответствует результатуеее")
	}
}

func TestCreateRootExpressionHandlerGoodCase(t *testing.T) {
	bytesresp, err := json.Marshal(structures.RootExpression{Expression: "2+22"})
	if err != nil {
		t.Error("ошибка при сереализации ")
	}
	req := httptest.NewRequest("POST", "http://localhost:8080/api/v1/calculate", strings.NewReader(string(bytesresp)))
	w := httptest.NewRecorder()
	important.CreateRootExpressionHandler(w, req)
	resp := w.Result()
	bytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("ошибка при попытке прочитать ответ")
	}
	var fieldForOneUse structures.Expression
	err = json.Unmarshal(bytesBody, &fieldForOneUse)
	if err != nil {
		t.Error("ошибка при десеарилизации данных в структуру")
	}

	if resp.StatusCode != http.StatusCreated {
		t.Error("ожидаемый статус код не соответствует результатуеее")
	}
	for _, ex := range variables.Expressions {
		if ex.Id == fieldForOneUse.Id && ex.Result != fieldForOneUse.Result {
			t.Error("ошибка, результат посчитан неправильно")
		}
	}
}

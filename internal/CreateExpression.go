package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"project_yandex_lms/Agent"
	"project_yandex_lms/structures"

	"project_yandex_lms/variables"
)

func CreateRootExpressionHandler(w http.ResponseWriter, r *http.Request) {
	var newExpression structures.RootExpression
	variables.Count_Root_Id++
	newExpression.Id = variables.Count_Root_Id
	bytesreq, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ошибка при считывании запроса", 400)
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(bytesreq, &newExpression)
	if err != nil {
		http.Error(w, "ошибка при десеарилезации запроса ", 422)
		return
	}

	resp := struct {
		Id int `json:"id"`
	}{Id: newExpression.Id}
	byesresp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "ошибка при сериализации ответа ", 422)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(byesresp)
	if err != nil {
		http.Error(w, "не удалось отправить ответ", 500)
		return
	}

	rpns, _ := infixToRPN(string(bytesreq))
	node, _ := buildAST(rpns)
	splitexpressions := splitAST(node)
	ParseExpression(splitexpressions)
	Agent.CreateAgent(6)

}

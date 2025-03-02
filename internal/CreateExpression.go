package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"project_yandex_lms/Agent"
	"project_yandex_lms/structures"
	"strconv"
	"time"

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
		log.Println(err)

		if len(newExpression.Expression) == 0 {
			fmt.Println("ошибка у выражения 0 сиимволов")
		}
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
	rpns, err := InfixToRPN(newExpression.Expression)
	if err != nil {
		http.Error(w, "ошибка при преобразовании выражения в rpn", 400)
		return

	}

	node, err := BuildAST(rpns)
	if err != nil {
		http.Error(w, "ошибка при преобразовании rpn выражения в ast", 400)
		return
	}
	os.Setenv("TIME_ADDITION_MS", "10")
	os.Setenv("TIME_SUBSTRACTION_MS", "10")
	os.Setenv("TIME_MULTIPLICATIONS_MS", "15")
	os.Setenv("TIME_DIVISIONS_MS", "16")
	splitexpressions := SplitAST(node)

	for i := 0; i < len(splitexpressions); i++ {

		if splitexpressions[i].Operation == "+" {
			plus := os.Getenv("TIME_ADDITION_MS")

			valueplus, err := strconv.Atoi(plus)

			if err != nil {
				http.Error(w, "ошибка при задании  времени операции для задачи", 522)
			}
			splitexpressions[i].Operation_time = time.Duration(valueplus) * time.Millisecond

		} else if splitexpressions[i].Operation == "-" {
			minus := os.Getenv("TIME_SUBSTRACTION_MS")
			valueminus, err := strconv.Atoi(minus)
			if err != nil {
				http.Error(w, "ошибка при задании времени операции для задачи", 522)
			}
			splitexpressions[i].Operation_time = time.Duration(valueminus) * time.Millisecond
		} else if splitexpressions[i].Operation == "*" {
			multiplication := os.Getenv("TIME_MULTIPLICATIONS_MS")
			valuemultiplication, err := strconv.Atoi(multiplication)
			if err != nil {
				http.Error(w, "ошибка при задании времени операции для задачи", 522)
			}
			splitexpressions[i].Operation_time = time.Duration(valuemultiplication) * time.Millisecond
		} else if splitexpressions[i].Operation == "/" {
			divide := os.Getenv("TIME_DIVISIONS_MS")
			valuedivide, err := strconv.Atoi(divide)
			if err != nil {
				http.Error(w, "ошибка при задании времени операции для задачи", 522)
			}
			splitexpressions[i].Operation_time = time.Duration(valuedivide) * time.Millisecond
		}
	}
	err = PostExpression(splitexpressions)
	if err != nil {
		http.Error(w, "ошибка при парсинге в task", 400)

		log.Println(err)
	}

	os.Setenv("COMPUTING_POWER", "6")

	value := os.Getenv("COMPUTING_POWER")
	valueInt, _ := strconv.Atoi(value)
	Agent.CreateAgent(valueInt)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(byesresp)

	if err != nil {
		http.Error(w, "не удалось отправить ответ", 500)
		return
	}

}

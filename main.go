package main

import (
	"net/http"
	"project_yandex_lms/important"
)

func main() {
	http.HandleFunc("/", important.HandlerHome)
	http.HandleFunc("/api/v1/calculate", important.CreateRootExpressionHandler)
	http.HandleFunc("/api/v1/expressions", important.ExpressionsHandler)
	http.HandleFunc("/internal", important.InteralHandler)
	http.HandleFunc("/internal/task", important.TaskHandler)
	http.HandleFunc("/api/v1/expressionsSpecial", important.ExpressionsHandlerSpecial)
	http.HandleFunc("/api/v1/calculateSpecial", important.CreateRootExpressionHandlerSpecial)
	http.HandleFunc("/api/v1/expressions/", important.HandlerId)
	http.HandleFunc("/api/v1/expressionsSpecial/", important.HandlerIdSprcial)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"net/http"
	"project_yandex_lms/internal"
)

func main() {
	http.HandleFunc("/", internal.HandlerHome)
	http.HandleFunc("/api/v1/calculate", internal.CreateRootExpressionHandler)
	http.HandleFunc("/api/v1/expressions", internal.ExpressionsHandler)
	http.HandleFunc("/internal", internal.InteralHandler)
	http.HandleFunc("/internal/task", internal.TaskHandler)
	http.HandleFunc("/api/v1/expressionsSpecial", internal.ExpressionsHandlerSpecial)
	http.HandleFunc("/api/v1/calculateSpecial", internal.CreateRootExpressionHandlerSpecial)
	http.HandleFunc("/api/v1/expressions/", internal.HandlerId)
	http.HandleFunc("/api/v1/expressionsSpecial/", internal.HandlerIdSprcial)
	http.ListenAndServe(":8080", nil)
}

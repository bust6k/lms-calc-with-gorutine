package main

import (
	"net/http"
	"project_yandex_lms/internal"
)

func main() {
	http.HandleFunc("/api/v1/calculate", internal.CreateRootExpressionHandler)
	http.HandleFunc("/api/v1/expressions", internal.ExpressionsHandler)
	http.HandleFunc("/internal", internal.InteralHandler)
	http.ListenAndServe(":8080", nil)
}

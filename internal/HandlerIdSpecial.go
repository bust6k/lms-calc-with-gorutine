package internal

import (
	"html/template"
	"net/http"
	"project_yandex_lms/variables"
	"strconv"
)

func HandlerIdSprcial(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Path[len("/api/v1/expressionsSpecial/"):]

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		http.Error(w, "ошибка при преобразовании id в целочисленный тип", 422)
	}
	for _, expression := range variables.Expressions {
		if expression.Id == valueInt {
			var ResultofEX = template.Must(template.New("one id").Parse(`
<!DOCTYPE html>
<html>
<head>

    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
</head>

<body>
<h1>Результат задачи:</h1> <br> <br>
<p>id задачи:  {{.Id}} </p> <br>
<p> статус выполнения задачи:  {{.Status}} </p> <br>
<p> результат:  {{.Result}}</p> <br>
</body>
</html>
`))
			ResultofEX.Execute(w, expression)
		}
	}
}

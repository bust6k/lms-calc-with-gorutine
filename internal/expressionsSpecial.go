package internal

import (
	"html/template"
	"net/http"
	"project_yandex_lms/variables"
)

func ExpressionsHandlerSpecial(w http.ResponseWriter, r *http.Request) {
	var exprs = template.Must(template.New("expressions").Parse(`
<!DOCTYPE html>
<html>
<head>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
</head>

<body>

<h1>Результаты задач</h1> <br>

{{range .}}
<p>Id данного выражения:  {{.Id}}<br>
<p>статус работы данного выражения:   {{.Status}}<br>
<p>результат работы данного выражения:  {{.Result}}<br> <br> <br> <br>
{{end}}
</body>
</html>
`))
	exprs.Execute(w, variables.Expressions)
}

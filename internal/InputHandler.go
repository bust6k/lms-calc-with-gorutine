package internal

import (
	"encoding/json"
	"html/template"
	"net/http"
	"project_yandex_lms/structures"
	"project_yandex_lms/variables"
	"strings"
)

func CreateRootExpressionHandlerSpecial(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("getexpression")
	var expression structures.RootExpression
	expression.Expression = value
	bytesvalue, err := json.Marshal(expression)
	if err != nil {
		http.Error(w, "ошибка при сериализации выражения в json", 522)
	}

	readerValue := strings.NewReader(string(bytesvalue))
	http.Post("http://localhost:8080/api/v1/calculate", "application/json", readerValue)
	if err != nil {
		http.Error(w, "ошибка при добавлении выражения на сервер", 422)
	}

	var expr structures.RootExpression
	expr.Id = variables.Count_Root_Id

	var ID = template.Must(template.New("showjson").Parse(`
<!DOCTYPE html>
<html>
<head>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
    <title>Информация о задаче</title>  
</head>
<body>
<div class="container">  
    <h1> Id задачи: </h1>

    <p>  ID данной задачи является: {{.}} </p>

    <a href="http://localhost:8080" class="btn btn-primary"> Назад</a>
    <a href="http://localhost:8080/api/v1/expressionsSpecial/{{.}}" class="btn btn-primary">Посмотреть результат выполнения задачи</a>
    <a href="http://localhost:8080/api/v1/expressionsSpecial" class="btn btn-primary">Посмотреть результаты всех задач</a>
</div>

<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"></script>  
</body>
</html>

`))
	ID.Execute(w, variables.Count_Root_Id)

}

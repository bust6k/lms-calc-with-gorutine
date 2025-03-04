package important

import (
	"html/template"
	"net/http"
)

func HandlerHome(w http.ResponseWriter, r *http.Request) {

	tmp, err := template.ParseFiles("ui/html/homepage.html")

	if err != nil {
		http.Error(w, "ошибка при парсинге html шаблона", 500)
	}

	tmp.Execute(w, nil)

}

package important

import (
	"html/template"
	"net/http"
)

func ExpressionSpecialHandler(w http.ResponseWriter, r *http.Request) {
	var Expressions = template.Must(template.New("expressions").Parse(`

<!DOCTYPE html>

<html>

<body>

</html>

</html>
`))
	Expressions.Execute(w, nil)
}

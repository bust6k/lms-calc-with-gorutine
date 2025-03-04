package important

import (
	"encoding/json"
	"io"
	"net/http"
	"project_yandex_lms/structures"
	"project_yandex_lms/variables"
)

var f bool = true

func ExpressionsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var expression structures.Expression
		bytesresp, _ := io.ReadAll(r.Body)
		json.Unmarshal(bytesresp, &expression)
		variables.Expressions = append(variables.Expressions, expression)
	} else if r.Method == http.MethodGet {

		bytesexpressions, _ := json.Marshal(variables.Expressions)
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytesexpressions)

	}
}

package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

func ExpressionsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var expression Expression
		bytesresp, _ := io.ReadAll(r.Body)
		json.Unmarshal(bytesresp, &expression)
		Expressions = append(Expressions, expression)
	} else if r.Method == http.MethodGet {
		bytesexpressions, _ := json.Marshal(Expressions)
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytesexpressions)
	}
}

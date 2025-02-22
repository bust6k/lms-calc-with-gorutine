package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode"
)

type ExpressionToCalcuate struct {
	Expression string ` json:"expression"`
}

func DoCorrect(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("не удалось обратиться по адресу: %s, ошибка: %v", url, err)
		return false
	}
	defer resp.Body.Close()

	bytesJson, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("не удалось прочитать тело ответа, ошибка: %v", err)
		return false
	}

	var ex ExpressionToCalcuate
	err = json.Unmarshal(bytesJson, &ex)
	if err != nil {
		fmt.Errorf("не удалось десериализовать JSON, ошибка: %v", err)
		return false
	}

	expr := strings.TrimSpace(ex.Expression)
	if expr == "" {
		return false
	}

	lastWasOperator := true

	for _, ch := range expr {
		if unicode.IsDigit(ch) {
			lastWasOperator = false
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			if lastWasOperator {
				return false
			}
			lastWasOperator = true
		} else {
			return false
		}
	}

	if lastWasOperator {
		return false
	}

	return true
}

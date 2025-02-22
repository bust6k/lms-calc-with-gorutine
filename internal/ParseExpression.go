package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"project_yandex_lms/structures"
	"strconv"
	"strings"
	"time"
)

var TheTasks []structures.Task

func splitAST(node *structures.ASTNode) []string {
	if node == nil {
		return nil
	}

	var subExpressions []string
	if node.Type == structures.OperatorNode {
		leftExpr := collectExpression(node.Left)
		rightExpr := collectExpression(node.Right)

		subExpressions = append(subExpressions, leftExpr)
		subExpressions = append(subExpressions, node.Value+" "+rightExpr)
	}

	return subExpressions
}

func collectExpression(node *structures.ASTNode) string {
	if node == nil {
		return ""
	}
	if node.Type == structures.NumberNode {
		return node.Value
	}
	return collectExpression(node.Left) + " " + node.Value + " " + collectExpression(node.Right)
}

func ParseExpression(tasks []string) error {
	for i := 0; i < len(tasks); i++ {
		task := tasks[i]

		if task[0] == '+' {
			correctTaskPlusAndMinus := fmt.Sprintf("0 + %s", tasks[i])
			tasks[i] = correctTaskPlusAndMinus
		} else if task[0] == '-' {
			correctTaskPlusAndMinus := fmt.Sprintf("0 - %s", tasks[i])
			tasks[i] = correctTaskPlusAndMinus
		} else if task[0] == '*' {
			correctTaskMultiplieAndDivide := fmt.Sprintf("1 * %s", tasks[i])
			tasks[i] = correctTaskMultiplieAndDivide
		} else if task[0] == '/' {
			correctTaskMultiplieAndDivide := fmt.Sprintf("1 / %s", tasks[i])
			tasks[i] = correctTaskMultiplieAndDivide
		} else if strings.Contains(task[:len(task)-1], "+ - * /") {
			return errors.New("ошибка, последний элемент вражения содержит не операнд, а оператор")
		}

		parsingToTask := func(ex string) structures.Task {
			ex = strings.ReplaceAll(ex, " ", "")
			var FloatArg1 float64
			var FloatArg2 float64
			var Operator string
			for i := 0; i < len(ex); i++ {
				if ex[i] == '+' || ex[i] == '-' || ex[i] == '*' || ex[i] == '/' {
					FloatArg1, _ = strconv.ParseFloat(ex[:i], 64)
					Operator = ex[i : i+1]
					FloatArg2, _ = strconv.ParseFloat(ex[i:], 64)

				}
			}

			return structures.Task{Id: 3, Arg1: FloatArg1, Arg2: FloatArg2, Operation: Operator, Operation_time: 1 * time.Second}
		}(tasks[i])

		TheTasks = append(TheTasks, parsingToTask)

	}
	jsonbytes, err := json.Marshal(TheTasks)
	if err != nil {
		fmt.Errorf("ошибка при сериализации задач в json: %v", err)
	}
	body := bytes.NewReader(jsonbytes)
	http.Post("http://localhost:8080/internal", "application/json", body)

	return nil
}

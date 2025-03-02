package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project_yandex_lms/structures"
	"project_yandex_lms/variables"
	"strconv"
	"strings"
)

func SplitAST(node *structures.ASTNode) []structures.Task {
	if node == nil {
		return nil
	}

	var tasks []structures.Task

	if node.Type == structures.OperatorNode {
		leftTasks := SplitAST(node.Left)
		tasks = append(tasks, leftTasks...)

		rightTasks := SplitAST(node.Right)
		tasks = append(tasks, rightTasks...)

		task := structures.Task{
			Id:        variables.Count_Root_Id,
			Arg1:      parseOperand(node.Left),
			Operation: node.Value,
			Arg2:      parseOperand(node.Right),
		}
		tasks = append(tasks, task)
	}

	return tasks
}
func parseOperand(node *structures.ASTNode) float64 {
	if node.Type == structures.NumberNode {
		value, _ := strconv.ParseFloat(node.Value, 64)
		return value
	} else if node.Type == structures.OperatorNode {
		left := parseOperand(node.Left)
		right := parseOperand(node.Right)
		switch node.Value {
		case "+":
			return left + right
		case "-":
			return left - right
		case "*":
			return left * right
		case "/":
			return left / right
		}
	}
	return 0
}

func collectExpression(node *structures.ASTNode) string {
	if node == nil {
		return ""
	}
	if node.Type == structures.NumberNode {
		return node.Value
	}

	if node.Value == "+" || node.Value == "-" {
		return "(" + collectExpression(node.Left) + " " + node.Value + " " + collectExpression(node.Right) + ")"
	}
	return collectExpression(node.Left) + " " + node.Value + " " + collectExpression(node.Right)
}

func PostExpression(tasks []structures.Task) error {

	bytesjson, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	body := strings.NewReader(string(bytesjson))
	resp, err := http.Post("http://localhost:8080/internal", "application/json", body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка, отправка не является успешной")
	}

	if err != nil {
		return err
	}
	return nil
}

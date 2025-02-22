package internal

import (
	"fmt"
	"project_yandex_lms/structures"
	"strconv"
	"strings"
)

func buildAST(rpn []string) (*structures.ASTNode, error) {
	var stack []*structures.ASTNode

	isOperator := func(s string) bool {
		_, ok := structures.Operators[s]
		return ok
	}

	for _, token := range rpn {
		if !isOperator(token) {
			// Это число.
			node := &structures.ASTNode{Type: structures.NumberNode, Value: token}
			stack = append(stack, node)
		} else {
			// Это оператор.
			if len(stack) < 2 {
				return nil, fmt.Errorf("недостаточно операндов для оператора: %s", token)
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			node := &structures.ASTNode{Type: structures.OperatorNode, Value: token, Left: left, Right: right}
			stack = append(stack, node)
		}
	}

	if len(stack) != 1 {
		return nil, fmt.Errorf("недопустимое RPN выражение")
	}

	return stack[0], nil
}
func splitExpression(expression string) ([]string, error) {
	rpn, err := infixToRPN(expression)
	if err != nil {
		return nil, err
	}

	var subExpressions []string
	currentExpression := ""
	operandCount := 0

	for _, token := range rpn {
		currentExpression += token + " "

		if _, err := strconv.Atoi(token); err == nil {
			operandCount++
		} else if _, ok := structures.Operators[token]; ok {
			operandCount--
		}

		if operandCount == 1 && len(currentExpression) > 0 {
			// Найден конец подвыражения.
			subExpressions = append(subExpressions, strings.TrimSpace(currentExpression))
			currentExpression = ""
			operandCount = 0
		}
	}

	// Если что-то осталось в currentExpression, добавляем это как последнее подвыражение.
	if len(currentExpression) > 0 {
		subExpressions = append(subExpressions, strings.TrimSpace(currentExpression))
	}

	return subExpressions, nil
}

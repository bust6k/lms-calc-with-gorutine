package important

import (
	"fmt"
	"project_yandex_lms/structures"
	"project_yandex_lms/variables"
)

func BuildAST(rpn []string) (*structures.ASTNode, error) {
	var stack []*structures.ASTNode

	isOperator := func(s string) bool {
		_, ok := variables.Operators[s]
		return ok
	}

	for _, token := range rpn {
		if !isOperator(token) {

			node := &structures.ASTNode{Type: structures.NumberNode, Value: token}
			stack = append(stack, node)
		} else {

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

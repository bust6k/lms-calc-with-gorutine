package Tests

import (
	"fmt"
	"project_yandex_lms/important"
	"project_yandex_lms/structures"
	"reflect"
	"testing"
)

func equalAST(node1 *structures.ASTNode, node2 *structures.ASTNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return reflect.DeepEqual(node1, node2)
}
func TestBuildAST(t *testing.T) {
	tests := []struct {
		rpn      []string
		expected *structures.ASTNode
		err      error
	}{
		{[]string{"3", "4", "+"}, &structures.ASTNode{Type: structures.OperatorNode, Value: "+", Left: &structures.ASTNode{Type: structures.NumberNode, Value: "3"}, Right: &structures.ASTNode{Type: structures.NumberNode, Value: "4"}}, nil},
		{[]string{"3", "4", "2", "*", "+"}, &structures.ASTNode{Type: structures.OperatorNode, Value: "+", Left: &structures.ASTNode{Type: structures.NumberNode, Value: "3"}, Right: &structures.ASTNode{Type: structures.OperatorNode, Value: "*", Left: &structures.ASTNode{Type: structures.NumberNode, Value: "4"}, Right: &structures.ASTNode{Type: structures.NumberNode, Value: "2"}}}, nil},
		{[]string{"3", "4", "+", "2", "*"}, &structures.ASTNode{Type: structures.OperatorNode, Value: "*", Left: &structures.ASTNode{Type: structures.OperatorNode, Value: "+", Left: &structures.ASTNode{Type: structures.NumberNode, Value: "3"}, Right: &structures.ASTNode{Type: structures.NumberNode, Value: "4"}}, Right: &structures.ASTNode{Type: structures.NumberNode, Value: "2"}}, nil},
		{[]string{"3", "4", "+", "*"}, nil, fmt.Errorf("недостаточно операндов для оператора: *")},
	}

	for _, test := range tests {
		result, err := important.BuildAST(test.rpn)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Ошибка не совпадает для RPN %v: ожидалось %v, получено %v", test.rpn, test.err, err)
		}
		if !equalAST(result, test.expected) {
			t.Errorf("Результат не совпадает для RPN %v: ожидалось %v, получено %v", test.rpn, *test.expected, *result)
		}
	}
}

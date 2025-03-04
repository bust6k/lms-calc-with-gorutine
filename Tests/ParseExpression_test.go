package Tests

import (
	"project_yandex_lms/important"
	"project_yandex_lms/structures"
	"project_yandex_lms/variables"
	"reflect"
	"testing"
)

func TestSplitAST(t *testing.T) {
	tests := []struct {
		node     *structures.ASTNode
		expected []structures.Task
	}{
		{&structures.ASTNode{Type: structures.OperatorNode, Value: "+", Left: &structures.ASTNode{Type: structures.NumberNode, Value: "3"}, Right: &structures.ASTNode{Type: structures.NumberNode, Value: "4"}}, []structures.Task{{Id: variables.Count_Root_Id, Arg1: 3, Operation: "+", Arg2: 4}}},
		{&structures.ASTNode{Type: structures.OperatorNode, Value: "*", Left: &structures.ASTNode{Type: structures.OperatorNode, Value: "+", Left: &structures.ASTNode{Type: structures.NumberNode, Value: "3"}, Right: &structures.ASTNode{Type: structures.NumberNode, Value: "4"}}, Right: &structures.ASTNode{Type: structures.NumberNode, Value: "2"}}, []structures.Task{{Id: variables.Count_Root_Id, Arg1: 3, Operation: "+", Arg2: 4}, {Id: variables.Count_Root_Id, Arg1: 7, Operation: "*", Arg2: 2}}},
	}

	for _, test := range tests {
		result := important.SplitAST(test.node)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Результат не совпадает для AST %v: ожидалось %v, получено %v", test.node, test.expected, result)
		}
	}

}

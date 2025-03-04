package Tests

import (
	"fmt"
	"project_yandex_lms/important"
	"reflect"

	"testing"
)

func TestInfixToRPN(t *testing.T) {
	tests := []struct {
		expression string
		expected   []string
		err        error
	}{
		{"3 + 4", []string{"3", "4", "+"}, nil},
		{"3 + 4 * 2 / (1 - 5)", []string{"3", "4", "2", "*", "1", "5", "-", "/", "+"}, nil},
		{"3 + (4 * 2)", []string{"3", "4", "2", "*", "+"}, nil},
		{"3 + 4 * (2 a)", nil, fmt.Errorf("недопустимый символ: a")},
		{"3 + 4 * 2 - 1)(", nil, fmt.Errorf("несоответствие скобок")},
	}

	for _, test := range tests {
		result, err := important.InfixToRPN(test.expression)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Ошибка не совпадает для выражения %s: ожидалось %v, получено %v", test.expression, test.err, err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Результат не совпадает для выражения %s: ожидалось %v, получено %v", test.expression, test.expected, result)
		}
	}
}

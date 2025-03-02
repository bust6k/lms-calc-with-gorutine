package structures

import (
	"time"
)

type RootExpression struct {
	Id         int    `json:"id,omitempty"`
	Expression string `json:"expression,omitempty" `
}

type NodeType int

const (
	NumberNode NodeType = iota
	OperatorNode
)

type ASTNode struct {
	Type  NodeType
	Value string // Значение (число или оператор).
	Left  *ASTNode
	Right *ASTNode
}

type Task struct {
	Id             int           `json:"id"`
	Arg1           float64       `json:"arg1"`
	Arg2           float64       `json:"arg2"`
	Operation      string        `json:"operation"`
	Operation_time time.Duration `json:"operation_time"`
}

type Expression struct {
	Id     int     `json:"id"`
	Status string  `json:"status"`
	Result float64 `json:"result"`
}

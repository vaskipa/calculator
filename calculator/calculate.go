package calculator

import (
	"errors"
	"unicode"
)

type NodeData struct {
	isOperation bool
	operation   rune
	digit       int
}

type Node struct {
	left  *Node
	right *Node
	NodeData
}

func ToPolishNotation(expression string) ([]NodeData, error) {
	data := []rune(expression)
	polishNotation := make([]NodeData, 0)

	stack := make([]NodeData, 0)
	digit := 0
	isDigit := false

	for _, value := range data {
		if unicode.IsDigit(value) {
			digit = digit*10 + int(value-'0')
			isDigit = true
			continue
		}
		if isDigit {
			isDigit = false
			polishNotation = append(polishNotation, NodeData{false, 0, digit})
			digit = 0
		}

		if value == '(' {
			stack = append(stack, NodeData{true, value, 0})
		}
		if value == ')' {
			for {
				if len(stack) == 0 {
					return nil, errors.New("no (")
				}
				var operator NodeData
				stack, operator = stack[:len(stack)-1], stack[len(stack)-1]
				if operator.isOperation && operator.operation == '(' {
					break
				}
				polishNotation = append(polishNotation, operator)
			}
		}
		if value == '+' || value == '-' || value == '*' || value == '/' {
			stack = append(stack, NodeData{true, value, 0})
		}
	}
	if isDigit {
		isDigit = false
		polishNotation = append(polishNotation, NodeData{false, 0, digit})
		digit = 0
	}
	var value NodeData
	for {
		if len(stack) == 0 {
			break
		}
		stack, value = stack[:len(stack)-1], stack[len(stack)-1]
		if value.operation == '(' {
			return nil, errors.New("too many operators")
		}
		polishNotation = append(polishNotation, value)
	}
	return polishNotation, nil

}

func fromNodeData(nodeData NodeData) *Node {
	return &Node{NodeData: nodeData, left: nil, right: nil}
}

func GenerateAST(data []NodeData, i *int) *Node {
	/*
		Перевод из обратной пользской записи в дерево
	*/
	head := fromNodeData(data[*i])
	*i--

	if head.isOperation {
		head.left = GenerateAST(data, i)
		head.right = GenerateAST(data, i)
	}
	return head
}

func Calculate(node *Node) (float64, error) {

	if !node.isOperation {
		return float64(node.digit), nil
	}

	left, err := Calculate(node.left)
	if err != nil {
		return 0, err
	}
	right, err := Calculate(node.right)
	if err != nil {
		return 0, err
	}
	if node.operation == '+' {
		return left + right, nil
	}
	if node.operation == '-' {
		// нужно проверить потом, тут может перепутано
		return left - right, nil
	}
	if node.operation == '*' {
		return left * right, nil
	}

	// это тоже нжуно будет проверить
	if right == 0 {
		return 0, errors.New("division by zero")
	}
	return left / right, nil

}

func Calc(expression string) (float64, error) {
	polishNotation, err := ToPolishNotation(expression)
	if err != nil {
		return 0, err
	}
	size := len(polishNotation) - 1
	head := GenerateAST(polishNotation, &size)
	return Calculate(head)
}

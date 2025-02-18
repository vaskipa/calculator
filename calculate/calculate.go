package calculate

import (
	"errors"
	"unicode"
)

func Calculate(polishNotation []rune) (float64, error) {
	resultStack := make([]float64, 0)
	for _, value := range polishNotation {
		if unicode.IsDigit(value) {
			resultStack = append(resultStack, float64(value)-48)
		} else {
			if len(resultStack) < 2 {
				return 0, errors.New("too few digits")
			}
			var el float64
			if value == '+' {
				el = resultStack[len(resultStack)-2] + resultStack[len(resultStack)-1]
			}
			if value == '*' {
				el = resultStack[len(resultStack)-2] * resultStack[len(resultStack)-1]
			}
			if value == '-' {
				el = resultStack[len(resultStack)-2] - resultStack[len(resultStack)-1]
			}
			if value == '/' {
				el = resultStack[len(resultStack)-2] / resultStack[len(resultStack)-1]
			}
			resultStack = resultStack[:len(resultStack)-2]
			resultStack = append(resultStack, el)
		}
	}
	if len(resultStack) == 1 {
		return resultStack[0], nil
	}
	return 0, errors.New("too many digits")
}

func ToPolishNotation(expression string) ([]rune, error) {
	data := []rune(expression)
	polishNotation := make([]rune, 0)
	stack := make([]rune, 0)
	for _, value := range data {
		if unicode.IsDigit(value) {
			polishNotation = append(polishNotation, value)
		}
		if value == '(' {
			stack = append(stack, value)
		}
		if value == ')' {
			for {
				if len(stack) == 0 {
					return []rune(""), errors.New("no (")
				}
				stack, value = stack[:len(stack)-1], stack[len(stack)-1]
				if value == '(' {
					break
				}
				polishNotation = append(polishNotation, value)
			}
		}
		if value == '+' || value == '-' || value == '*' || value == '/' {
			stack = append(stack, value)
		}
	}
	var value rune
	for {
		if len(stack) == 0 {
			break
		}
		stack, value = stack[:len(stack)-1], stack[len(stack)-1]
		if value == '(' {
			return []rune(""), errors.New("too many operators")
		}
		polishNotation = append(polishNotation, value)
	}
	return polishNotation, nil

}

func Calc(expression string) (float64, error) {
	polishNotation, err := ToPolishNotation(expression)
	if err != nil {
		return 0, err
	}
	return Calculate(polishNotation)
}

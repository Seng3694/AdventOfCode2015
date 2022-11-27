package main

import (
	"aocutil"
	"fmt"
	"regexp"
	"strconv"
)

// opcode
const (
	OpNot = iota
	OpAssign
	OpAnd
	OpOr
	OpRShift
	OpLShift
)

// operandType
const (
	OperandConstant = iota
	OperandVariable
)

type Operand struct {
	operandType int
	value       string
}

type Expression struct {
	lhs, rhs  Operand
	operation int
}

func get_regex_params(r *regexp.Regexp, str string) (parameters map[string]string) {
	match := r.FindStringSubmatch(str)
	parameters = make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			parameters[name] = match[i]
		}
	}
	return
}

func get_operand_type(value string) int {
	if _, err := strconv.Atoi(value); err == nil {
		return OperandConstant
	} else {
		return OperandVariable
	}
}

func parse_line(line string, variablesMap map[string]Expression) {
	binaryOperationRegex := regexp.MustCompile(`^(?P<operand1>[a-z0-9]+)\s+(?P<operation>[A-Z]+)\s+(?P<operand2>[a-z0-9]+)\s+(\-\>)\s+(?P<dest>[a-z0-9]+)$`)
	notOperationRegex := regexp.MustCompile(`^(NOT)\s+(?P<operand>[a-z0-9]+)\s+(\-\>)\s+(?P<dest>[a-z0-9]+)$`)
	assignmentRegex := regexp.MustCompile(`^(?P<operand>[a-z0-9]+)\s+(\-\>)\s+(?P<dest>[a-z0-9]+)$`)

	if params := get_regex_params(binaryOperationRegex, line); len(params) > 0 {
		var (
			operand1, operand2 Operand
		)

		operand1.value = params["operand1"]
		operand1.operandType = get_operand_type(operand1.value)

		operand2.value = params["operand2"]
		operand2.operandType = get_operand_type(operand2.value)

		expr := Expression{lhs: operand1, rhs: operand2}

		switch params["operation"] {
		case "AND":
			expr.operation = OpAnd
		case "OR":
			expr.operation = OpOr
		case "LSHIFT":
			expr.operation = OpLShift
		case "RSHIFT":
			expr.operation = OpRShift
		default:
			fmt.Printf("Undefined operation %v\n", params["operation"])
		}

		variablesMap[params["dest"]] = expr
	} else if params := get_regex_params(notOperationRegex, line); len(params) > 0 {
		operand := Operand{
			value:       params["operand"],
			operandType: get_operand_type(params["operand"]),
		}
		variablesMap[params["dest"]] = Expression{lhs: operand, operation: OpNot}
	} else if params := get_regex_params(assignmentRegex, line); len(params) > 0 {
		operand := Operand{
			value:       params["operand"],
			operandType: get_operand_type(params["operand"]),
		}
		variablesMap[params["dest"]] = Expression{lhs: operand, operation: OpAssign}
	}
}

func evaluate(variable string, variablesMap map[string]Expression, evaluationCache map[string]uint16) uint16 {
	if value, contains := evaluationCache[variable]; contains {
		return value
	}
	expr := variablesMap[variable]

	lhsIntValue, _ := strconv.Atoi(expr.lhs.value)
	rhsIntValue, _ := strconv.Atoi(expr.rhs.value)

	lhsExpr := func() uint16 {
		if expr.lhs.operandType == OperandConstant {
			return uint16(lhsIntValue)
		} else {
			return evaluate(expr.lhs.value, variablesMap, evaluationCache)
		}
	}
	rhsExpr := func() uint16 {
		if expr.rhs.operandType == OperandConstant {
			return uint16(rhsIntValue)
		} else {
			return evaluate(expr.rhs.value, variablesMap, evaluationCache)
		}
	}

	var returnValue uint16

	switch expr.operation {
	case OpNot:
		returnValue = ^lhsExpr()
	case OpAssign:
		returnValue = lhsExpr()
	case OpAnd:
		returnValue = lhsExpr() & rhsExpr()
	case OpOr:
		returnValue = lhsExpr() | rhsExpr()
	case OpRShift:
		returnValue = lhsExpr() >> rhsExpr()
	case OpLShift:
		returnValue = lhsExpr() << rhsExpr()
	default:
		fmt.Printf("invald operation %v\n", expr.operation)
		return 0
	}

	evaluationCache[variable] = returnValue
	return returnValue
}

func main() {
	variables := make(map[string]Expression)

	aocutil.FileReadAllLines("day7/input.txt", func(s string) {
		parse_line(s, variables)
	})

	evaluationCache := make(map[string]uint16)
	resultPart1 := evaluate("a", variables, evaluationCache)

	evaluationCache = make(map[string]uint16)
	evaluationCache["b"] = resultPart1
	resultPart2 := evaluate("a", variables, evaluationCache)

	aocutil.AOCFinish(resultPart1, resultPart2)
}

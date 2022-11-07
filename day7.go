package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

/*
--- Day 7: Some Assembly Required ---

This year, Santa brought little Bobby Tables a set of wires and bitwise logic gates! Unfortunately, little Bobby is a little under the recommended age range, and he needs help assembling the circuit.

Each wire has an identifier (some lowercase letters) and can carry a 16-bit signal (a number from 0 to 65535). A signal is provided to each wire by a gate, another wire, or some specific value. Each wire can only get a signal from one source, but can provide its signal to multiple destinations. A gate provides no signal until all of its inputs have a signal.

The included instructions booklet describes how to connect the parts together: x AND y -> z means to connect wires x and y to an AND gate, and then connect its output to wire z.

For example:

    123 -> x means that the signal 123 is provided to wire x.
    x AND y -> z means that the bitwise AND of wire x and wire y is provided to wire z.
    p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 and then provided to wire q.
    NOT e -> f means that the bitwise complement of the value from wire e is provided to wire f.

Other possible gates include OR (bitwise OR) and RSHIFT (right-shift). If, for some reason, you'd like to emulate the circuit instead, almost all programming languages (for example, C, JavaScript, or Python) provide operators for these gates.

For example, here is a simple circuit:

123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i

After it is run, these are the signals on the wires:

d: 72
e: 507
f: 492
g: 114
h: 65412
i: 65079
x: 123
y: 456

In little Bobby's kit's instructions booklet (provided as your puzzle input), what signal is ultimately provided to wire a?

--- Part Two ---

Now, take the signal you got on wire a, override wire b to that signal, and reset the other wires (including wire a). What new signal is ultimately provided to wire a?

*/

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

func day7() (string, string) {
	file, err := os.Open("input/day7.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	variables := make(map[string]Expression)

	for scanner.Scan() {
		line := scanner.Text()
		parse_line(line, variables)
	}

	evaluationCache := make(map[string]uint16)
	resultPart1 := evaluate("a", variables, evaluationCache)

	evaluationCache = make(map[string]uint16)
	evaluationCache["b"] = resultPart1
	resultPart2 := evaluate("a", variables, evaluationCache)

	return fmt.Sprint(resultPart1), fmt.Sprint(resultPart2)
}

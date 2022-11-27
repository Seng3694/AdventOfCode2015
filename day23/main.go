package main

import (
	"aocutil"
	"strconv"
	"strings"
)

type Registers struct {
	a, b uint32
	pc   int32
}

const (
	INSTR_HLF_A uint8 = iota
	INSTR_HLF_B
	INSTR_TPL_A
	INSTR_TPL_B
	INSTR_INC_A
	INSTR_INC_B
	INSTR_JMP
	INSTR_JIE_A
	INSTR_JIE_B
	INSTR_JIO_A
	INSTR_JIO_B
	INSTR_EOF
)

type Instruction struct {
	instr uint8
	arg   int8
}

func parse_file(fileName string) []Instruction {
	instructions := make([]Instruction, 0, 32)
	aocutil.FileReadAllLines(fileName, func(s string) {
		s = strings.Replace(s, ",", "", -1)
		split := strings.Split(s, " ")

		offset := uint8(0)
		if split[1] == "b" {
			offset = 1
		}

		switch split[0] {
		case "hlf":
			instructions = append(instructions, Instruction{instr: INSTR_HLF_A + offset})
		case "tpl":
			instructions = append(instructions, Instruction{instr: INSTR_TPL_A + offset})
		case "inc":
			instructions = append(instructions, Instruction{instr: INSTR_INC_A + offset})
		case "jmp":
			jmpOffset, _ := strconv.Atoi(split[1])
			instructions = append(instructions, Instruction{instr: INSTR_JMP, arg: int8(jmpOffset)})
		case "jie":
			jmpOffset, _ := strconv.Atoi(split[2])
			instructions = append(instructions, Instruction{instr: INSTR_JIE_A + offset, arg: int8(jmpOffset)})
		case "jio":
			jmpOffset, _ := strconv.Atoi(split[2])
			instructions = append(instructions, Instruction{instr: INSTR_JIO_A + offset, arg: int8(jmpOffset)})
		}
	})
	instructions = append(instructions, Instruction{instr: INSTR_EOF})
	return instructions
}

func run_program(r *Registers, instructions []Instruction) {
	finished := false
	for !finished {
		switch instructions[r.pc].instr {
		case INSTR_HLF_A:
			r.a /= 2
			r.pc++
		case INSTR_HLF_B:
			r.b /= 2
			r.pc++
		case INSTR_TPL_A:
			r.a *= 3
			r.pc++
		case INSTR_TPL_B:
			r.b *= 3
			r.pc++
		case INSTR_INC_A:
			r.a++
			r.pc++
		case INSTR_INC_B:
			r.b++
			r.pc++
		case INSTR_JMP:
			r.pc += int32(instructions[r.pc].arg)
		case INSTR_JIE_A:
			if r.a%2 == 0 {
				r.pc += int32(instructions[r.pc].arg)
			} else {
				r.pc++
			}
		case INSTR_JIE_B:
			if r.b%2 == 0 {
				r.pc += int32(instructions[r.pc].arg)
			} else {
				r.pc++
			}
		case INSTR_JIO_A:
			if r.a == 1 {
				r.pc += int32(instructions[r.pc].arg)
			} else {
				r.pc++
			}
		case INSTR_JIO_B:
			if r.b == 1 {
				r.pc += int32(instructions[r.pc].arg)
			} else {
				r.pc++
			}
		case INSTR_EOF:
			finished = true
		}
	}
}

func main() {
	instructions := parse_file("day23/input.txt")
	r := Registers{}

	run_program(&r, instructions)
	part1 := r.b

	r = Registers{a: 1}
	run_program(&r, instructions)
	part2 := r.b

	aocutil.AOCFinish(part1, part2)
}

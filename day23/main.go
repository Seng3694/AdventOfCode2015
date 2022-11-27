package main

import (
	"aocutil"
	"strconv"
	"strings"
)

/*
--- Day 23: Opening the Turing Lock ---

Little Jane Marie just got her very first computer for Christmas from some unknown benefactor. It comes with instructions and an example program, but the computer itself seems to be malfunctioning. She's curious what the program does, and would like you to help her run it.

The manual explains that the computer supports two registers and six instructions (truly, it goes on to remind the reader, a state-of-the-art technology). The registers are named a and b, can hold any non-negative integer, and begin with a value of 0. The instructions are as follows:

    hlf r sets register r to half its current value, then continues with the next instruction.
    tpl r sets register r to triple its current value, then continues with the next instruction.
    inc r increments register r, adding 1 to it, then continues with the next instruction.
    jmp offset is a jump; it continues with the instruction offset away relative to itself.
    jie r, offset is like jmp, but only jumps if register r is even ("jump if even").
    jio r, offset is like jmp, but only jumps if register r is 1 ("jump if one", not odd).

All three jump instructions work with an offset relative to that instruction. The offset is always written with a prefix + or - to indicate the direction of the jump (forward or backward, respectively). For example, jmp +1 would simply continue with the next instruction, while jmp +0 would continuously jump back to itself forever.

The program exits when it tries to run an instruction beyond the ones defined.

For example, this program sets a to 2, because the jio instruction causes it to skip the tpl instruction:

inc a
jio a, +2
tpl a
inc a

What is the value in register b when the program in your puzzle input is finished executing?

--- Part Two ---

The unknown benefactor is very thankful for releasi-- er, helping little Jane Marie with her computer. Definitely not to distract you, what is the value in register b after the program is finished executing if register a starts as 1 instead?

*/

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

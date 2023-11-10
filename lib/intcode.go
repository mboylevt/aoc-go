package lib

import (
	"fmt"

	"github.com/mboylevt/aoc-go/cast"
)

type instruction struct {
	operation int
	modeP1    byte
	modeP2    byte
	modeP3    byte
	p1        int
	p2        int
	p3        int
	advance   int
}

func parseInstruction(program []int, ip int) instruction {
	opcode := cast.ToString(program[ip])
	if len(opcode) < 5 {
		for {
			opcode = "0" + opcode
			if len(opcode) == 5 {
				break
			}
		}
	}
	operation := cast.ToInt(opcode[len(opcode)-2:])
	modes := opcode[0 : len(opcode)-2]
	ins := instruction{}
	switch operation {
	case 1, 2, 7, 8: //add, multiply, lessthan, equals
		ins = instruction{operation: operation,
			modeP1:  modes[2] - 48,
			modeP2:  modes[1] - 48,
			modeP3:  modes[0] - 48,
			p1:      program[ip+1],
			p2:      program[ip+2],
			p3:      program[ip+3],
			advance: 4}
	case 3: //input
		ins = instruction{operation: operation, p1: program[ip+1], advance: 2}
	case 4: //output
		ins = instruction{operation: operation, p1: program[ip+1], modeP1: modes[2] - 48, advance: 2}
	case 5, 6: //jump-if-true/false,
		ins = instruction{operation: operation,
			modeP1:  modes[2] - 48,
			modeP2:  modes[1] - 48,
			p1:      program[ip+1],
			p2:      program[ip+2],
			advance: 0}
	case 99: //halt
		ins = instruction{operation: 99, advance: 1}
	}
	return ins
}

func executeInstruction(program []int, ins instruction, ip *int) int {
	_ = 1
	fmt.Printf("\tExecuting %v\n", ins)
	switch ins.operation {
	case 1, 2, 7, 8:
		// handle parameter modes
		param1 := ins.p1
		param2 := ins.p2
		param3 := ins.p3
		if ins.modeP1 == 0 {
			param1 = program[param1]
		}
		if ins.modeP2 == 0 {
			param2 = program[param2]
		}

		// perform operation
		if ins.operation == 1 {
			program[param3] = param1 + param2
		} else if ins.operation == 2 {
			program[param3] = param1 * param2
		} else if ins.operation == 7 {
			if param1 < param2 {
				program[param3] = 1
			} else {
				program[param3] = 0
			}
		} else if ins.operation == 8 {
			if param1 == param2 {
				program[param3] = 1
			} else {
				program[param3] = 0
			}
		}
	case 3:
		fmt.Print("Enter input value: ")
		var input string
		fmt.Scanf("%s", &input)
		program[ins.p1] = cast.ToInt(input)
	case 4:
		param1 := ins.p1
		if ins.modeP1 == 0 {
			param1 = program[param1]
		}
		fmt.Printf("Output: %v\n", (param1))
	case 5, 6:
		//handle params and modes
		param1 := ins.p1
		param2 := ins.p2
		if ins.modeP1 == 0 {
			param1 = program[param1]
		}
		if ins.modeP2 == 0 {
			param2 = program[param2]
		}
		// perform jumps
		if ins.operation == 5 {
			if param1 != 0 {
				*ip = param2
			} else {
				return 3
			}

		} else if ins.operation == 6 {
			if param1 == 0 {
				*ip = param2
			} else {
				return 3
			}
		}
	}
	return 0
}

func RunProgram(program []int) int {
	ip := 0
	for {
		instruction := parseInstruction(program, ip)
		if instruction.operation == 99 {
			return 0
		}
		jmp := executeInstruction(program, instruction, &ip)
		if jmp > 0 {
			ip += jmp
		} else {
			ip += instruction.advance
		}
	}
}

func RunProgramDay2(program []int) int {
	var instruction_pointer = 0
	for {
		var operation = program[instruction_pointer]
		var arg1 = program[instruction_pointer+1]
		var arg2 = program[instruction_pointer+2]
		var result = program[instruction_pointer+3]
		switch operation {
		case 1: // addition
			program[result] = program[arg1] + program[arg2]
			instruction_pointer += 4
		case 2: // multiplication
			program[result] = program[arg1] * program[arg2]
			instruction_pointer += 4
		case 99:
			instruction_pointer += 1
			return program[0]
		}
	}
}

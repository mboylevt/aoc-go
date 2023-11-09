package lib

func RunProgram(program []int) int {
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

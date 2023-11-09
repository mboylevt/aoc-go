package util

func IntCode(program []int) int {
	var program_counter = 0
	for {
		var operation = program[program_counter]
		var arg1 = program[program_counter+1]
		var arg2 = program[program_counter+2]
		var result = program[program_counter+3]
		switch operation {
		case 1:
			program[result] = program[arg1] + program[arg2]
		case 2:
			program[result] = program[arg1] * program[arg2]
		case 99:
			return program[0]
		}
		program_counter += 4
	}
}

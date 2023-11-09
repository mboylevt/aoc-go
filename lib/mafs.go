package lib

func AbsInt(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

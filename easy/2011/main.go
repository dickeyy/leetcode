package main

func finalValueAfterOperations(operations []string) int {
	x := 0
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "--X":
			x--
		case "X--":
			x--
		case "++X":
			x++
		case "X++":
			x++
		}
	}

	return x
}

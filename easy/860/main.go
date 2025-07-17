package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 9.68mb (beats 64.49% of Go solutions)
 */

func lemonadeChange(bills []int) bool {
	fives := 0
	tens := 0

	for _, bill := range bills {
		switch bill {
		case 5:
			// Customer pays $5, no change needed
			fives++
		case 10:
			// Customer pays $10, need to give $5 change
			if fives == 0 {
				return false // Can't make change
			}
			fives--
			tens++
		case 20:
			// Customer pays $20, need to give $15 change
			// Prefer giving one $10 + one $5 over three $5s
			if tens > 0 && fives > 0 {
				tens--
				fives--
			} else if fives >= 3 {
				fives -= 3
			} else {
				return false // Can't make change
			}
		}
	}
	return true
}

func main() {
	got := lemonadeChange([]int{5, 5, 10, 10, 20})
	println(got, "should be false")
}

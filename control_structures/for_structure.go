package control_structures

func collatzChainLength(n int) int {
	count := 0
	if n == 1 {
		return count;
	} else if n % 2 == 0 {
		count++
		count := count + collatzChainLength(n / 2)
		return count
	} else {
		count++
		count := count + collatzChainLength(3 * n + 1)
		return count
	}
}

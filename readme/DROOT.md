func cor(n int) int {
	for n > 9 {
		s := 0
		for n > 0 {
			s += n % 10
			n = n / 10
		}
		n = s
	}
	return n
}

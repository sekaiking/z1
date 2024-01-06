package piscine

func Abort(a, b, c, d, e int) int {
	// bubble sort
	for i := 0; i < 5; i++ {
		if a > b {
			a, b = b, a
		}
		if b > c {
			b, c = c, b
		}
		if c > d {
			c, d = d, c
		}
		if d > e {
			d, e = e, d
		}
	}
	// return middle value
	return c
}

package add

// Sub returns a minus b.
func Sub(a, b int) int {
	return a - b
}

// Neg returns -a.
func Neg(a int) int {
	return Sub(0, a)
}

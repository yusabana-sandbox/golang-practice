package main

// Sqrt ok
func Sqrt(x float64) float64 {
	z := 2.0
	for i := 0; i < 10000000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

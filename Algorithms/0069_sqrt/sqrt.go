package sqrt_0069

func mySqrt(x int) int {
	sqrt := x
	for sqrt*sqrt > x {
		sqrt = (sqrt + x/sqrt) / 2
	}
	return sqrt
}

package math

func Sum(nums ...float64) float64 {
	var s float64
	for _, v := range nums {
		s += v
	}
	return s
}

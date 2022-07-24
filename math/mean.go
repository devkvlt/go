package math

func Mean(nums ...float64) float64 {
	return Sum(nums...) / float64(len(nums))
}

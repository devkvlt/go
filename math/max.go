package math

func Max(nums ...float64) float64 {
	m := nums[0]
	for _, v := range nums {
		if m < v {
			m = v
		}
	}
	return float64(m)
}

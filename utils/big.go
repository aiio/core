package utils

func GetBig(data []float64) float64 {
	var tmp float64
	for _, datum := range data {
		if datum > tmp {
			tmp = datum
		}
	}
	return tmp
}

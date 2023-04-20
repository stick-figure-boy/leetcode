package leetcode

// https://leetcode.com/problems/convert-the-temperature/
// Runtime: 1 ms
// Memory Usage: 2 MB

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.80 + 32.00}
}
